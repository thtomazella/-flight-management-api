package controllers

import (
	"encoding/json"
	"errors"
	"flight/src/authentication"
	"flight/src/database"
	"flight/src/models"
	"flight/src/repository"
	"flight/src/response"
	"flight/src/security"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Login é responsável por autenticar um usuário na API
func Login(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
	}

	var usuario models.Usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuarios(db)
	usuarioSalvoNoBanco, erro := repositorio.BuscarPorEmail(usuario.Email)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerificarSenha(usuarioSalvoNoBanco.Senha, usuario.Senha); erro != nil {
		// response.Erro(w, http.StatusUnauthorized, erro)
		response.Erro(w, http.StatusInternalServerError, errors.New(" Usuário ou senha não identificado! "))
		return
	}

	token, erro := authentication.CriarToken(usuarioSalvoNoBanco.ID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	usuarioID := strconv.FormatUint(usuarioSalvoNoBanco.ID, 10)

	response.JSON(w, http.StatusOK, models.DadosAutenticacao{ID: usuarioID, Token: token})

}

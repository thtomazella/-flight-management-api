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
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarUsuario insere um usuario no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioDuplicado := repository.NovoRepositoDeUsuarios(db)
	if codAnacLancto, _ := repositorioDuplicado.BuscarPor_CANAC_Email(usuario.Id_Anac, usuario.Email); codAnacLancto == "YES" {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro já informado para este EMAIL ou CANAC! Por favor verifique. "))
		return
	}

	repositorio := repository.NovoRepositoDeUsuarios(db)
	usuario.ID, erro = repositorio.Criar(usuario)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios retorna todos usuarios
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	// nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuarios(db)

	usuarios, erro := repositorio.BuscarTodos(1)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, usuarios)

}

// BuscarUsuario retorna somente um usuario
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuarios(db)
	usuario, erro := repositorio.BuscarPorCANAC(usuarioID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	response.JSON(w, http.StatusOK, usuario)
}

// AtualizandoUsuario altera as informacoes de somente um usuario
func AtualizandoUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	usarioIDNoToken, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}
	if usuarioID != usarioIDNoToken {
		response.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar um usuário que não seja o seu"))
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var usuario models.Usuario

	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		fmt.Println("!")
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}
	if erro = usuario.Preparar("edicao"); erro != nil {

		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()
	// fmt.Println(usuarioID, usuario)
	repositorio := repository.NovoRepositoDeUsuarios(bd)
	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// AtualizarSenha permite alterar a senha de um usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	usuarioIDNoToken, erro := authentication.ExtrairUsuarioID(r)
	if erro != nil {
		response.Erro(w, http.StatusUnauthorized, erro)
		return
	}

	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if usuarioIDNoToken != usuarioID {
		response.Erro(w, http.StatusForbidden, errors.New("Não é possível atualizar a senha de uma usuário diferente do seu"))
	}
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	var senha models.Senha
	if erro = json.Unmarshal(corpoRequisicao, &senha); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeUsuarios(db)
	senhaSalvaNodatabase, erro := repositorio.BuscarSenha(usuarioID)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	if erro = security.VerificarSenha(senhaSalvaNodatabase, senha.Atual); erro != nil {
		response.Erro(w, http.StatusInternalServerError, errors.New("A Senha atual não codiz com a que está atual"))
		return
	}

	senhaComHash, erro := security.Hash(senha.Nova)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = repositorio.AtualizarSenha(usuarioID, string(senhaComHash)); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	response.JSON(w, http.StatusNoContent, nil)

}

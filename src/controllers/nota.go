package controllers

import (
	"encoding/json"
	"errors"
	"flight/src/database"
	"flight/src/models"
	"flight/src/repository"
	"flight/src/response"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarNota insere uma nota no banco de dados
func CriarNota(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if !json.Valid([]byte(corpoRequest)) {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro enviado para servidor não compativel com o sistema. "))
		return
	}

	var nota models.Nota

	if erro = json.Unmarshal(corpoRequest, &nota); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = nota.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioDuplicado := repository.NovoRepositoDeNota(db)
	if notaDuplicada := repositorioDuplicado.BuscarPorDescricao_Duplicado(nota.Nome); notaDuplicada == "YES" {
		response.Erro(w, http.StatusBadRequest, errors.New(" Esta NOTA já foi informada! Operação Bloqueada. "))
		return
	}

	repositorio := repository.NovoRepositoDeNota(db)
	nota.ID, erro = repositorio.Criar(nota)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, nota)
}

// AtualizandoNota altera as informacoes de somente de uma nota
func AtualizandoNota(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tipoVooID, erro := strconv.ParseUint(parametros["notaId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if !json.Valid([]byte(corpoRequisicao)) {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro enviado para servidor não compativel com o sistema. "))
		return
	}

	var nota models.Nota

	if erro = json.Unmarshal(corpoRequisicao, &nota); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = nota.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repositorio := repository.NovoRepositoDeNota(bd)
	if erro = repositorio.Atualizar(tipoVooID, nota); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// BuscarNotasretorna todos as notas
func BuscarNotas(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeNota(db)

	nota, erro := repositorio.BuscarTodos(1)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, nota)
}

// BuscarNota retorna uma nota por descricao
func BuscarNota(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	tipoVooPesquisa := parametros["descricao"]

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeNota(db)

	nota, erro := repositorio.BuscarNota(tipoVooPesquisa)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, nota)

}

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

// CriarTipoVoo insere um tipo de voo no banco de dados
func CriarTipoVoo(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if !json.Valid([]byte(corpoRequest)) {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro enviado para servidor não compativel com o sistema. "))
		return
	}

	var tipoVoo models.TipoVoo

	if erro = json.Unmarshal(corpoRequest, &tipoVoo); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = tipoVoo.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioDuplicado := repository.NovoRepositoDeTipoVoo(db)
	if tipoVooDuplicada := repositorioDuplicado.BuscarPorDescricao_Duplicado(tipoVoo.Nome); tipoVooDuplicada == "YES" {
		response.Erro(w, http.StatusBadRequest, errors.New(" TIPO DE VOO já informado! Operação Bloqueada. "))
		return
	}

	repositorio := repository.NovoRepositoDeTipoVoo(db)
	tipoVoo.ID, erro = repositorio.Criar(tipoVoo)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, tipoVoo)
}

// AtualizandoTipoVoo altera as informacoes de somente de um tipo de voo
func AtualizandoTipoVoo(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tipoVooID, erro := strconv.ParseUint(parametros["tipoVooId"], 10, 64)
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

	var tipoVoo models.TipoVoo

	if erro = json.Unmarshal(corpoRequisicao, &tipoVoo); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = tipoVoo.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repositorio := repository.NovoRepositoDeTipoVoo(bd)
	if erro = repositorio.Atualizar(tipoVooID, tipoVoo); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// BuscarTipoVoo retorna todos os tipos de voo
func BuscarTipoVoos(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeTipoVoo(db)

	tipoVoo, erro := repositorio.BuscarTodos(1)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, tipoVoo)
}

// BuscarTipoVoo retorna um tipo de voo por nome
func BuscarTipoVoo(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	tipoVooPesquisa := parametros["descricao"]

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeTipoVoo(db)

	tipoVoo, erro := repositorio.BuscarTipoVoo(tipoVooPesquisa)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, tipoVoo)

}

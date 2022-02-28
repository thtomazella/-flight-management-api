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

// CriarTipoInstrucao insere um tipo de instrucao no banco de dados
func CriarTipoInstrucao(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if !json.Valid([]byte(corpoRequest)) {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro enviado para servidor não compativel com o sistema. "))
		return
	}

	var tipoInstrucao models.TipoInstrucao

	if erro = json.Unmarshal(corpoRequest, &tipoInstrucao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = tipoInstrucao.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioDuplicado := repository.NovoRepositoDeTipoInstrucao(db)
	if tipoVooDuplicada := repositorioDuplicado.BuscarPorDescricao_Duplicado(tipoInstrucao.Nome); tipoVooDuplicada == "YES" {
		response.Erro(w, http.StatusBadRequest, errors.New(" TIPO DE INSTRUCAO já informado! Operação Bloqueada. "))
		return
	}

	repositorio := repository.NovoRepositoDeTipoInstrucao(db)
	tipoInstrucao.ID, erro = repositorio.Criar(tipoInstrucao)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, tipoInstrucao)
}

// AtualizandoTipoInstrucao altera as informacoes de somente de um tipo de instrucao
func AtualizandoTipoInstrucao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	tipoVooID, erro := strconv.ParseUint(parametros["tipoInstrucaoId"], 10, 64)
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

	var tipoInstrucao models.TipoInstrucao

	if erro = json.Unmarshal(corpoRequisicao, &tipoInstrucao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = tipoInstrucao.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repositorio := repository.NovoRepositoDeTipoInstrucao(bd)
	if erro = repositorio.Atualizar(tipoVooID, tipoInstrucao); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// BuscarTipoInstrucoes retorna todos os tipos de instrucao
func BuscarTipoInstrucoes(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeTipoInstrucao(db)

	tipoVoo, erro := repositorio.BuscarTodos(1)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, tipoVoo)
}

// BuscarTipoInstrucao retorna um tipo de instrucao por nome
func BuscarTipoInstrucao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	tipoVooPesquisa := parametros["descricao"]

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeTipoInstrucao(db)

	tipoVoo, erro := repositorio.BuscarTipoInstrucao(tipoVooPesquisa)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, tipoVoo)

}

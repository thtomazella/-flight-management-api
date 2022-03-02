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

// CriarMovimentacao insere um tipo de voo no banco de dados
func CriarMovimentacao(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if !json.Valid([]byte(corpoRequest)) {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro enviado para servidor não compativel com o sistema. "))
		return
	}

	var movimentacao models.Movimentacao

	if erro = json.Unmarshal(corpoRequest, &movimentacao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = movimentacao.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeMovimentacao(db)
	movimentacao.ID, erro = repositorio.Criar(movimentacao)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, movimentacao)
}

// AtualizandoMovimentacao altera as informacoes de somente de uma movimentacao
func AtualizandoMovimentacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	movimentacaoID, erro := strconv.ParseUint(parametros["movimentacaoId"], 10, 64)
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

	var movimentacao models.Movimentacao

	if erro = json.Unmarshal(corpoRequisicao, &movimentacao); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = movimentacao.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repositorio := repository.NovoRepositoDeMovimentacao(bd)
	if erro = repositorio.Atualizar(movimentacaoID, movimentacao); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// BuscarMovimentacoes retorna todos as movimentacoes
func BuscarMovimentacoes(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeMovimentacao(db)

	tipoVoo, erro := repositorio.BuscarTodos(1)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, tipoVoo)
}

// BuscarMovimentacao retorna uma movimentacao por descricao
func BuscarMovimentacao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	movimentacaoPesquisa := parametros["descricao"]

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeMovimentacao(db)

	movimentacao, erro := repositorio.BuscarMovimentacao(movimentacaoPesquisa)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, movimentacao)

}

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

// CriarTipoEquipamento insere um tipo de equipamento no banco de dados
func CriarTipoEquipamento(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if !json.Valid([]byte(corpoRequest)) {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro enviado para servidor não compativel com o sistema. "))
		return
	}

	var tipoEquipamento models.TipoEquipamento

	if erro = json.Unmarshal(corpoRequest, &tipoEquipamento); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = tipoEquipamento.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioDuplicado := repository.NovoRepositoDeTipoEquipamento(db)
	if tipoEquipDuplicada := repositorioDuplicado.BuscarPorDescricao_Duplicado(tipoEquipamento.Nome); tipoEquipDuplicada == "YES" {
		response.Erro(w, http.StatusBadRequest, errors.New(" TIPO DE EQUIPAMENTO já informado! Operação Bloqueada. "))
		return
	}

	repositorio := repository.NovoRepositoDeTipoEquipamento(db)
	tipoEquipamento.ID, erro = repositorio.Criar(tipoEquipamento)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, tipoEquipamento)
}

// AtualizandoTipoEquipamento altera as informacoes de somente de um tipo de equipamento
func AtualizandoTipoEquipamento(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	aeronaveID, erro := strconv.ParseUint(parametros["tipoEquioamentoId"], 10, 64)
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

	var tipoequipamento models.TipoEquipamento

	if erro = json.Unmarshal(corpoRequisicao, &tipoequipamento); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = tipoequipamento.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repositorio := repository.NovoRepositoDeTipoEquipamento(bd)
	if erro = repositorio.Atualizar(aeronaveID, tipoequipamento); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// BuscarTipoEquipamentos retorna todos os tipos de equipamentos
func BuscarTipoEquipamentos(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeTipoEquipamento(db)

	tipoEquipamento, erro := repositorio.BuscarTodos(1)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, tipoEquipamento)
}

// BuscarTipoEquipamento retorna um tipo de equipamento por nome
func BuscarTipoEquipamento(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	tipoEquipPesquisa := parametros["descricao"]

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeTipoEquipamento(db)

	tipoEquipamento, erro := repositorio.BuscarTipoEquipamento(tipoEquipPesquisa)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, tipoEquipamento)

}

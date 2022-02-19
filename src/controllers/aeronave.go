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

// CriarAeronave insere uma aeronave no banco de dados
func CriarAeronave(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	if !json.Valid([]byte(corpoRequest)) {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro enviado para servidor não compativel com o sistema. "))
		return
	}

	var aeronave models.Aeronave

	if erro = json.Unmarshal(corpoRequest, &aeronave); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = aeronave.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repositorioDuplicado := repository.NovoRepositoDeAeronave(db)
	if codAeronaveDuplicada := repositorioDuplicado.BuscarPorPrefixo_Duplicado(aeronave.Prefixo); codAeronaveDuplicada == "YES" {
		response.Erro(w, http.StatusBadRequest, errors.New(" PREFIXO já informado! Operação Bloqueada. "))
		return
	}

	repositorio := repository.NovoRepositoDeAeronave(db)
	aeronave.ID, erro = repositorio.Criar(aeronave)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, aeronave)
}

// AtualizandoAeronave altera as informacoes de somente uma aeronave
func AtualizandoAeronave(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	aeronaveID, erro := strconv.ParseUint(parametros["aeronaveid"], 10, 64)
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

	var aeronave models.Aeronave

	if erro = json.Unmarshal(corpoRequisicao, &aeronave); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = aeronave.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repositorio := repository.NovoRepositoDeAeronave(bd)
	if erro = repositorio.Atualizar(aeronaveID, aeronave); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// BuscarAeronaves retorna todas as aeronaves
func BuscarAeronaves(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeAeronave(db)

	aeronaves, erro := repositorio.BuscarTodos(1)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, aeronaves)
}

// BuscarAeronave retorna a aeronave por prefixo
func BuscarAeronave(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	prefixoAeronave := parametros["prefixo"]

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeAeronave(db)

	aeronave, erro := repositorio.BuscarAeronavePorPrefixo(prefixoAeronave)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, aeronave)

}

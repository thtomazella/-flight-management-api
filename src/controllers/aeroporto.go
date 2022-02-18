package controllers

import (
	"encoding/json"
	"flight/src/database"
	"flight/src/models"
	"flight/src/repository"
	"flight/src/response"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CriarAeroporto insere um aeroporto no banco de dados
func CriarAeroporto(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)

	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var aeroporto models.Aeroporto
	if erro = json.Unmarshal(corpoRequest, &aeroporto); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = aeroporto.Preparar("cadastro"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	/* -- COLOCAR DUPLICIDADE DE AEROPORTO
	repositorioDuplicado := repository.NovoRepositoDeUsuarios(db)
	if codAnacLancto, _ := repositorioDuplicado.BuscarPor_CANAC_Email(usuario.Id_Anac, usuario.Email); codAnacLancto == "YES" {
		response.Erro(w, http.StatusBadRequest, errors.New(" Registro já informado para este EMAIL ou CANAC! Por favor verifique. "))
		return
	}*/

	repositorio := repository.NovoRepositoDeAeroportos(db)
	aeroporto.ID, erro = repositorio.Criar(aeroporto)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusCreated, aeroporto)
}

// AtualizandoAeroporto altera as informacoes de somente um usuario
func AtualizandoAeroporto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	aeroportoID, erro := strconv.ParseUint(parametros["aeroportoId"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	var aeroporto models.Aeroporto

	if erro = json.Unmarshal(corpoRequisicao, &aeroporto); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = aeroporto.Preparar("edicao"); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bd, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer bd.Close()

	repositorio := repository.NovoRepositoDeAeroportos(bd)
	if erro = repositorio.Atualizar(aeroportoID, aeroporto); erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusNoContent, nil)

}

// BuscarAeroportos retorna todos aeroportos
func BuscarAeroportos(w http.ResponseWriter, r *http.Request) {

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeAeroportos(db)

	usuarios, erro := repositorio.BuscarTodos(1)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, usuarios)
}

// BuscarAeroportosPorDescricao retorna todos aeroportos por descricao ou sigla
func BuscarAeroportosPorDescricao(w http.ResponseWriter, r *http.Request) {

	parametros := mux.Vars(r)
	descricaoAeroporto := parametros["descricao"]
	//descricaoAeroporto := strings.ToLower(r.URL.Query().Get("descricao"))

	db, erro := database.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}
	defer db.Close()

	repositorio := repository.NovoRepositoDeAeroportos(db)

	usuarios, erro := repositorio.BuscarTodosPorDescricaoSigla(descricaoAeroporto)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
	}

	response.JSON(w, http.StatusOK, usuarios)

}

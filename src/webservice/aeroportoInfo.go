package webservice

import (
	"encoding/json"
	"flight/src/config"
	"flight/src/response"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type API_InfoAeroporto struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    struct {
		Localidade     string `json:"localidade"`
		Nome           string `json:"nome"`
		Cidade         string `json:"cidade"`
		Lon            string `json:"lon"`
		Lat            string `json:"lat"`
		Localizacao    string `json:"localizacao"`
		Metar          string `json:"metar"`
		Data           string `json:"data"`
		Temperatura    string `json:"temperatura"`
		Ur             string `json:"ur"`
		Visibilidade   string `json:"visibilidade"`
		Teto           string `json:"teto"`
		Ceu            string `json:"ceu"`
		CondicoesTempo string `json:"condicoes_tempo"`
		TempoImagem    string `json:"tempoImagem"`
		Vento          string `json:"vento"`
	} `json:"data"`
}

func BuscaAPI_InfoAeroporto(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	descAeroporto := parametros["descricao"]

	url := fmt.Sprintf("https://api-redemet.decea.mil.br/aerodromos/info?api_key=%s&localidade=%s", config.SecretKey_Redemet, descAeroporto)

	resposta, erro := http.Get(url)
	if erro != nil {
		return
	}

	bodyBytes, _ := ioutil.ReadAll(resposta.Body)

	var api_info API_InfoAeroporto
	if erro := json.Unmarshal([]byte(bodyBytes), &api_info); erro != nil {
		return
	}

	response.JSON(w, http.StatusOK, api_info)
}

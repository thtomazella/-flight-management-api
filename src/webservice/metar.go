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

type API_Metar struct {
	Status  bool `json:"status"`
	Message int  `json:"message"`
	Data    struct {
		CurrentPage int `json:"current_page"`
		Data        []struct {
			IDLocalidade    string `json:"id_localidade"`
			ValidadeInicial string `json:"validade_inicial"`
			Mens            string `json:"mens"`
			Recebimento     string `json:"recebimento"`
		} `json:"data"`
	} `json:"data"`
}

func BuscaAPI_Metar(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	descAeroporto := parametros["descricao"]

	url := fmt.Sprintf("https://api-redemet.decea.mil.br/mensagens/metar/%s?api_key=%s", descAeroporto, config.SecretKey_Redemet)

	resposta, erro := http.Get(url)
	if erro != nil {
		return
	}

	bodyBytes, _ := ioutil.ReadAll(resposta.Body)

	var api_metar API_Metar
	if erro := json.Unmarshal([]byte(bodyBytes), &api_metar); erro != nil {
		return
	}

	response.JSON(w, http.StatusOK, api_metar)
}

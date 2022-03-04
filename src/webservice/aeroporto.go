package webservice

import (
	"encoding/json"
	"flight/src/config"
	"flight/src/response"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type API_Aeroporto struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    [][]string `json:"data"`
}

func BuscaAPI_Aeroporto(w http.ResponseWriter, r *http.Request) {
	url := fmt.Sprintf("https://api-redemet.decea.mil.br/aerodromos/status?api_key=%s", config.SecretKey_Redemet)


	resposta, erro := http.Get(url)
	if erro != nil {
		fmt.Println(erro)
		return
	}
	
	bodyBytes, _ := ioutil.ReadAll(resposta.Body)

	var api_aeroporto API_Aeroporto
	if erro := json.Unmarshal([]byte(bodyBytes), &api_aeroporto); erro != nil {
		fmt.Println(erro)
		log.Fatalln(erro)
		return
	}

	response.JSON(w, http.StatusOK, api_aeroporto)
}



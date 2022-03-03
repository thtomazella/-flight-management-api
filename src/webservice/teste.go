package webservice

import (
	"encoding/json"
	"flight/src/response"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Teste struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func BuscarWebService(w http.ResponseWriter, r *http.Request) {
	API := "https://jsonplaceholder.typicode.com/posts"
	// TOKEN := "abcdefg"

	resp, err := http.Get(API)
	// resp.Header.Set("Authentication", TOKEN)

	if err != nil {
		log.Fatalln(err)
	}

	var _teste []Teste

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	if erro := json.Unmarshal([]byte(bodyBytes), &_teste); erro != nil {
		fmt.Println(erro)
		return
	}

	//	bodyString := string(bodyBytes)
	// fmt.Printf("%+v\n", bodyString)

	response.JSON(w, http.StatusOK, _teste)
}

package rotas

import (
	"flight/src/webservice"
	"net/http"
)

var rotasWebService = []Rota{
	{ // Cadastro de uma aeronave
		URI:                "/testeweb",
		Metodo:             http.MethodGet,
		Funcao:             webservice.BuscarWebService,
		RequerAutenticacao: false,
	},
}

package rotas

import (
	"flight/src/webservice"
	"net/http"
)

var rotasWebService = []Rota{
	{ // Rota referente a API aeroporto
		URI:                "/api_aeroporto",
		Metodo:             http.MethodGet,
		Funcao:             webservice.BuscaAPI_Aeroporto,
		RequerAutenticacao: true,
	},
	{ // Cadastro de uma aeronave
		URI:                "/testeweb",
		Metodo:             http.MethodGet,
		Funcao:             webservice.BuscarWebService,
		RequerAutenticacao: false,
	},
}

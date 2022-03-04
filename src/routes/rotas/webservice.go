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
	{ // Rota referente a API Metar
		URI:                "/api_metar/{descricao}",
		Metodo:             http.MethodGet,
		Funcao:             webservice.BuscaAPI_Metar,
		RequerAutenticacao: true,
	},
	{ // Rota referente a API Informacao do Aeroporto
		URI:                "/api_aeroporto/{descricao}",
		Metodo:             http.MethodGet,
		Funcao:             webservice.BuscaAPI_InfoAeroporto,
		RequerAutenticacao: true,
	},

	{ // Cadastro de uma aeronave
		URI:                "/testeweb",
		Metodo:             http.MethodGet,
		Funcao:             webservice.BuscarWebService,
		RequerAutenticacao: false,
	},
}

package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasAeronave = []Rota{
	{ // Cadastro de uma aeronave
		URI:                "/aeronave",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarAeronave,
		RequerAutenticacao: true,
	},

	{ //Alterar uma aeronave por ID
		URI:                "/aeronave/{aeronaveid}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoAeronave,
		RequerAutenticacao: true,
	},

	{ // Busca Todas as Aeronaves
		URI:                "/aeronave",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarAeronaves,
		RequerAutenticacao: true,
	},

	{ // Busca Aeronave por filtro de prefixo
		URI:                "/aeronave/{prefixo}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarAeronave,
		RequerAutenticacao: true,
	},
}

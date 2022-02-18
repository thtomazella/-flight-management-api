package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasAeroporto = []Rota{
	{ // Cadastro do aeroporto
		URI:                "/aeroporto",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarAeroporto,
		RequerAutenticacao: true,
	},
	{ //Alterar o aeroporto por ID
		URI:                "/aeroporto/{aeroportoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoAeroporto,
		RequerAutenticacao: true,
	},

	{ // Busca Todos os Aeroportos
		URI:                "/aeroporto",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarAeroportos,
		RequerAutenticacao: true,
	},

	{ // Busca Aeroporto por filtro de nome ou Sigla
		URI:                "/aeroporto/{descricao}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarAeroportosPorDescricao,
		RequerAutenticacao: true,
	},
}

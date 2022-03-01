package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasNota = []Rota{
	{ // Cadastro da Nota
		URI:                "/nota",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarNota,
		RequerAutenticacao: true,
	},

	{ //Alterar a Nota por ID
		URI:                "/nota/{notaId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoNota,
		RequerAutenticacao: true,
	},

	{ // Busca Todas as Notas
		URI:                "/nota",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarNotas,
		RequerAutenticacao: true,
	},
	{ // Busca Notas por filtro de nome
		URI:                "/nota/{descricao}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarNota,
		RequerAutenticacao: true,
	},
}

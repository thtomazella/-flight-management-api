package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasMovimentacao = []Rota{
	{ // Cadastro de Movimentacao
		URI:                "/movimentacao",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarMovimentacao,
		RequerAutenticacao: true,
	},

	{ //Alterar o tipo de Movimentacao
		URI:                "/movimentacao/{movimentacaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoMovimentacao,
		RequerAutenticacao: true,
	},

	{ // Busca Todos as Movimentacao
		URI:                "/movimentacao",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarMovimentacao,
		RequerAutenticacao: true,
	},
	{ // Busca Movimentacao por descricao
		URI:                "/movimentacao/{descricao}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarMovimentacao,
		RequerAutenticacao: true,
	},
}

package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasTipoVoo = []Rota{
	{ // Cadastro do tipo de Voo
		URI:                "/tipoVoo",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTipoVoo,
		RequerAutenticacao: true,
	},

	{ //Alterar o tipo de Voo por ID
		URI:                "/tipoVoo/{tipoVooId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoTipoVoo,
		RequerAutenticacao: true,
	},

	{ // Busca Todos os tipo de Voo
		URI:                "/tipoVoo",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTipoVoos,
		RequerAutenticacao: true,
	},
	{ // Busca tipo de Voo por filtro de nome
		URI:                "/tipoVoo/{descricao}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTipoVoo,
		RequerAutenticacao: true,
	},
}

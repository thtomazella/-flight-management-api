package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasTipoInstrucao = []Rota{
	{ // Cadastro do tipo de Instrucao
		URI:                "/tipoInstrucao",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTipoInstrucao,
		RequerAutenticacao: true,
	},

	{ //Alterar o tipo de Instrucao por ID
		URI:                "/tipoInstrucao/{tipoInstrucaoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoTipoInstrucao,
		RequerAutenticacao: true,
	},

	{ // Busca Todos os tipo de Instrucao
		URI:                "/tipoInstrucao",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTipoInstrucoes,
		RequerAutenticacao: true,
	},
	{ // Busca tipo de Instrucao por filtro de nome
		URI:                "/tipoInstrucao/{descricao}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTipoInstrucao,
		RequerAutenticacao: true,
	},
}

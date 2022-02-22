package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasTipoEquipamento = []Rota{
	{ // Cadastro do tipo de equipamento
		URI:                "/tipoEquipamento",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarTipoEquipamento,
		RequerAutenticacao: true,
	},

	{ //Alterar o tipo de equipamento por ID
		URI:                "/tipoEquipamento/{tipoEquioamentoId}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizandoTipoEquipamento,
		RequerAutenticacao: true,
	},

	{ // Busca Todos os tipo de equipamento
		URI:                "/tipoEquipamento",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTipoEquipamentos,
		RequerAutenticacao: true,
	},
	{ // Busca tipo de equipamento por filtro de nome
		URI:                "/tipoEquipamento/{descricao}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTipoEquipamento,
		RequerAutenticacao: true,
	},
}

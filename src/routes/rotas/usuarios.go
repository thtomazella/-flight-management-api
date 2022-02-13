package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{ // Cadastro o usuario
		URI:                "/usuario",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{ // Busca Todos os Usuários
		URI:                "/usuarios",
		Metodo:             http.MethodGet,
		Funcao:             nil, // controllers.BuscarUsuarios,
		RequerAutenticacao: true,
	},

	{ // Buscar somente um usuario por ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodGet,
		Funcao:             nil, // controllers.BuscarUsuario,
		RequerAutenticacao: true,
	},

	{ //Alterar o usuario por ID
		URI:                "/usuarios/{usuarioId}",
		Metodo:             http.MethodPut,
		Funcao:             nil, // controllers.AtualizandoUsuario,
		RequerAutenticacao: true,
	},

	{
		URI:                "/usuarios/{usuarioId}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             nil, // controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}

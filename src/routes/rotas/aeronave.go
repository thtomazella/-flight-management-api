package rotas

import (
	"flight/src/controllers"
	"net/http"
)

var rotasAeronave = []Rota{
	{ // Cadastro do aeroporto
		URI:                "/aeronave",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarAeronave,
		RequerAutenticacao: true,
	},
}

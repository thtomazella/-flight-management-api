package rotas

import (
	"flight/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)
	rotas = append(rotas, rotasAeroporto...)
	rotas = append(rotas, rotasAeronave...)
	rotas = append(rotas, rotasTipoEquipamento...)
	rotas = append(rotas, rotasTipoVoo...)
	rotas = append(rotas, rotasTipoInstrucao...)
	rotas = append(rotas, rotasNota...)
	rotas = append(rotas, rotasMovimentacao...)

	for _, rota := range rotas {
		if rota.RequerAutenticacao {
			r.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else {

			r.HandleFunc(rota.URI, middlewares.Logger(rota.Funcao)).Methods(rota.Metodo)
		}
	}
	return r
}

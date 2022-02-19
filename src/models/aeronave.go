package models

import (
	"errors"
	"strings"
)

// Aeronave representa um estrutura da tabela da aeronave / bd
type Aeronave struct {
	ID        uint64  `json: "id, omitempty"`
	Nome      string  `json: "nome, omitempty"`
	Prefixo   string  `json: "prefixo, omitempty"`
	Custo     float64 `json: "custo"`
	Preco     float64 `json: "preco"`
	Inclusion string  `json: "inclusion,omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar o aeroporto recebido
func (aeronave *Aeronave) Preparar(etapa string) error {
	if erro := aeronave.validar(etapa); erro != nil {
		return erro
	}

	if erro := aeronave.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (aeronave *Aeronave) validar(etapa string) error {
	if aeronave.Nome == "" {
		return errors.New("O nome da aeronave é obrigatório e não pode estar em branco.")
	}
	if aeronave.Prefixo == "" {
		return errors.New("A Sigla da aeronave é obrigatório e não pode estar em branco.")
	}
	return nil
}

func (aeronave *Aeronave) fomatar(etapa string) error {
	aeronave.Nome = strings.TrimSpace(aeronave.Nome)
	aeronave.Prefixo = strings.TrimSpace(aeronave.Prefixo)
	return nil
}

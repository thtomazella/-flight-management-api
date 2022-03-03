package models

import (
	"errors"
	"strings"
)

// Aeroporto representa um estrutura da tabela do aeroporto / bd
type Aeroporto struct {
	ID        uint64 `json: "id, omitempty"`
	Nome      string `json: "nome, omitempty"`
	Sigla     string `json: "sigla, omitempty"`
	Inclusion string `json: "inclusion,omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar o aeroporto recebido
func (aeroporto *Aeroporto) Preparar(etapa string) error {
	if erro := aeroporto.validar(etapa); erro != nil {
		return erro
	}

	if erro := aeroporto.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (aeroporto *Aeroporto) validar(etapa string) error {
	if aeroporto.Nome == "" {
		return errors.New(" O nome do aeroporto é obrigatório e não pode estar em branco. ")
	}
	if aeroporto.Sigla == "" {
		return errors.New(" A Sigla é obrigatório e não pode estar em branco. ")
	}

	return nil
}

func (aeroporto *Aeroporto) fomatar(etapa string) error {
	aeroporto.Nome = strings.TrimSpace(aeroporto.Nome)
	aeroporto.Sigla = strings.TrimSpace(aeroporto.Sigla)
	return nil
}

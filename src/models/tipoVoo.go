package models

import (
	"errors"
	"strings"
)

// TipoVoo representa um estrutura da tabela da tipoVoo / bd
type TipoVoo struct {
	ID        uint64 `json: "id, omitempty"`
	Nome      string `json: "nome, omitempty"`
	Inclusion string `json: "inclusion,omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar o tipo de Voo recebido
func (tipoVoo *TipoVoo) Preparar(etapa string) error {
	if erro := tipoVoo.validar(etapa); erro != nil {
		return erro
	}

	if erro := tipoVoo.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (tipoVoo *TipoVoo) validar(etapa string) error {
	if tipoVoo.Nome == "" {
		return errors.New(" A descrição do Tipo de Voo é obrigatório e não pode estar em branco. ")
	}
	return nil
}

func (tipoVoo *TipoVoo) fomatar(etapa string) error {
	tipoVoo.Nome = strings.TrimSpace(tipoVoo.Nome)
	return nil
}

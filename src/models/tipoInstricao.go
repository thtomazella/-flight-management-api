package models

import (
	"errors"
	"strings"
)

// TipoInstrucao representa um estrutura da tabela da tipo_instrucao / bd
type TipoInstrucao struct {
	ID        uint64 `json: "id, omitempty"`
	Nome      string `json: "nome, omitempty"`
	Inclusion string `json: "inclusion,omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar o tipo de Voo recebido
func (tipoInstrucao *TipoInstrucao) Preparar(etapa string) error {
	if erro := tipoInstrucao.validar(etapa); erro != nil {
		return erro
	}

	if erro := tipoInstrucao.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (tipoInstrucao *TipoInstrucao) validar(etapa string) error {
	if tipoInstrucao.Nome == "" {
		return errors.New(" A descrição do Tipo de Voo é obrigatório e não pode estar em branco. ")
	}
	return nil
}

func (tipoInstrucao *TipoInstrucao) fomatar(etapa string) error {
	tipoInstrucao.Nome = strings.TrimSpace(tipoInstrucao.Nome)
	return nil
}

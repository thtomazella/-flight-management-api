package models

import (
	"errors"
	"strings"
)

// TipoEquipamento representa um estrutura da tabela da tipoequipamento / bd
type TipoEquipamento struct {
	ID        uint64 `json: "id, omitempty"`
	Nome      string `json: "nome, omitempty"`
	Inclusion string `json: "inclusion,omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar o tipo de equipamento recebido
func (tipoEquipamento *TipoEquipamento) Preparar(etapa string) error {
	if erro := tipoEquipamento.validar(etapa); erro != nil {
		return erro
	}

	if erro := tipoEquipamento.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (tipoEquipamento *TipoEquipamento) validar(etapa string) error {
	if tipoEquipamento.Nome == "" {
		return errors.New(" A descrição do Tipo de Equipamento é obrigatório e não pode estar em branco. ")
	}
	return nil
}

func (tipoEquipamento *TipoEquipamento) fomatar(etapa string) error {
	tipoEquipamento.Nome = strings.TrimSpace(tipoEquipamento.Nome)
	return nil
}

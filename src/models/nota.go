package models

import (
	"errors"
	"strings"
)

// TipoVoo representa um estrutura da tabela da tipoVoo / bd
type Nota struct {
	ID        uint64 `json: "id, omitempty"`
	Nome      string `json: "nome, omitempty"`
	Inclusion string `json: "inclusion,omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar a nota recebida
func (nota *Nota) Preparar(etapa string) error {
	if erro := nota.validar(etapa); erro != nil {
		return erro
	}

	if erro := nota.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (nota *Nota) validar(etapa string) error {
	if nota.Nome == "" {
		return errors.New(" A descrição da Nota é obrigatório e não pode estar em branco. ")
	}
	return nil
}

func (nota *Nota) fomatar(etapa string) error {
	nota.Nome = strings.TrimSpace(nota.Nome)
	return nil
}

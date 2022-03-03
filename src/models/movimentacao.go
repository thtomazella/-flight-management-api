package models

// Movimentacao representa um estrutura da tabela de movimentacao / bd
type Movimentacao struct {
	ID                uint64 `json: "id, omitempty"`
	TipoEquipamentoID uint64 `json: "tipoEquipamentoId, omitempty"`
	TipoEquipamento   string `json: "tipoEquipamento, omitempty"`
	TipoInstrucaoID   uint64 `json: "tipoInstrucaoId, omitempty"`
	TipoInstrucao     string `json: "tipoInstrucao, omitempty"`
	NotaID            uint64 `json: "notaId, omitempty"`
	Nota              string `json: "nota, omitempty"`
	TipoVooID         uint64 `json: "tipoVooId, omitempty"`
	TipoVoo           string `json: "tipoVoo, omitempty"`
	Inclusion         string `json: "inclusion,omitempty"`
}

// MovimentacaoRetorno representa um estrutura da tabela de movimentacao / bd
type MovimentacaoRetorno struct {
	ID uint64 `json: "id, omitempty"`

	TipoEquipamento struct {
		ID   string `json:"ID"`
		Nome string `json:"Nome"`
	} `json:"tipoEquipamento"`

	TipoInstrucao struct {
		ID   string `json:"ID"`
		Nome string `json:"Nome"`
	} `json:"tipoInstrucao"`

	Nota struct {
		ID   string `json:"ID"`
		Nome string `json:"Nome"`
	} `json:"Nota"`

	TipoVoo struct {
		ID   string `json:"ID"`
		Nome string `json:"Nome"`
	} `json:"tipoVoo"`

	Inclusion string `json: "inclusion,omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar a nota recebida
func (movimentacao *Movimentacao) Preparar(etapa string) error {
	if erro := movimentacao.validar(etapa); erro != nil {
		return erro
	}

	if erro := movimentacao.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (movimentacao *Movimentacao) validar(etapa string) error {
	/**	if movimentacao.TipoEquipamentoID == "" {
		return errors.New(" A descrição da Nota é obrigatório e não pode estar em branco. ")
	}
	*/
	return nil
}

func (movimentacao *Movimentacao) fomatar(etapa string) error {
	//movimentacao.TipoEquipamentoID = strings.TrimSpace(movimentacao.TipoEquipamentoID)
	return nil
}

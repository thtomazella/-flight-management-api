package repository

import (
	"database/sql"
	"flight/src/models"
	"fmt"
)

type Movimentacoes struct {
	db *sql.DB
}

// NovoRepositoDeMovimentacao cria um repositorio de movimentacao
func NovoRepositoDeMovimentacao(db *sql.DB) *Movimentacoes {
	return &Movimentacoes{db}
}

// Criar insere uma movimentacao no bando de dados
func (repositorio Movimentacoes) Criar(movimentacao models.Movimentacao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		`INSERT INTO movimentacao(tipoequipamento_id, tipoinstrucao_id, nota_id, tipovoo_id) 
		VALUES(?, ?, ?, ?)`,
	)

	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(movimentacao.TipoEquipamentoID,
		movimentacao.TipoInstrucaoID,
		movimentacao.NotaID,
		movimentacao.TipoVooID)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Atualizar altera as informações de um movimentacao no banco de dados
func (repositorio Movimentacoes) Atualizar(ID uint64, movimentacao models.Movimentacao) error {
	statement, erro := repositorio.db.Prepare(`
		UPDATE movimentacao SET tipoequipamento_id= ?, 
								tipoinstrucao_id= ?, 
								nota_id= ?, 
								tipovoo_id= ?
		WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(movimentacao.TipoEquipamentoID,
		movimentacao.TipoInstrucaoID,
		movimentacao.NotaID,
		movimentacao.TipoVooID,
		ID); erro != nil {
		return erro

	}
	return nil
}

// BuscarTodos retorna todas as Movimentacoes
func (repositorio Movimentacoes) BuscarTodos(codInt int) ([]models.MovimentacaoRetorno, error) {

	linhas, erro := repositorio.db.Query(`
	SELECT m.id,
		   tEquipamento.id as idEquip,
		   tEquipamento.nome as Equipamento,
		   tInstrucao.id as idInstrucao,
		   tInstrucao.nome as Instrucao,
		   nota.id as idnota,
		   nota.nome as Nota,
		   tVoo.id as idVoo,
		   tVoo.nome as Voo,
		   m.inclusion

	FROM movimentacao m
	  
	LEFT OUTER JOIN tipoequipamento tEquipamento
	    ON(m.tipoequipamento_id = tEquipamento.id)
	  		
	LEFT OUTER JOIN tipo_instrucao tInstrucao
		ON(m.tipoinstrucao_id = tInstrucao.id)

	LEFT OUTER JOIN nota
		ON(m.nota_id = nota.id)

	LEFT OUTER JOIN tipo_voo tVoo
	    ON(m.tipovoo_id = tVoo.id)

	WHERE m.id >= ? ORDER BY m.id`, codInt)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var movimentacoes []models.MovimentacaoRetorno

	for linhas.Next() {
		var movimentacao models.MovimentacaoRetorno

		if erro = linhas.Scan(&movimentacao.ID,
			&movimentacao.TipoEquipamento.ID,
			&movimentacao.TipoEquipamento.Nome,
			&movimentacao.TipoInstrucao.ID,
			&movimentacao.TipoInstrucao.Nome,
			&movimentacao.Nota.ID,
			&movimentacao.Nota.Nome,
			&movimentacao.TipoVoo.ID,
			&movimentacao.TipoVoo.Nome,
			&movimentacao.Inclusion); erro != nil {
			return nil, erro
		}
		movimentacoes = append(movimentacoes, movimentacao)
	}
	return movimentacoes, nil

}

// BuscarMovimentacao retorna os dados de movimentacoes por descricao
func (repositorio Movimentacoes) BuscarMovimentacao(_descricao string) ([]models.Movimentacao, error) {
	_descricao = fmt.Sprintf("%%%s%%", _descricao)
	linhas, erro := repositorio.db.Query(`
			SELECT id
			  FROM movimentacao where id = ?`, _descricao)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var movimentacoes []models.Movimentacao

	for linhas.Next() {
		var movimentacao models.Movimentacao

		if erro = linhas.Scan(&movimentacao.ID); erro != nil {
			return nil, erro
		}
		movimentacoes = append(movimentacoes, movimentacao)
	}
	return movimentacoes, nil
}

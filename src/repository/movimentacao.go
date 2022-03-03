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
	       t.id as idEquip,
		   t.nome as Equipamento
	  FROM movimentacao m
	  LEFT OUTER JOIN tipoequipamento t
	    ON(m.tipoequipamento_id = t.id)
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
			&movimentacao.TipoEquipamento.Nome); erro != nil {
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

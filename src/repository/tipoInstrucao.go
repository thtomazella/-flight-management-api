package repository

import (
	"database/sql"
	"flight/src/models"
	"fmt"
)

type TipoInstrucoes struct {
	db *sql.DB
}

// NovoRepositoDeTipoVoo cria um repositorio de tipo de instrucao
func NovoRepositoDeTipoInstrucao(db *sql.DB) *TipoInstrucoes {
	return &TipoInstrucoes{db}
}

// Criar insere um tipo de instrucao no bando de dados
func (repositorio TipoInstrucoes) Criar(tipoInstrucao models.TipoInstrucao) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		`INSERT INTO tipo_instrucao(nome) 
		VALUES(?)`,
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(tipoInstrucao.Nome)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Atualizar altera as informações de um tipo de instrucao no banco de dados
func (repositorio TipoInstrucoes) Atualizar(ID uint64, tipoInstrucao models.TipoInstrucao) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE tipo_instrucao SET nome= ?
		WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(
		tipoInstrucao.Nome,
		ID); erro != nil {
		return erro

	}
	return nil
}

//BuscarPorDescricao_Duplicado para bloqueio de registro duplicado
func (repositorio TipoInstrucoes) BuscarPorDescricao_Duplicado(_descricao string) string {
	linha, erro := repositorio.db.Query(`
			SELECT id
			  FROM tipo_instrucao where nome= ? `, _descricao)
	if erro != nil {
		return "NO"
	}
	defer linha.Close()

	if linha.Next() {
		return "YES"
	} else {
		return "NO"
	}
}

// BuscarTodos retorna todos os Tipos de instrucao
func (repositorio TipoInstrucoes) BuscarTodos(codInt int) ([]models.TipoInstrucao, error) {

	linhas, erro := repositorio.db.Query(`
	SELECT id, 
		   nome, 
		   inclusion
	  FROM tipo_instrucao
	  WHERE id >= ? ORDER BY nome`, codInt)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var tipoInstrucoes []models.TipoInstrucao

	for linhas.Next() {
		var tipoInstrucao models.TipoInstrucao

		if erro = linhas.Scan(&tipoInstrucao.ID,
			&tipoInstrucao.Nome,
			&tipoInstrucao.Inclusion); erro != nil {
			return nil, erro
		}
		tipoInstrucoes = append(tipoInstrucoes, tipoInstrucao)
	}
	return tipoInstrucoes, nil

}

// BuscarTipoVoo retorna os dados de um tipo de instrucao específico por descricao
func (repositorio TipoInstrucoes) BuscarTipoInstrucao(_descricao string) (models.TipoInstrucao, error) {
	_descricao = fmt.Sprintf("%%%s%%", _descricao)
	linha, erro := repositorio.db.Query(`
			SELECT id, nome
			  FROM tipo_instrucao where nome LIKE ?`, _descricao)
	if erro != nil {
		return models.TipoInstrucao{}, erro
	}
	defer linha.Close()

	var tipoInstrucao models.TipoInstrucao

	if linha.Next() {
		if erro = linha.Scan(&tipoInstrucao.ID,
			&tipoInstrucao.Nome); erro != nil {
			return models.TipoInstrucao{}, erro
		}
	}
	return tipoInstrucao, nil

}

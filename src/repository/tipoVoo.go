package repository

import (
	"database/sql"
	"flight/src/models"
)

type TipoVoos struct {
	db *sql.DB
}

// NovoRepositoDeTipoVoo cria um repositorio de tipo de voo
func NovoRepositoDeTipoVoo(db *sql.DB) *TipoVoos {
	return &TipoVoos{db}
}

// Criar insere um tipo de voo no bando de dados
func (repositorio TipoVoos) Criar(TipoVoo models.TipoVoo) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		`INSERT INTO tipo_voo(nome) 
		VALUES(?)`,
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(TipoVoo.Nome)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Atualizar altera as informações de um tipo de voo no banco de dados
func (repositorio TipoVoos) Atualizar(ID uint64, tipovoo models.TipoVoo) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE tipo_voo SET nome= ?
		WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(
		tipovoo.Nome,
		ID); erro != nil {
		return erro

	}
	return nil
}

//BuscarPorDescricao_Duplicado para bloqueio de registro duplicado
func (repositorio TipoVoos) BuscarPorDescricao_Duplicado(_descricao string) string {
	linha, erro := repositorio.db.Query(`
			SELECT id
			  FROM tipo_voo where nome= ? `, _descricao)
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

// BuscarTodos retorna todos os Tipos de Voo
func (repositorio TipoVoos) BuscarTodos(codInt int) ([]models.TipoVoo, error) {

	linhas, erro := repositorio.db.Query(`
	SELECT id, 
		   nome, 
		   inclusion
	  FROM tipo_voo
	  WHERE id >= ? ORDER BY nome`, codInt)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var tipoVoos []models.TipoVoo

	for linhas.Next() {
		var tipoVoo models.TipoVoo

		if erro = linhas.Scan(&tipoVoo.ID,
			&tipoVoo.Nome,
			&tipoVoo.Inclusion); erro != nil {
			return nil, erro
		}
		tipoVoos = append(tipoVoos, tipoVoo)
	}
	return tipoVoos, nil

}

// BuscarTipoVoo retorna os dados de um tipo de voo específico por descricao
func (repositorio TipoVoos) BuscarTipoVoo(_descricao string) (models.TipoVoo, error) {
	// _descricao = fmt.Sprintf("%%%s%%", _descricao)
	linha, erro := repositorio.db.Query(`
			SELECT id, nome
			  FROM tipo_voo where nome = ?`, _descricao)
	if erro != nil {
		return models.TipoVoo{}, erro
	}
	defer linha.Close()

	var tipoVoo models.TipoVoo

	if linha.Next() {
		if erro = linha.Scan(&tipoVoo.ID,
			&tipoVoo.Nome); erro != nil {
			return models.TipoVoo{}, erro
		}
	}
	return tipoVoo, nil

}

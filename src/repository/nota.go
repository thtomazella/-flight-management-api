package repository

import (
	"database/sql"
	"flight/src/models"
	"fmt"
)

type Notas struct {
	db *sql.DB
}

// NovoRepositoDeNota cria um repositorio de nota
func NovoRepositoDeNota(db *sql.DB) *Notas {
	return &Notas{db}
}

// Criar insere uma nota no bando de dados
func (repositorio Notas) Criar(Nota models.Nota) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		`INSERT INTO nota(nome) 
		VALUES(?)`,
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(Nota.Nome)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Atualizar altera as informações de uma nota no banco de dados
func (repositorio Notas) Atualizar(ID uint64, nota models.Nota) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE nota SET nome= ?
		WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(
		nota.Nome,
		ID); erro != nil {
		return erro

	}
	return nil
}

//BuscarPorDescricao_Duplicado para bloqueio de registro duplicado
func (repositorio Notas) BuscarPorDescricao_Duplicado(_descricao string) string {
	linha, erro := repositorio.db.Query(`
			SELECT id
			  FROM nota where nome= ? `, _descricao)
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

// BuscarTodos retorna todas as Notas
func (repositorio Notas) BuscarTodos(codInt int) ([]models.Nota, error) {

	linhas, erro := repositorio.db.Query(`
	SELECT id, 
		   nome, 
		   inclusion
	  FROM nota
	  WHERE id >= ? ORDER BY nome`, codInt)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var notas []models.Nota

	for linhas.Next() {
		var nota models.Nota

		if erro = linhas.Scan(&nota.ID,
			&nota.Nome,
			&nota.Inclusion); erro != nil {
			return nil, erro
		}
		notas = append(notas, nota)
	}
	return notas, nil

}

// BuscarNota retorna os dados de uma nota específico por descricao
func (repositorio Notas) BuscarNota(_descricao string) (models.Nota, error) {
	_descricao = fmt.Sprintf("%%%s%%", _descricao)
	linha, erro := repositorio.db.Query(`
			SELECT id, nome
			  FROM nota where nome LIKE ?`, _descricao)
	if erro != nil {
		return models.Nota{}, erro
	}
	defer linha.Close()

	var nota models.Nota

	if linha.Next() {
		if erro = linha.Scan(&nota.ID,
			&nota.Nome); erro != nil {
			return models.Nota{}, erro
		}
	}
	return nota, nil

}

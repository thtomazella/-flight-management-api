package repository

import (
	"database/sql"
	"flight/src/models"
	"fmt"
)

// Aeroporto representa um repositorio de aeroportos
type Aeroportos struct {
	db *sql.DB
}

// NovoRepositoDeAeroportos cria um repositorio de aeroportos
func NovoRepositoDeAeroportos(db *sql.DB) *Aeroportos {
	return &Aeroportos{db}
}

// Criar insere um aeroporto no bando de dados
func (repositorio Aeroportos) Criar(aeroporto models.Aeroporto) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		`INSERT INTO aeroporto(nome, sigla) 
		VALUES(?, ?)`,
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(aeroporto.Nome, aeroporto.Sigla)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Atualizar  altera as informações de um aeroporto no banco de dados
func (repositorio Aeroportos) Atualizar(ID uint64, aeroporto models.Aeroporto) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE aeroporto SET nome= ?, 
		sigla = ?
		WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(
		aeroporto.Nome,
		aeroporto.Sigla,
		ID); erro != nil {
		return erro

	}
	return nil
}

// BuscarTodos retorna todos os aeroportos
func (repositorio Aeroportos) BuscarTodos(codInt int) ([]models.Aeroporto, error) {

	linhas, erro := repositorio.db.Query(`
	SELECT id, 
		   nome, 
		   sigla,
		   inclusion
	  FROM aeroporto 
	  WHERE id >= ? ORDER BY nome`, codInt)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var aeroportos []models.Aeroporto

	for linhas.Next() {
		var aeroporto models.Aeroporto

		if erro = linhas.Scan(&aeroporto.ID,
			&aeroporto.Nome,
			&aeroporto.Sigla,
			&aeroporto.Inclusion); erro != nil {
			return nil, erro
		}
		aeroportos = append(aeroportos, aeroporto)
	}
	return aeroportos, nil

}

// BuscarPorANAC retorna os dados de um usuário específico por ID ANAC
func (repositorio Aeroportos) BuscarTodosPorDescricaoSigla(_descricao string) ([]models.Aeroporto, error) {
	_descricao = fmt.Sprintf("%%%s%%", _descricao)
	linhas, erro := repositorio.db.Query(`
			SELECT id, nome, sigla
			  FROM aeroporto where sigla like ? or nome like ?`, _descricao, _descricao)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var aeroportos []models.Aeroporto

	for linhas.Next() {
		var aeroporto models.Aeroporto
		if erro = linhas.Scan(&aeroporto.ID,
			&aeroporto.Nome,
			&aeroporto.Sigla); erro != nil {
			return nil, erro
		}
		aeroportos = append(aeroportos, aeroporto)
	}
	return aeroportos, nil

}

// BuscarAeroportoUnico retorna os dados de um aeroporto específico por ID
func (repositorio Aeroportos) BuscarAeroportoUnico(_ID uint64) (models.Aeroporto, error) {
	linha, erro := repositorio.db.Query(`
			SELECT id, nome, sigla
			  FROM aeroporto where id = ? `, _ID)
	if erro != nil {
		return models.Aeroporto{}, erro
	}
	defer linha.Close()

	var aeroporto models.Aeroporto

	if linha.Next() {
		if erro = linha.Scan(&aeroporto.ID,
			&aeroporto.Nome,
			&aeroporto.Sigla); erro != nil {
			return models.Aeroporto{}, erro
		}
	}
	return aeroporto, nil

}

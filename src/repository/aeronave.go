package repository

import (
	"database/sql"
	"flight/src/models"
)

// Aeronave representa um repositorio de aeronaves
type Aeronaves struct {
	db *sql.DB
}

// NovoRepositoDeAeronave cria um repositorio de aeronaves
func NovoRepositoDeAeronave(db *sql.DB) *Aeronaves {
	return &Aeronaves{db}
}

// Criar insere uma aeronave no bando de dados
func (repositorio Aeronaves) Criar(aeronave models.Aeronave) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		`INSERT INTO aeronave(nome, prefixo, preco, custo) 
		VALUES(?, ?, ?, ?)`,
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(aeronave.Nome, aeronave.Prefixo, aeronave.Preco, aeronave.Custo)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Atualizar altera as informações de uma aeronave no banco de dados
func (repositorio Aeronaves) Atualizar(ID uint64, aeronave models.Aeronave) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE aeronave SET nome= ?, 
		prefixo = ?,
		custo = ?,
		preco = ?
		WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(
		aeronave.Nome,
		aeronave.Prefixo,
		aeronave.Custo,
		aeronave.Preco,
		ID); erro != nil {
		return erro

	}
	return nil
}

//BuscarPorPrefixo para bloqueio de registro duplicado
func (repositorio Aeronaves) BuscarPorPrefixo_Duplicado(_prefixo string) string {
	linha, erro := repositorio.db.Query(`
			SELECT id
			  FROM aeronave where prefixo= ? `, _prefixo)
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

// BuscarTodos retorna todas as aeronaves
func (repositorio Aeronaves) BuscarTodos(codInt int) ([]models.Aeronave, error) {

	linhas, erro := repositorio.db.Query(`
	SELECT id, 
		   nome, 
		   prefixo,
		   custo,
		   preco,
		   inclusion
	  FROM aeronave
	  WHERE id >= ? ORDER BY nome`, codInt)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var aeronaves []models.Aeronave

	for linhas.Next() {
		var aeronave models.Aeronave

		if erro = linhas.Scan(&aeronave.ID,
			&aeronave.Nome,
			&aeronave.Prefixo,
			&aeronave.Custo,
			&aeronave.Preco,
			&aeronave.Inclusion); erro != nil {
			return nil, erro
		}
		aeronaves = append(aeronaves, aeronave)
	}
	return aeronaves, nil

}

// BuscarPorANAC retorna os dados de um usuário específico por ID ANAC
func (repositorio Aeronaves) BuscarAeronavePorPrefixo(_prefixo string) (models.Aeronave, error) {
	// _descricao = fmt.Sprintf("%%%s%%", _descricao)
	linha, erro := repositorio.db.Query(`
			SELECT id, nome, prefixo, custo, preco
			  FROM aeronave where prefixo = ?`, _prefixo)
	if erro != nil {
		return models.Aeronave{}, erro
	}
	defer linha.Close()

	var aeronave models.Aeronave

	if linha.Next() {
		if erro = linha.Scan(&aeronave.ID,
			&aeronave.Nome,
			&aeronave.Prefixo,
			&aeronave.Custo,
			&aeronave.Preco); erro != nil {
			return models.Aeronave{}, erro
		}
	}
	return aeronave, nil

}

package repository

import (
	"database/sql"
	"flight/src/models"
)

// Aeronave representa um repositorio de aeronaves
type Aeronave struct {
	db *sql.DB
}

// NovoRepositoDeAeronave cria um repositorio de aeronaves
func NovoRepositoDeAeronave(db *sql.DB) *Aeronave {
	return &Aeronave{db}
}

// Criar insere uma aeronave no bando de dados
func (repositorio Aeronave) Criar(aeronave models.Aeronave) (uint64, error) {
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

//BuscarPorPrefixo
func (repositorio Aeronave) BuscarPorPrefixo(_prefixo string) string {
	linha, erro := repositorio.db.Query(`
			SELECT id
			  FROM aeronave where prefixo= ? or email = ?`, _prefixo)
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

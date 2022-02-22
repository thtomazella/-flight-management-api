package repository

import (
	"database/sql"
	"flight/src/models"
)

type TipoEquipamentos struct {
	db *sql.DB
}

// NovoRepositoDeAeronave cria um repositorio de aeronaves
func NovoRepositoDeTipoEquipamento(db *sql.DB) *TipoEquipamentos {
	return &TipoEquipamentos{db}
}

// Criar insere um tipo de equipamento no bando de dados
func (repositorio TipoEquipamentos) Criar(tipoEquipamento models.TipoEquipamento) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		`INSERT INTO tipoequipamento(nome) 
		VALUES(?)`,
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(tipoEquipamento.Nome)

	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Atualizar altera as informações de um tipo de equipamento no banco de dados
func (repositorio TipoEquipamentos) Atualizar(ID uint64, tipoequipamento models.TipoEquipamento) error {
	statement, erro := repositorio.db.Prepare(`
	UPDATE tipoequipamento SET nome= ?
		WHERE id = ?`)

	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(
		tipoequipamento.Nome,
		ID); erro != nil {
		return erro

	}
	return nil
}

//BuscarPorDescricao_Duplicado para bloqueio de registro duplicado
func (repositorio TipoEquipamentos) BuscarPorDescricao_Duplicado(_descricao string) string {
	linha, erro := repositorio.db.Query(`
			SELECT id
			  FROM tipoequipamento where nome= ? `, _descricao)
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

// BuscarTodos retorna todos os Tipos de Equipamentos
func (repositorio TipoEquipamentos) BuscarTodos(codInt int) ([]models.TipoEquipamento, error) {

	linhas, erro := repositorio.db.Query(`
	SELECT id, 
		   nome, 
		   inclusion
	  FROM tipoequipamento
	  WHERE id >= ? ORDER BY nome`, codInt)

	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var tipoEquipamentos []models.TipoEquipamento

	for linhas.Next() {
		var tipoEquipamento models.TipoEquipamento

		if erro = linhas.Scan(&tipoEquipamento.ID,
			&tipoEquipamento.Nome,
			&tipoEquipamento.Inclusion); erro != nil {
			return nil, erro
		}
		tipoEquipamentos = append(tipoEquipamentos, tipoEquipamento)
	}
	return tipoEquipamentos, nil

}

// BuscarTipoEquipamento retorna os dados de um tipo equipamento específico por descricao
func (repositorio TipoEquipamentos) BuscarTipoEquipamento(_descricao string) (models.TipoEquipamento, error) {
	// _descricao = fmt.Sprintf("%%%s%%", _descricao)
	linha, erro := repositorio.db.Query(`
			SELECT id, nome
			  FROM tipoequipamento where nome = ?`, _descricao)
	if erro != nil {
		return models.TipoEquipamento{}, erro
	}
	defer linha.Close()

	var tipoEquipamento models.TipoEquipamento

	if linha.Next() {
		if erro = linha.Scan(&tipoEquipamento.ID,
			&tipoEquipamento.Nome); erro != nil {
			return models.TipoEquipamento{}, erro
		}
	}
	return tipoEquipamento, nil

}

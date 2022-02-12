package repository

import (
	"database/sql"
	"flight/src/models"
)

// Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositoDeUsuarios cria um repositorio de usuarios
func NovoRepositoDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// BuscarPorEmail busca usuário por email e retorna o seu ID e senha com HASH
func (repositorio Usuarios) BuscarPorEmail(email string) (models.User, error) {
	linha, erro := repositorio.db.Query("SELECT id, senha FROM users WHERE email = ?", email)
	if erro != nil {
		return models.User{}, nil
	}
	defer linha.Close()

	var usuario models.User
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Passsword); erro != nil {
			return models.User{}, nil
		}
	}
	return usuario, nil
}

package repository

import (
	"database/sql"
	"flight/src/models"
	"fmt"
)

// Usuarios representa um repositorio de usuarios
type Usuarios struct {
	db *sql.DB
}

// NovoRepositoDeUsuarios cria um repositorio de usuarios
func NovoRepositoDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no bando de dados
func (repositorio Usuarios) Criar(usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO usuario(nome, type_user, nick, id_anac, email, senha) VALUES(?, ?, ?, ?, ?, ?)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Type_User, usuario.Nick, usuario.Id_Anac, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}

// Buscar retorna todos os usuarios de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]models.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) //%nomeOuNick%

	linhas, erro := repositorio.db.Query(
		"SELECT id, nome, nick, id_Anac, email, inclusion FROM usuario where nome LIKE ? OR nick LIKE ?", nomeOuNick, nomeOuNick,
	)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Id_Anac,
			&usuario.Email,
			&usuario.Inclusion); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil

}

// Buscar retorna todos os usuarios de nome ou nick
func (repositorio Usuarios) BuscarTodos(codInt int) ([]models.Usuario, error) {

	linhas, erro := repositorio.db.Query("SELECT id, nome, nick, id_Anac, email, inclusion FROM usuario WHERE id >= ? ORDER BY nome", codInt)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var usuarios []models.Usuario
	for linhas.Next() {
		var usuario models.Usuario

		if erro = linhas.Scan(&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Id_Anac,
			&usuario.Email,
			&usuario.Inclusion); erro != nil {
			return nil, erro
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil

}

// BuscarPorID retorna os dados de um usuário específico
/*func (repositorio Usuarios) BuscarPorID(ID uint64) (models.Usuario, error) {
	linhas, erro := repositorio.db.Query("SELECT id, nome, id_anac, nick, email, inclusion from usuarios where id= ?", ID)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Id_Anac,
			&usuario.Email,
			&usuario.Inclusion); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}*/

// BuscarPorANAC retorna os dados de um usuário específico por ID ANAC
func (repositorio Usuarios) BuscarPorCANAC(ID uint64) (models.Usuario, error) {
	linhas, erro := repositorio.db.Query("SELECT id, nome, id_anac, nick, email, inclusion from usuarios where Id_Anac= ?", ID)
	if erro != nil {
		return models.Usuario{}, erro
	}
	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next() {
		if erro = linhas.Scan(&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Id_Anac,
			&usuario.Email,
			&usuario.Inclusion); erro != nil {
			return models.Usuario{}, erro
		}
	}
	return usuario, nil
}

// Atualizar  altera as informações de um usuário no banco de dados
func (repositorio Usuarios) Atualizar(ID uint64, usuario models.Usuario) error {
	fmt.Println(ID, usuario)
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET nome= ?, type_user = ?, nick = ?, email = ? WHERE id = ?")

	/*
	 cpf = ?,
	 id_anac = ?,
	 address = ?,
	 district = ?,
	 id_city = ?,
	 contact= ?,
	 cell = ?,
	 email = ?

	*/
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro := statement.Exec(
		usuario.Nome,
		usuario.Type_User,
		usuario.Nick,
		usuario.Email,
		ID); erro != nil {
		return erro

		/**
		  usuario.Cpf,
		  		usuario.Id_Anac,
		  		usuario.Address,
		  		usuario.District,
		  		usuario.Id_City,
		  		usuario.Contact,
		  		usuario.Cell,
		  		usuario.Email,
		*/

	}
	return nil

}

//  Deletar excluir as informações de um usuario no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, erro := repositorio.db.Prepare("DELETE FROM usuarios WHERE id= ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(ID); erro != nil {
		return erro
	}
	return nil
}

// BuscarPorEmail busca usuário por email e retorna o seu ID e senha com HASH
func (repositorio Usuarios) BuscarPorEmail(email string) (models.Usuario, error) {
	linha, erro := repositorio.db.Query("SELECT id, senha FROM usuarios WHERE email = ?", email)
	if erro != nil {
		return models.Usuario{}, nil
	}
	defer linha.Close()

	var usuario models.Usuario
	if linha.Next() {
		if erro = linha.Scan(&usuario.ID, &usuario.Senha); erro != nil {
			return models.Usuario{}, nil
		}
	}
	return usuario, nil
}

// BuscaSenha traz a senha de um usuário pelo ID
func (repositorio Usuarios) BuscarSenha(usuarioID uint64) (string, error) {
	linha, erro := repositorio.db.Query("SELECT senha FROM usuarios WHERE id= ?", usuarioID)
	if erro != nil {
		return "", erro
	}
	defer linha.Close()

	var usuario models.Usuario

	if linha.Next() {
		if erro = linha.Scan(&usuario.Senha); erro != nil {
			return "", erro
		}
	}
	return usuario.Senha, nil
}

// AtualizarSenha altera a senha de um usuario
func (repositorio Usuarios) AtualizarSenha(usuarioID uint64, senha string) error {
	statement, erro := repositorio.db.Prepare("UPDATE usuarios SET senha =? WHERE id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(senha, usuarioID); erro != nil {
		return erro
	}
	return nil
}

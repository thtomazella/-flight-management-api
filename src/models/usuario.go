package models

import (
	"errors"
	"flight/src/security"
	"strings"

	"github.com/badoux/checkmail"
)

// User representa um usuario utilizando o software
type Usuario struct {
	ID        uint64 `json: "id,omitempty"`
	Nome      string `json: "nome,omitempty"`
	Type_User string `json: "type_user,omitempty"`
	Nick      string `json: "nick,omitempty"`
	Cpf       string `json: "cpf"`
	Id_Anac   string `json: "id_anac,omitempty"`
	//	Cma       time.Time `json: "cma"`
	Address   string `json: "address"`
	Number    string `json: "number"`
	District  string `json: "district"`
	Id_City   string `json: "id_city"`
	Contact   string `json: "contact"`
	Cell      string `json: "cell"`
	Email     string `json: "email,omitempty"`
	Senha     string `json: "senha,omitempty"`
	Inclusion string `json: "inclusion,omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar o usuário recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco.")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco.")
	}
	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco.")
	}

	if usuario.Id_Anac == "" {
		return errors.New("O CANAC é obrigatório e não pode estar em branco")
	}
	if usuario.Type_User == "" {
		return errors.New("O Tipo de Usuário é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email inserido é inválido. Tente novamente!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco.")
	}
	return nil
}

func (usuario *Usuario) fomatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}

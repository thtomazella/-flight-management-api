package models

import (
	"errors"
	"flight/src/security"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User representa um usuario utilizando o software
type User struct {
	ID        uint64    `json: "id, omitempty"`
	Name      string    `json: "name, omitempty"`
	Type_User string    `json: "type_user, omitempty"`
	Nick      string    `json: "nick, omitempty"`
	Cpf       string    `json: "cpf"`
	Id_Anac   string    `json: "id_anac, omitempty"`
	Cma       time.Time `json: "cma"`
	Address   string    `json: "address"`
	Number    string    `json: "number"`
	District  string    `json: "district"`
	Id_City   string    `json: "id_city"`
	Contact   string    `json: "contact"`
	Cell      string    `json: "cell"`
	Email     string    `json: "email, omitempty"`
	Passsword string    `json: "password, omitempty"`
	Inclusion string    `json: "inclusion, omitempty"`
}

// Prepara vai chamar is métodos para validar e formatar o usuário recebido
func (_user *User) Preparar(etapa string) error {
	if erro := _user.validar(etapa); erro != nil {
		return erro
	}

	if erro := _user.fomatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *User) validar(etapa string) error {
	if usuario.Name == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Id_Anac == "" {
		return errors.New("O CANAC é obrigatório e não pode estar em branco")
	}
	if usuario.Type_User == "" {
		return errors.New("O Tipo de Usuário é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email inserido é inválido. Tente novamente")
	}
	if etapa == "cadastro" && usuario.Passsword == "" {
		return errors.New("A senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (usuario *User) fomatar(etapa string) error {
	usuario.Name = strings.TrimSpace(usuario.Name)
	usuario.Type_User = strings.TrimSpace(usuario.Type_User)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Cpf = strings.TrimSpace(usuario.Cpf)
	usuario.Id_Anac = strings.TrimSpace(usuario.Id_Anac)
	usuario.Email = strings.TrimSpace(usuario.Email)
	usuario.Address = strings.TrimSpace(usuario.Address)
	usuario.Number = strings.TrimSpace(usuario.Number)
	usuario.District = strings.TrimSpace(usuario.District)
	usuario.Id_City = strings.TrimSpace(usuario.Id_City)
	usuario.Contact = strings.TrimSpace(usuario.Contact)
	usuario.Cell = strings.TrimSpace(usuario.Cell)

	if etapa == "cadastro" {
		senhaComHash, erro := security.Hash(usuario.Passsword)
		if erro != nil {
			return erro
		}
		usuario.Passsword = string(senhaComHash)
	}
	return nil
}

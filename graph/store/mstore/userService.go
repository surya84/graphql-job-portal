package mstore

import (
	"fmt"
	custommodels "graph/customModels"
	"graph/graph/model"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *Conn) AddUser(input *model.NewUser) (*model.Users, error) {

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return &model.Users{}, fmt.Errorf("generating password hash: %w", err)
	}
	user := custommodels.Users{

		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashedPass),
	}
	tx := s.db.Create(&user)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
		return &model.Users{}, tx.Error
	}
	return &model.Users{
		Name:  input.Name,
		Email: input.Email,
	}, nil
}

// func (s *Conn) FetchUser(email string) ()

// func (s *Conn) Login(email string, Password string) (*model.Users, error) {

// 	var u []custommodels.Users

// 	val,err:=FetchUser(email)
// 	tx=s.db.Where("email")

// 	err := bcrypt.CompareHashAndPassword([]byte(.PasswordHash), []byte(password))
// 	if err != nil {
// 		return
// 	}
// }

package userservices

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

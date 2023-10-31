package mstore

import (
	custommodels "graph/customModels"
	"graph/graph/model"

	"log"
)

func (s *Conn) AddCompany(input *model.NewCompany) (*model.Company, error) {
	com := custommodels.Company{
		Name:      input.Name,
		Location:  input.Location,
		CompanyID: input.CompanyID,
	}
	tx := s.db.Create(&com)
	if tx.Error != nil {
		log.Fatalln(tx.Error)
		return nil, tx.Error
	}
	return &model.Company{
		Name:      input.Name,
		Location:  input.Location,
		CompanyID: input.CompanyID,
	}, nil
}

func (s *Conn) DisplayCompany() ([]*model.Company, error) {

	var com = make([]*model.Company, 0, 10)
	err := s.db.Find(&com).Error

	if err != nil {
		return nil, nil
	}

	return com, nil

}

func (s *Conn) CompanyByID(id int) (*model.Company, error) {

	var c model.Company

	tx := s.db.Where("ID=?", id)

	err := tx.Find(&c).Error

	if err != nil {
		return nil, err
	}

	return &c, nil
}

package store

import (
	"graph/graph/model"
)

type Storer interface {
	AddUser(user *model.NewUser) (*model.Users, error)
	AddCompany(com *model.NewCompany) (*model.Company, error)
	DisplayCompany() ([]*model.Company, error)

	CompanyByID(id int) (*model.Company, error)

	CreateJob(input *model.NewJob) (*model.Job, error)

	GetAllJobs() ([]*model.Job, error)

	GetJobByID(id int) (*model.Job, error)

	GetAllJobsInCompany(id int) ([]*model.Job, error)
}

type Store struct {
	Storer
}

func NewStore(storer Storer) Store {
	return Store{Storer: storer}
}

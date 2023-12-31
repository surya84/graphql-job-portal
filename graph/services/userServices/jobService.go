package userservices

import (
	custommodels "graph/customModels"
	"graph/graph/model"
)

func (s *Conn) CreateJob(input *model.NewJob) (*model.Job, error) {

	job := &custommodels.Job{
		Title:       input.Name,
		Description: input.Description,
		CompanyID:   input.CompanyID,
	}

	tx := s.db.Create(&job)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &model.Job{
		Title:       input.Name,
		Description: input.Description,
		CompanyID:   input.CompanyID,
	}, nil
}

func (s *Conn) GetAllJobs() ([]*model.Job, error) {

	var job []*model.Job

	tx := s.db.Find(&job)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return job, nil
}

func (s *Conn) GetJobByID(id int) (*model.Job, error) {

	var job *model.Job

	tx := s.db.Where("ID=?", id)

	err := tx.Find(&job).Error

	if err != nil {
		return nil, err
	}

	return job, nil
}

func (s *Conn) GetAllJobsInCompany(id int) ([]*model.Job, error) {

	var job []*model.Job

	tx := s.db.Where("ID=?", id)

	err := tx.Find(&job).Error

	if err != nil {
		return nil, nil
	}
	return job, nil
}

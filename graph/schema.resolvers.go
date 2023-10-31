package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.40

import (
	"context"
	"fmt"
	"graph/graph/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.Users, error) {
	// We hash the user's password for storage in the database.
	// hashedPass, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	return &model.Users{}, fmt.Errorf("generating password hash: %w", err)
	// }
	// user := &model.Users{

	// 	Name:     input.Name,
	// 	Email:    input.Email,
	// 	Password: string(hashedPass),
	// }

	return r.S.AddUser(&input)
}

// CreateCompany is the resolver for the createCompany field.
func (r *mutationResolver) CreateCompany(ctx context.Context, input model.NewCompany) (*model.Company, error) {
	//panic(fmt.Errorf("not implemented: CreateCompany - createCompany"))

	return r.S.AddCompany(&input)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*model.Users, error) {
	panic(fmt.Errorf("not implemented: Login - login"))

	//return r.S.Login(email, password)
}

// CreateJob is the resolver for the createJob field.
func (r *mutationResolver) CreateJob(ctx context.Context, input model.NewJob) (*model.Job, error) {
	//panic(fmt.Errorf("not implemented: CreateJob - createJob"))

	return r.S.CreateJob(&input)
}

// FindUserByEmail is the resolver for the findUserByEmail field.
func (r *queryResolver) FindUserByEmail(ctx context.Context, email string) (*model.Users, error) {
	panic(fmt.Errorf("not implemented: FindUserByEmail - findUserByEmail"))
}

// DisplayCompany is the resolver for the displayCompany field.
func (r *queryResolver) DisplayCompany(ctx context.Context) ([]*model.Company, error) {
	//panic(fmt.Errorf("not implemented: DisplayCompany - displayCompany"))

	return r.S.DisplayCompany()
}

// CompanyByID is the resolver for the companyById field.
func (r *queryResolver) CompanyByID(ctx context.Context, id int) (*model.Company, error) {
	//panic(fmt.Errorf("not implemented: CompanyByID - companyById"))

	return r.S.CompanyByID(id)
}

// GetAllJobs is the resolver for the getAllJobs field.
func (r *queryResolver) GetAllJobs(ctx context.Context) ([]*model.Job, error) {
	//panic(fmt.Errorf("not implemented: GetAllJobs - getAllJobs"))

	return r.S.GetAllJobs()
}

// GetJobByID is the resolver for the getJobById field.
func (r *queryResolver) GetJobByID(ctx context.Context, id int) (*model.Job, error) {
	//panic(fmt.Errorf("not implemented: GetJobByID - getJobById"))

	return r.S.GetJobByID(id)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Company(ctx context.Context) (*model.Company, error) {
	panic(fmt.Errorf("not implemented: FindUserByEmail - findUserByEmail"))
}
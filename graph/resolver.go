package graph

import (
	"graph/graph/store"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen
type Resolver struct {
	S store.Store
}

type Resolvers struct {
	DB *gorm.DB
}


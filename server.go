package main

import (
	"graph/graph"
	"graph/graph/store"
	"graph/graph/store/mstore"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	dsn := "host=localhost user=postgres password=admin dbname=job-portal-graphql port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	//db, err := gorm.Open("postgres", "user=postgres dbname=postgres sslmode=disable")

	if err != nil {
		panic(err)
	}

	service, err := mstore.NewService(db)
	if err != nil {
		log.Fatalln(err)
	}
	err = service.AutoMigrate()
	if err != nil {
		log.Fatalln(err)
	}

	s := store.NewStore(service)
	// r := &graph.Resolvers{DB: db}
	// srv := NewExecutableSchema(Config{Resolvers: r})

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{S: s}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

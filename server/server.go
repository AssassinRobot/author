package server

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/AssassinRobot/author/graph"
	"github.com/AssassinRobot/author/internal/repository"
	"github.com/vektah/gqlparser/v2/ast"
)

func InitializeServer(port string,authorRepo repository.AuthorRepository,bookRepo repository.BookRepository,genre repository.GenreRepository,languageRepo repository.LanguageRepository) {
	srv := handler.New(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			AuthorRepo:authorRepo,
			BookRepo: bookRepo,
			GenreRepo: genre,
			LanguageRepo: languageRepo,
		}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

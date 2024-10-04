package graphql

import (
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"

	"log"
)

type Book struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Genre  string `json:"genre"`
	Author string `json:"author"`
}

var books = []Book{}

var bookType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Book",
		Fields: graphql.Fields{
			"id":     &graphql.Field{Type: graphql.String},
			"name":   &graphql.Field{Type: graphql.String},
			"genre":  &graphql.Field{Type: graphql.String},
			"author": &graphql.Field{Type: graphql.String},
		},
	},
)

var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"book": &graphql.Field{
				Type: bookType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					id, ok := params.Args["id"].(string)
					if ok {
						for _, book := range books {
							if book.ID == id {
								return book, nil
							}
						}
					}
					return nil, nil
				},
			},
			"books": &graphql.Field{
				Type: graphql.NewList(bookType),
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					return books, nil
				},
			},
		},
	},
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"addBook": &graphql.Field{
			Type: bookType,
			Args: graphql.FieldConfigArgument{
				"name":   &graphql.ArgumentConfig{Type: graphql.String},
				"genre":  &graphql.ArgumentConfig{Type: graphql.String},
				"author": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				book := Book{
					ID:     uuid.New().String(),
					Name:   params.Args["name"].(string),
					Genre:  params.Args["genre"].(string),
					Author: params.Args["author"].(string),
				}
				books = append(books, book)
				return book, nil
			},
		},
	},
})

var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)

func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		log.Printf("ошибка выполнения запроса: %+v", result.Errors)
	}
	return result
}

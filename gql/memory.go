package gql

import "github.com/graphql-go/graphql"

func Swap() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Swap",
			Fields: graphql.Fields{
				"total":     &graphql.Field{Type: graphql.Int},
				"used":      &graphql.Field{Type: graphql.Int},
				"available": &graphql.Field{Type: graphql.Int},
			},
		},
	)
}

func Virtual() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Virtual",
			Fields: graphql.Fields{
				"total":     &graphql.Field{Type: graphql.Int},
				"used":      &graphql.Field{Type: graphql.Int},
				"available": &graphql.Field{Type: graphql.Int},
				"cached":    &graphql.Field{Type: graphql.Int},
			},
		},
	)
}

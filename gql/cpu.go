package gql

import "github.com/graphql-go/graphql"

func CPU() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "CPU",
			Fields: graphql.Fields{
				"model": &graphql.Field{Type: graphql.String},
				"cores": &graphql.Field{Type: graphql.Int},
			},
		},
	)
}

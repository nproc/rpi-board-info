package gql

import "github.com/graphql-go/graphql"

func LoadAverage() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "LoadAverage",
			Fields: graphql.Fields{
				"load1":  &graphql.Field{Type: graphql.Float},
				"load5":  &graphql.Field{Type: graphql.Float},
				"load15": &graphql.Field{Type: graphql.Float},
			},
		},
	)
}

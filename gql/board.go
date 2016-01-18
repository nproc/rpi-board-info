package gql

import "github.com/graphql-go/graphql"

func Board() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Board",
			Fields: graphql.Fields{
				"hostname": &graphql.Field{Type: graphql.String},
				"os":       &graphql.Field{Type: graphql.String},
				"platform": &graphql.Field{Type: graphql.String},
				"family":   &graphql.Field{Type: graphql.String},
				"version":  &graphql.Field{Type: graphql.String},
			},
		},
	)
}

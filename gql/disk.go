package gql

import "github.com/graphql-go/graphql"

func Disk() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "Disk",
			Fields: graphql.Fields{
				"device":     &graphql.Field{Type: graphql.String},
				"mountpoint": &graphql.Field{Type: graphql.String},
				"type":       &graphql.Field{Type: graphql.String},
				"total":      &graphql.Field{Type: graphql.Int},
				"used":       &graphql.Field{Type: graphql.Int},
				"available":  &graphql.Field{Type: graphql.Int},
			},
		},
	)
}

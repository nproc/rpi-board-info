package gql

import "github.com/graphql-go/graphql"

func CPUTemperature() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "CPUTemperature",
			Fields: graphql.Fields{
				"celsius":    &graphql.Field{Type: graphql.Float},
				"fahrenheit": &graphql.Field{Type: graphql.Float},
			},
		},
	)
}

package gql

import "github.com/graphql-go/graphql"

func Network() *graphql.Object {
	return graphql.NewObject(graphql.ObjectConfig{
		Name: "Network",
		Fields: graphql.Fields{
			"name":            &graphql.Field{Type: graphql.String},
			"hardwareAddress": &graphql.Field{Type: graphql.String},
			"ips":             &graphql.Field{Type: graphql.NewList(graphql.String)},
		},
	})
}

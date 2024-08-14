package main

import (
	"github.com/graphql-go/graphql"
)

// Define the User GraphQL object.
var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// Define the root query.
var rootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type:        userType,
				Description: "Get a single user by ID",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.Int,
					},
				},
				Resolve: getUserResolver,
			},
			"users": &graphql.Field{
				Type:        graphql.NewList(userType),
				Description: "Get a list of users",
				Resolve:     listUsersResolver,
			},
		},
	},
)

// Define the root mutation.
var rootMutation = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "RootMutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type:        userType,
				Description: "Create a new user",
				Args: graphql.FieldConfigArgument{
					"name": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: createUserResolver,
			},
			"updateUser": &graphql.Field{
				Type:        userType,
				Description: "Update an existing user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"email": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				},
				Resolve: updateUserResolver,
			},
			"deleteUser": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Delete a user",
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.Int),
					},
				},
				Resolve: deleteUserResolver,
			},
		},
	},
)

// Initialize the GraphQL schema.
var schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	},
)

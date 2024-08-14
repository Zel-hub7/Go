package main

import (
	"github.com/graphql-go/graphql"
)

// Resolver for getting a single user by ID.
func getUserResolver(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Resolver for getting a list of users.
func listUsersResolver(p graphql.ResolveParams) (interface{}, error) {
	var users []User
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Resolver for creating a new user.
func createUserResolver(p graphql.ResolveParams) (interface{}, error) {
	user := User{
		Name:  p.Args["name"].(string),
		Email: p.Args["email"].(string),
	}
	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// Resolver for updating an existing user.
func updateUserResolver(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}

	if name, ok := p.Args["name"].(string); ok {
		user.Name = name
	}

	if email, ok := p.Args["email"].(string); ok {
		user.Email = email
	}

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Resolver for deleting a user.
func deleteUserResolver(p graphql.ResolveParams) (interface{}, error) {
	id, _ := p.Args["id"].(int)
	if err := db.Delete(&User{}, id).Error; err != nil {
		return nil, err
	}
	return true, nil
}

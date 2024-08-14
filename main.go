package main

import (
	"encoding/json"
	"net/http"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/graphql-go/graphql"
)

var db *gorm.DB
var err error

// Initialize the database and migrate models.
func init() {
	dsn := "host=localhost user=postgres dbname=crud_db sslmode=disable password=1234"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}

	db.AutoMigrate(&User{})
}

// Start the HTTP server and handle GraphQL requests.
func main() {
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query string `json:"query"`
		}

		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
		})

		json.NewEncoder(w).Encode(result)
	})

	http.ListenAndServe(":8080", nil)
}

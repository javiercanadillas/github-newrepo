// newghr command uses go-github as a cli tool for
// creating new GitHub repositories.
// It takes the var GITHUB_AUTH_TOKEN from the environment
// and creates a new repo under the account associated with
// the token.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v41/github"
	"golang.org/x/oauth2"
)

var (
	name        = flag.String("name", "", "Repo name to be created.")
	description = flag.String("description", "", "Repo description.")
	private     = flag.Bool("private", false, "Check repo as private.")
)

func main() {
	flag.Parse()
	tokenEnvVar := "GITHUB_AUTH_TOKEN"
	token := os.Getenv(tokenEnvVar)

	// Check that basic flags are provided
	if token == "" {
		log.Fatalf("No token is present, please set the %v var properly", tokenEnvVar)
	}
	if *name == "" {
		log.Fatal("Repo name not provided, you must specify one.")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)

	// Create a new repo with data from flags
	r := &github.Repository{
		Name:        github.String(*name),
		Private:     github.Bool(*private),
		Description: github.String(*description),
	}

	// Create the repo, providing no GitHub Org
	repo, _, err := client.Repositories.Create(ctx, "", r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created new repo %v\n", repo.GetName())
}

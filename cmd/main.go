package main

import (
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BigOplO/GO_INTERVIEW/graph"
	"github.com/BigOplO/GO_INTERVIEW/internal/contentful"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "bread-info",
	Short: "Get bread information from Contentful API and save to DB",
}

func main() {
	contentfulCmd := &cobra.Command{
		Use:   "contentful",
		Short: "Fetch data from Contentful API",
		RunE: func(cmd *cobra.Command, args []string) error {
			return contentful.FetchAndSave()
		},
	}

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Fetch data from GraphQL server",

		Run: func(cmd *cobra.Command, args []string) {
			srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))
			http.Handle("/", playground.Handler("GraphQL playground", "/query"))
			http.Handle("/query", srv)
			http.ListenAndServe(":8080", nil)
		},
	}

	rootCmd.AddCommand(contentfulCmd)
	rootCmd.AddCommand(serverCmd)
	rootCmd.Execute()
}

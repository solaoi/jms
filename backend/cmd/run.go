package cmd

import (
	"path"
	"os"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/solaoi/jms/cmd/graph"
	"github.com/solaoi/jms/cmd/graph/generated"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "generate JSON with form and form-template",
	Long: `generate JSON with form and form-template.

This command launches our default browser
and gives us a simple solution to generate JSON.`,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(currentDir)
	}

	var jmsDir = path.Join(currentDir, ".jms")
	if f, err := os.Stat(jmsDir); os.IsNotExist(err) || !f.IsDir() {
		log.Fatal("we should run jms init")
	}

	var jmsDb = path.Join(jmsDir, "jms.db")
	db, err := gorm.Open("sqlite3", jmsDb)
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.GET("/", welcome())

	graphqlHandler := handler.NewDefaultServer(
		generated.NewExecutableSchema(
			generated.Config{Resolvers: &graph.Resolver{DB: db}},
		),
	)
	playgroundHandler := playground.Handler("GraphQL", "/query")

	e.POST("/query", func(c echo.Context) error {
		graphqlHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/playground", func(c echo.Context) error {
		playgroundHandler.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err = e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}

func welcome() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Welcome!")
	}
}

func init() {
	rootCmd.AddCommand(runCmd)
}

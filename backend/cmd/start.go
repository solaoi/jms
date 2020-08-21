package cmd

import (
	"runtime"
	"fmt"
	"os/exec"
	"path"
	"os"
	"log"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/spf13/cobra"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"github.com/solaoi/jms/cmd/graph"
	"github.com/solaoi/jms/cmd/graph/generated"

	"github.com/solaoi/jms/lib"
    "github.com/elazarl/go-bindata-assetfs"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/solaoi/jms/static"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "generate JSON with form and form-template",
	Long: `generate JSON with form and form-template.

This command launches our default browser
and gives us a simple solution to generate JSON.`,
	Run: start,
}

func start(cmd *cobra.Command, args []string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var jmsDir = path.Join(currentDir, ".jms")
	if f, err := os.Stat(jmsDir); os.IsNotExist(err) || !f.IsDir() {
		log.Fatal("we should run jms init at first")
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

	e.Use(lib.ServeRoot("/", NewAssets("out")))

	if b, err := cmd.Flags().GetBool("open"); err == nil && b {
		openbrowser("http://localhost:3000/template/add")
	}

	err = e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}

func NewAssets(root string) *assetfs.AssetFS {
    return &assetfs.AssetFS{
        Asset:     static.Asset,
        AssetDir:  static.AssetDir,
        AssetInfo: static.AssetInfo,
        Prefix:    root,
    }
}

func openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolP("open", "o", false, "open browser automatically")
}

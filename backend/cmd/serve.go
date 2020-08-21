package cmd

import (
	"os"
	"log"
	"path"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/solaoi/jms/lib"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "serve JSON on the current public directory",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: serve,
}

func serve(cmd *cobra.Command, args []string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var jmsDir = path.Join(currentDir, "public")
	if f, err := os.Stat(jmsDir); os.IsNotExist(err) || !f.IsDir() {
		log.Fatal("there is no public dir to serve...")
	}

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(lib.ServeRootSimple("/", http.Dir("public/")))
	err = e.Start(":3000")
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

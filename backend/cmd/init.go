package cmd

import (
	"path"
	"os"
	"log"
	"io/ioutil"

	"github.com/spf13/cobra"
	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose"

	_ "github.com/solaoi/jms/migrations"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		var jmsDir = path.Join(currentDir, ".jms")
		if f, err := os.Stat(jmsDir); os.IsNotExist(err) || !f.IsDir() {
			if err := os.Mkdir(jmsDir, 0777); err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal("we've already run jms init in the current dir")
		}

		var jmsDb = path.Join(jmsDir, "jms.db")
		db, err := sqlx.Connect("sqlite3", jmsDb)
		if err != nil {
			log.Fatal("Failed to connect to sqlite3")
		}
		defer db.Close()

		// Run migrations
		goose.SetDialect("sqlite3")
		// Use a temporary directory for goose.Up() - we don't have any .sql files
		// to run, it's all embedded in the binary
		tmpdir, err := ioutil.TempDir(jmsDir, "")
		if err != nil {
			log.Fatal(err)
		}
		defer os.RemoveAll(tmpdir)

		// Discard Goose's log output
		goose.SetLogger(log.New(ioutil.Discard, "", 0))

		err = goose.Up(db.DB, tmpdir)
		if err != nil {
			log.Fatalf("Error running database migrations: %v\n", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

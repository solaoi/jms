package cmd

import (
	"os"
	"log"
	"encoding/json"

	"github.com/spf13/cobra"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"

	"github.com/solaoi/jms/cmd/graph/model"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "export JSON to the public directory under the current.",
	Long: `export JSON to the public directory under the current.

This command exports JSON generated by the run command.
JSON is exported to the public directory under the current.
If we do not find the public directory, we make it!`,
	Run: export,
}

func export(cmd *cobra.Command, args []string) {
	db, err := gorm.Open("sqlite3", "jms.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	var template model.Template
	db.Find(&template, 1)

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(currentDir)
	}
	var publicDir = currentDir + "/public"
	if f, err := os.Stat(publicDir); os.IsNotExist(err) || !f.IsDir() {
		if err := os.Mkdir(publicDir, 0777); err != nil {
			log.Fatal(err)
		}
	}

	file, err := os.Create(publicDir + "/sample.json")
    if err != nil {
		log.Fatal(err)
    }
    defer file.Close()

	json, err := json.Marshal(template)
	if err != nil {
		log.Fatal(err)
    }
    file.Write(([]byte)(json))
}

func init() {
	rootCmd.AddCommand(exportCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exportCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exportCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

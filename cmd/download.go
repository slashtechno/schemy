package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"

	"github.com/charmbracelet/log"
	"github.com/slashtechno/schemy/internal"
	"github.com/slashtechno/schemy/pkg/airtable"
	"github.com/spf13/cobra"
)

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download baseId output",
	Short: "Download a schema from Airtable",
	Args: func(cmd *cobra.Command, args []string) error {
		// Ensure there are exactly two arguments
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}

		// Check if the base ID is valid via regex
		re, err := regexp.Compile(`^app[a-zA-Z0-9]{14}$`)
		if err != nil {
			return fmt.Errorf("failed to compile regex: %v; this is a bug", err)
		}
		if !re.MatchString(args[0]) {
			return fmt.Errorf("base ID %s is not valid", args[0])
		}
		log.Debug("verified base ID", "baseId", args[0])

		// TODO: the file shouldn't exist, but the directory should. This should be changed later
		if info, err := os.Stat(args[1]); errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("path %s does not exist", args[1])
		} else if info.IsDir() {
			return fmt.Errorf("path %s is a directory", args[1])
		}
		log.Debug("verified output path", "path", args[1])

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		baseId := args[0]
		output := args[1]
		log.Debug("output path", "path", output)
		schema, err := airtable.NewClient(internal.Viper.GetString("airtable-token")).GetBaseSchema(baseId)
		if err != nil {
			log.Fatal("failed to get base schema", "baseId", baseId, "error", err)
		}
		log.Info("got base schema", "baseId", baseId)
		prettyJsonBytes, err := json.MarshalIndent(schema, "", "  ")
		if err != nil {
			log.Fatal("failed to marshal schema to JSON", "error", err)
		}
		fmt.Println(string(prettyJsonBytes))
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// downloadCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// downloadCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

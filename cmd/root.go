package cmd

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/slashtechno/schemy/internal"
	"github.com/slashtechno/schemy/pkg/utils"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "schemy",
	Short: "Create Airtable schemas!",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(func() {
		utils.SetupLogger(internal.Viper.GetString("log-level"))
		log.Debug("logger set up")
	})

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file (default is $XDG_CONFIG_HOME/generate-ddg/config.yaml)")

	rootCmd.PersistentFlags().String("log-level", "", "Log level")
	internal.Viper.BindPFlag("log-level", rootCmd.PersistentFlags().Lookup("log-level"))
	internal.Viper.SetDefault("log-level", "info")

	rootCmd.PersistentFlags().String("airtable-token", "", "Airtable API token")
	internal.Viper.BindPFlag("airtable-token", rootCmd.PersistentFlags().Lookup("airtable-token"))
}

func initConfig() {

	utils.LoadConfig(
		internal.Viper,
		cfgFile,
		"schemy/config.yaml",
		log.Default(),
		false,
	)
}

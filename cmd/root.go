package cmd

import (
	"os"
	"qaecli/config"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qae",
	Short: "Q App Engine Client",
	Long:  `Q App Engine Client`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	rootCmd.AddCommand(appCmd, checkCmd)
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	// init rootCmd flags
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.qaecli.yaml)")

	// init viper config
	viper.SetDefault("server", ":8000")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".qaecli")
	}

	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file: %v", viper.ConfigFileUsed())
	}

	config.InitConfig()
}

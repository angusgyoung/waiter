package cmd

import (
	"os"

	"github.com/angusgyoung/waiter/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/spf13/viper"
)

var cfgFile string
var logLevel string

var rootCmd = &cobra.Command{
	Use:   "waiter",
	Short: "Runs your tasks so you don't have to.",
	Long: `Waiter is a task runner, that exeutes simple or complex
	tasks, with support for batching, retries and schedules.`,
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.waiter.yaml)")
	rootCmd.PersistentFlags().StringVar(&logLevel, "level", "", "log level")
}

func initConfig() {
	if logLevel != "" {
		level, err := logrus.ParseLevel(logLevel)

		if err != nil {
			internal.Log.WithField("level", logLevel).Error("Failed to determine log level from argument, defaulting to warn")
			internal.Log.SetLevel(logrus.WarnLevel)
		}

		internal.Log.SetLevel(level)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".waiter")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		internal.Log.WithField("Path", viper.ConfigFileUsed()).Debug("Read config file")
	}

	internal.LoadConfig()
}

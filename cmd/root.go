package cmd

import (
	"awesomeProject/cmd2"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {

	rootCmd.AddCommand(cmd2.Cmd)
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config2", "", "config file (default is $HOME/.cobra.yaml)")
}

func initConfig() {
	fmt.Println(cfgFile)
}

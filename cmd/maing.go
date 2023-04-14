package cmd

import (
	"fmt"
	"os"

	"github.com/minnek-digital-studio/monorepo-ctrl/pkg"
	"github.com/spf13/cobra"
)

var IgnoreCheckVersion = false

var rootCmd = &cobra.Command{
	Use:     "monorepo-ctrl [command]",
	Short:   "Monorepo Control for Husky",
	Version: "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || args[0] == "" {
			fmt.Println("Error: No command specified.")
			os.Exit(1)
		}

		pkg.Init(args[0])
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetVersionTemplate("cominnek {{.Version}}\n")
	rootCmd.VersionTemplate()
}
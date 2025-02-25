/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"manifest/pkg/makeManifest"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use: `init <branch>
  ex) manifest init dev`,
	Long: `Render values in yaml to Template Manifests`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetString("values")
		makeManifest.MakeManifest(v, args[0])
		fmt.Println("")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().String("values", "./values.yaml", "path for values.yaml")
}

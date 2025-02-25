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
<<<<<<< HEAD
	Use: `init <branch>
  ex) manifest init dev`,
	Short: "A brief description of your command",
	Long:  `Render values in yaml to Template Manifests`,

	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := cmd.Flags().GetString("values")
		makeManifest.MakeManifest(v, args[0])
		fmt.Println("")
=======
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		makeManifest.MakeManifest(args[0])
		fmt.Println("Sucessfully make manifest")
>>>>>>> refs/remotes/origin/main
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
<<<<<<< HEAD
	initCmd.Flags().String("values", "./values.yaml", "path for values.yaml")
=======
>>>>>>> refs/remotes/origin/main
}

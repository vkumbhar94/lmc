/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	ar    string
	acf   string
	cscr  string
	csccf string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(ar)
		if ar != "" {
			helm := os.Getenv("HELM_BIN")
			fmt.Println(helm)
		}
	},
}

func init() {
	configCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	generateCmd.Flags().StringVar(&ar, "argus-release", "argus", "Argus Helm Release Name")
	generateCmd.Flags().StringVar(&cscr, "csc-release", "collectorset-controller", "Collectorset Controller Helm Release Name")
	generateCmd.Flags().StringVar(&acf, "argus-conf-file", "argus-configuration.yaml", "Argus Configuration Yaml File Path")
	generateCmd.Flags().StringVar(&csccf, "csc-conf-file", "collectorset-controller-configuration.yaml", "Collectorset Controller Configuration Yaml File Path")
}

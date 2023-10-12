/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		if err := Init(); err != nil {
			fmt.Println("init func error")
			return
		}

	},
}

func init() {
	rootCmd.AddCommand(initCmd)

}

func Init() error {
	fd, err := os.Create(".env")
	if err != nil {
		return err
	}

	_, err = fd.Write([]byte("USER="))
	if err != nil {
		return err
	}
	return nil
}

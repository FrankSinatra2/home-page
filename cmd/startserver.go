/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// startserverCmd represents the startserver command
var startserverCmd = &cobra.Command{
	Use:   "startserver",
	Short: "Start Server",
	Long:  ``,
	Run:   startserver,
}

func init() {
	rootCmd.AddCommand(startserverCmd)
}

func startserver(cmd *cobra.Command, args []string) {

	website := http.FileServer(http.Dir("dist"))
	http.Handle("/", website)

	fmt.Println("Listening on :3030")
	http.ListenAndServe(":3030", nil)
}

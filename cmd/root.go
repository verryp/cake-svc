package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/verryp/cake-store-service/app/server"
)

func Run() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var RootCmd = &cobra.Command{
	Use:   "root",
	Short: "cake-store-service",
	Long:  "cake-store service",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
	},
}

func init() {
	cobra.OnInitialize()
}

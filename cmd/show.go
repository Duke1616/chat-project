package cmd

import (
	"chat-project/client"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "查看登陆用户",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		client.GetClientList(remoteServerHost)
	},
}

func init() {
	rootCmd.AddCommand(showCmd)
	showCmd.Flags().StringVarP(&remoteServerHost, "remove-server-host", "s", "localhost:8080", "Remote server host where you want to join chat e.g 10.11.12.13:8080, default is localhost")
}

package commands

import (
	webserver "blog-go/embed/web-server"
	"blog-go/pkg/cmdroot"
	"blog-go/pkg/db"
	"github.com/spf13/cobra"
)
func Execute() {
	cmdroot.InitCommand("blog","blog-server")
	cmdroot.AddCommand(newServe())
	cmdroot.Execute()
}
func newServe() *cobra.Command {
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Start server",
		RunE:  startServer,
	}
	return serverCmd
}
func startServer(cmd *cobra.Command, args []string) error {
	db.InitMySQL()
	cmd.Println("Server starting ...")
	webserver.StartWebServer()
	cmdroot.WaitSignal()
	cmd.Println("Server stopping...")
	return nil
}
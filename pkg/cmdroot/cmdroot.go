package cmdroot

import (
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"
)
var (
	programName string
	cmdRoot  *cobra.Command
)

 func init() {
	 file, _ := exec.LookPath(os.Args[0])
	 _, programName = filepath.Split(file)
	 cmdRoot=getRootCommand()
}
func getRootCommand() *cobra.Command {
	root := &cobra.Command{
		Use:   programName,
		Short: "A brief description of command were not set",
		Long:  "A longer description were not set",
	}
	return root
}
func InitCommand(short, long string) {
	cmdRoot.Short = short
	cmdRoot.Long = long
}
// AddCommand add child command,
func AddCommand(cmd *cobra.Command) {
	cmdRoot.AddCommand(cmd)
}
// Execute command
func Execute() {
	rand.Seed(time.Now().UnixNano())
	if c, err := cmdRoot.ExecuteC(); err != nil {
		cmdRoot.Println(c.UsageString())
		os.Exit(-1)
	}
}
func WaitSignal() {
	sigChannel := make(chan os.Signal, 1)
	signal.Notify(sigChannel, os.Interrupt, os.Kill, syscall.SIGTERM)
	for {
		stop := false
		select {
		case <-sigChannel:
			stop = true
		}
		if stop {
			break
		}
	}
}

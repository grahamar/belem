package version

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/grahamar/belem/root"
)

// Command config.
var Command = &cobra.Command{
	Use:              "version",
	Short:            "Print version of belem",
	PersistentPreRun: root.PreRunNoop,
	Run:              run,
}

// Initialize.
func init() {
	root.Register(Command)
}

// Run command.
func run(c *cobra.Command, args []string) {
	fmt.Printf(root.PrintBelem())
	fmt.Printf("Version %s\n", root.Version)
	fmt.Printf("Build %s - %s\n", root.Commit, root.BuildTime)
}

package root

import (
	"bytes"

	"github.com/apex/log"

	"github.com/arsham/rainbow/rainbow"
	"github.com/mbndr/figlet4go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// Register `cmd`.
func Register(cmd *cobra.Command) {
	Command.AddCommand(cmd)
}

// PrintBelem returns BELÉM figlet rainbow
func PrintBelem() string {
	ascii := figlet4go.NewAsciiRender()
	renderStr, _ := ascii.Render("BELEM")

	r := bytes.NewReader([]byte(renderStr))
	w := new(bytes.Buffer)
	rb := rainbow.Light{
		Reader: r,
		Writer: w,
		Seed:   15, // 33 = Blue-Green, 15 = Orange-Purple
	}
	rb.Paint()

	return w.String()
}

// Command represents the base command when called without any subcommands
var Command = &cobra.Command{
	Use:               "belem",
	Short:             PrintBelem(),
	PersistentPreRunE: preRun,
}

func init() {
	f := Command.PersistentFlags()

	f.StringVar(&cfgFile, "config", "", "config file (default is $HOME/.belem.yaml)")
}

// PreRunNoop noop for other commands.
func PreRunNoop(c *cobra.Command, args []string) {
}

// preRun sets up global tasks used for most commands, some use PreRunNoop
// to remove this default behaviour.
func preRun(c *cobra.Command, args []string) error {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Search config in home directory with name ".belem.yaml".
		viper.AddConfigPath("$HOME")
		viper.SetConfigName(".belem")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	if err != nil {
		log.Warnf("Error reading config: %s", err)
	}

	return nil
}

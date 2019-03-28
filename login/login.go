package login

import (
	"errors"

	"github.com/apex/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/grahamar/belem/root"
)

// aws profile.
var profile string

// broker url.
var brokerURL string

// localhost callback port.
var port string

// example output.
const example = `
    Login with default AWS profile
    $ belem login
    Login with a specific AWS profile
    $ belem login -p <profile>`

// Command config.
var Command = &cobra.Command{
	Use:     "login [options]",
	Short:   "Obtain temporary AWS credentials",
	Example: example,
	RunE:    run,
}

// Initialize.
func init() {
	root.Register(Command)

	viper.SetDefault("broker", "https://broker.grhodes.io")
	viper.SetDefault("port", "8765")

	f := Command.Flags()
	f.StringVarP(&profile, "profile", "", "default", "Set AWS profile")
	f.StringVarP(&brokerURL, "broker", "b", viper.GetString("broker"), "Set AWS credentials broker URL")
	f.StringVarP(&port, "port", "p", viper.GetString("port"), "Set localhost callback port")

	viper.BindPFlag("broker", f.Lookup("broker"))
	viper.BindPFlag("port", f.Lookup("port"))
	viper.BindEnv("profile", "AWS_PROFILE")
}

// Run command.
func run(c *cobra.Command, args []string) error {
	log.Info("Waiting on browser callback...")
	if brokerURL == "" || port == "" {
		return errors.New("You must supply port and broker url")
	}
	return GetCredentials(profile, brokerURL, port)
}

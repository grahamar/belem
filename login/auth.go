package login

import (
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os/user"
	"time"

	"github.com/apex/log"
	"github.com/skratchdot/open-golang/open"
	ini "gopkg.in/ini.v1"
)

const awsCredentialsPath = ".aws/credentials"

func shutdown(srv *http.Server) {
	time.Sleep(500 * time.Millisecond)
	srv.Shutdown(nil)
}

func captureCredentials(profile string, awsBrokerURL string, srv *http.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("access_key_id")
		secretKey := r.URL.Query().Get("secret_access_key")
		sessionToken := r.URL.Query().Get("session_token")

		usr, _ := user.Current()
		awsCredentials := fmt.Sprintf("%s/%s", usr.HomeDir, awsCredentialsPath)
		cfg, err := ini.Load(awsCredentials)
		if err != nil {
			log.WithError(err).Fatalf("Exception loading AWS credentials file: %s", awsCredentials)
		}

		cfg.Section(profile).Key("aws_access_key_id").SetValue(key)
		cfg.Section(profile).Key("aws_secret_access_key").SetValue(secretKey)
		// apparently the different AWS SDKs either use "session_token" or "security_token", so set both
		cfg.Section(profile).Key("aws_session_token").SetValue(sessionToken)
		cfg.Section(profile).Key("aws_security_token").SetValue(sessionToken)
		cfg.SaveTo(awsCredentials)

		go shutdown(srv)

		log.Info("Successfully stored temporary AWS credentials!")
		http.Redirect(w, r, fmt.Sprintf("%s/success", awsBrokerURL), 301)
	}
}

// GetCredentials command.
func GetCredentials(profile string, awsBrokerURL string, port string) error {
	selfHost := fmt.Sprintf("localhost:%s", port)

	srv := &http.Server{Addr: selfHost}
	l, err := net.Listen("tcp", selfHost)
	if err != nil {
		return err
	}

	http.HandleFunc("/", captureCredentials(profile, awsBrokerURL, srv))

	uri, err := url.Parse(awsBrokerURL)
	if err != nil {
		return err
	}

	parameters := url.Values{}
	parameters.Add("callback_uri", fmt.Sprintf("http://%s", selfHost))
	uri.RawQuery = parameters.Encode()

	err = open.Start(uri.String())
	if err != nil {
		return err
	}

	srv.Serve(l)

	return nil
}

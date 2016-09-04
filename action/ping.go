package action

import (
	"fmt"
	"github.com/slapers/go-fastcgi-client"
	"github.com/urfave/cli"
	"strings"
)

func PingAction(c *cli.Context) error {

	fcgi, err := fcgiclient.New(c.String("host"), c.Int("port"))

	if err != nil {
		msg := fmt.Sprintf("Error connecting to fastcgi (%s:%v): %v", c.String("host"), c.Int("port"), err)
		return cli.NewExitError(msg, 1)
	}

	ret, err := fcgi.Request(buildRequestEnv(c.String("uri")), "")

	if err != nil {
		fmt.Printf("err: %v", err)
		msg := fmt.Sprintf("Error performing request %s: %v", c.String("uri"), err)
		return cli.NewExitError(msg, 1)
	}

	if !strings.Contains(string(ret), c.String("response")) {
		msg := fmt.Sprintf("Output did not contain expected response: '%s'\n", c.String("response"))
		msg += fmt.Sprintf("actual response: \n%s", string(ret))
		return cli.NewExitError(msg, 1)
	}

	fmt.Println("OK")
	return nil
}

func buildRequestEnv(uri string) map[string]string {

	return map[string]string{
		"REQUEST_METHOD":  "GET",
		"SCRIPT_NAME":     uri,
		"SCRIPT_FILENAME": uri,
		"SERVER_SOFTWARE": "go / ping-fastcgi",
		"REMOTE_ADDR":     "127.0.0.1",
		"SERVER_PROTOCOL": "HTTP/1.1",
	}
}

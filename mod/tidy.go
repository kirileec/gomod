package mod

import (
	"os"

	"github.com/hqpko/gosh"
	"github.com/urfave/cli"
)

const (
	defProxyURL  = "https://athens.azurefd.net"
	defProxyURL1 = "https://goproxy.io"
)

func ActionGoModTidy(c *cli.Context) error {
	if e := initEnvProxy(); e != nil {
		return e
	}
	return gosh.Run("go mod tidy")
}

func ActionGomodTidy1(c *cli.Context) error {
	if e := initEnvProxy1(); e != nil {
		return e
	}
	return gosh.Run("go mod tidy")
}

func initEnvProxy1() error {
	if os.Getenv("GOPROXY") == "" {
		return os.Setenv("GOPROXY", defProxyURL1)
	}
	return nil
}
func initEnvProxy() error {
	if os.Getenv("GOPROXY") == "" {
		return os.Setenv("GOPROXY", defProxyURL)
	}
	return nil
}

package main

import (
	"os"
	"time"

	"bak.server/pkg/utils/jsonx"
	"github.com/golang/glog"

	"github.com/koding/multiconfig"
)

type (
	Server struct {
		Name       string `required:"true"`
		Port       int    `default:"6060"`
		ID         int64
		Labels     []int
		Enabled    bool
		Users      []string
		Postgres   Postgres
		unexported string
		Interval   time.Duration
	}

	// Postgres holds Postgresql database related configuration
	Postgres struct {
		Enabled           bool
		Port              int      `required:"true" customRequired:"yes"`
		Hosts             []string `required:"true"`
		DBName            string   `default:"configdb"`
		AvailabilityRatio float64
		unexported        string
	}

	TaggedServer struct {
		Name     string `required:"true"`
		Postgres `structs:",flatten"`
	}
)

func main() {

	os.Setenv("SERVERCONFIG_NAME", "koding")
	os.Setenv("SERVERCONFIG_PORT", "6060")

	// Create a custom multi loader intance based on your needs.
	f := &multiconfig.FlagLoader{}
	e := &multiconfig.EnvironmentLoader{}
	fi := multiconfig.NewWithPath("config.toml")
	l := multiconfig.MultiLoader(f, e, fi)

	// Load configs into our s variable from the sources above
	s := &Server{}
	err := l.Load(s)
	if err != nil {
		panic(err)
	}
	glog.Info(jsonx.ToJson(s))
}

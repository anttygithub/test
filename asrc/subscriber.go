package subscriber

import (
	"test/config"
	"test/healthcheck"
	"test/rancherevents"

	log "github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
)

var ()

func Main(c *cli.Context) error {
	resultChan := make(chan error)
	conf := config.Conf(c)

	go func(rc chan error) {
		err := rancherevents.ConnectToEventStream(conf)
		log.Errorf("Rancher stream listener exited with error: %s", err)
		rc <- err
	}(resultChan)

	go func(rc chan error) {
		err := healthcheck.StartHealthCheck(conf.HealthCheckPort)
		log.Errorf("HealthCheck exit with error : %s", err)
		rc <- err
	}(resultChan)

	<-resultChan
	log.Info("Exiting...")

	return nil
}

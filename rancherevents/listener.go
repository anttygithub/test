package rancherevents

import (
	"github.com/anttygithub/test/config"
	reventhandlers "github.com/anttygithub/test/rancherevents/eventhandlers"
	revents "github.com/rancher/event-subscriber/events"
)

func ConnectToEventStream(conf config.Config) error {
	ehs := map[string]revents.EventHandler{
		"resource.change": reventhandlers.NewResourceChangeHandler().Handler,
	}
	router, err := revents.NewEventRouter("", 0, conf.CattleURL, conf.CattleAccessKey, conf.CattleSecretKey, nil, ehs, "", conf.WorkerCount, revents.DefaultPingConfig)
	if err != nil {
		return err
	}
	err = router.StartWithoutCreate(nil)
	return err
}

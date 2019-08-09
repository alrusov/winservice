package winservice

import (
	"time"

	"github.com/kardianos/service"

	"github.com/alrusov/log"
	"github.com/alrusov/misc"
)

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

// Service --
type Service struct {
	Config  *service.Config
	handler Handler
}

// Handler --
type Handler func(*service.Service)

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

// Interactive --
func (me *Service) Interactive() bool {
	return service.Interactive()
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

// Start --
func (me *Service) Start(serv service.Service) error {
	log.Message(log.INFO, "Service started")
	go me.handler(&serv)
	return nil
}

// Stop --
func (me *Service) Stop(serv service.Service) error {
	log.Message(log.INFO, "Service stopped")
	return nil
}

// Restart --
func (me *Service) Restart(serv service.Service) error {
	log.Message(log.INFO, "Service restarted")
	return nil
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

// Logger --
func (me *Service) Logger(errs chan<- error) (log.ServiceLogger, error) {
	return log.ServiceLogger{}, nil
}

// SystemLogger --
func (me *Service) SystemLogger(errs chan<- error) (log.ServiceLogger, error) {
	return log.ServiceLogger{}, nil
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

// Go --
func (me *Service) Go(command string, handler Handler) {
	me.handler = handler

	var serv service.Service
	var err error

	if serv, err = service.New(me, me.Config); err == nil {
		if command == "" {
			log.Message(log.INFO, `Service start initiated...`)
			err = serv.Run()
		} else {
			log.Message(log.INFO, `Service command "%s" processing...`, command)
			if command == "restart" {
				service.Control(serv, "stop")
				misc.Sleep(time.Duration(3) * time.Second)
				err = service.Control(serv, "start")
			} else {
				err = service.Control(serv, command)
			}
		}
	}

	if err != nil {
		log.Message(log.CRIT, "%s", err)
	}
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

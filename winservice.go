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

// New --
func New(config *service.Config, handler Handler) (*Service, error) {
	me := &Service{
		Config:  config,
		handler: handler,
	}
	return me, nil
}

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
func (me *Service) Go(command string) (err error) {
	var serv service.Service

	serv, err = service.New(me, me.Config)
	if err != nil {
		return
	}

	if command == "" {
		log.Message(log.INFO, `Service start initiated...`)
		err = serv.Run()
		return
	}

	log.Message(log.INFO, `Service command "%s" processing...`, command)

	if command == "restart" {
		// workaround...
		service.Control(serv, "stop")
		misc.Sleep(time.Duration(3) * time.Second)
		err = service.Control(serv, "start")
	} else {
		err = service.Control(serv, command)
	}

	return
}

//--------------------------------------------------------------------------------------------------------------------------------------------------------------//

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/kardianos/service"

	//"golang.org/x/sys/windows/svc/eventlog"
)

var logger service.Logger
type program struct{}

func (p *program) Start(s service.Service) error {
	// Inicialize o serviço aqui, se necessário.
	dt := time.Now()
	logger.Info(fmt.Sprintf("Iniciando o serviço: %s",  dt.Format("2006/01/02 15:04:05")))
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	// Pare o serviço aqui, se necessário.
	dt := time.Now()
	logger.Info(fmt.Sprintf("Parando o serviço: %s",  dt.Format("2006/01/02 15:04:05")))
	return nil
}

func (p *program) run() {
	for {
		dt := time.Now()
		logger.Info(fmt.Sprintf("mensagem: %s",  dt.Format("2006/01/02 15:04:05")))

		time.Sleep(5 * time.Second)
	}
}

func getConfig() *service.Config {
	svcConfig := &service.Config{
		Name:        "MeuServico",
		DisplayName: "Meu Serviço do Windows",
		Description: "Um exemplo de serviço do Windows em Go",
	}
	return svcConfig
}

func main() {
	prg := &program{}

	s, err := service.New(prg, getConfig())
	if err != nil {
		log.Fatal(err)
	}

	logger, err = s.Logger(nil)
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		logger.Error(err)
	}
}

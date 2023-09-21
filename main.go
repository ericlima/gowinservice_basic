package main

import (
	"log"
	"time"

	"github.com/kardianos/service"

	"golang.org/x/sys/windows/svc/eventlog"
)

var logger service.Logger
var eventLog *eventlog.Log
type program struct{}

func (p *program) Start(s service.Service) error {
	// Inicialize o serviço aqui, se necessário.
	logger.Info("Iniciando o serviço")
	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	// Pare o serviço aqui, se necessário.
	logger.Info("Parando o serviço")
	return nil
}

func (p *program) run() {
	for {
		// Coloque a lógica do seu serviço aqui.
		logger.Info("Serviço em execução...")
		eventLog.Info(1, "Mensagem de exemplo")

		time.Sleep(5 * time.Second)
	}
}

func main() {
	svcConfig := &service.Config{
		Name:        "MeuServico",
		DisplayName: "Meu Serviço do Windows",
		Description: "Um exemplo de serviço do Windows em Go",
	}

	eventLog, err := eventlog.Open(svcConfig.Name)
	if err != nil {
		log.Fatal(err)
	}
	defer eventLog.Close()

	prg := &program{}

	s, err := service.New(prg, svcConfig)
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

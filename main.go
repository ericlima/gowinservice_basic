package main

import (
	"log"
	"time"

	"github.com/kardianos/service"
)

var logger service.Logger

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
		time.Sleep(5 * time.Second)
	}
}

func main() {
	svcConfig := &service.Config{
		Name:        "MeuServico",
		DisplayName: "Meu Serviço do Windows",
		Description: "Um exemplo de serviço do Windows em Go",
	}

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

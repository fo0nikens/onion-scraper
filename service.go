package main

import (
	"context"
	"encoding/csv"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
)

const (
	onionV2 = 2
	onionV3 = 3
)

// Service handler
type Service struct {
	cfg    *Config
	client *Client
	csv    *csv.Writer
}

// Config data
type Config struct {
	TorSocksPort string
	Timeout      int
	Filename     string
	OnionVersion int
}

// Start the scraping
func (s *Service) Start() error {
	var err error

	// create our http proxy tor client
	s.client, err = NewClient(s.cfg.TorSocksPort, s.cfg.Timeout)
	if err != nil {
		return err
	}

	// create our storage layer
	s.csv, err = NewStore().CSV(s.cfg.Filename)
	if err != nil {
		return err
	}

	// gracefully stop service
	c1, cancel := context.WithCancel(context.Background())
	exitCh := make(chan struct{})
	go func(ctx context.Context) {
		for {
			s.run(ctx)
			select {
			case <-ctx.Done():
				exitCh <- struct{}{}
				return
			default:
			}
		}
	}(c1)

	// kill signal
	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, os.Interrupt)
	go func() {
		select {
		case <-signalCh:
			log.Info("Kill signal received, please wait to finish task")
			cancel()
			return
		}
	}()
	<-exitCh

	log.Info("Service stopped")
	return nil
}

func (s *Service) run(ctx context.Context) {
	var onionURL Address

	switch s.cfg.OnionVersion {
	case onionV2:
		onionURL = RandOnionV2()
		break

	case onionV3:
		onionURL = RandOnionV3()
		break

	default:
		onionURL = RandOnionV2()
	}

	log.Debugf("Scraping %s \n", onionURL.Addr())
	defer log.Debugf("Finished scraping %s\n", onionURL.Addr())

	resp, err := s.client.Request(onionURL)
	if err != nil {
		log.Debugf("%s is not available, err: %s", onionURL.Addr(), err.Error())
		return
	}

	err = s.csv.Write([]string{onionURL.Addr(), resp.Title, resp.Description})
	if err != nil {
		log.Errorf("error writing to file, err: %s \n", err.Error())
	}

	s.csv.Flush()
}

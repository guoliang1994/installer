package program

import (
	"fmt"
	"github.com/kardianos/service"
	"time"
)

type Server struct{}

func (t *Server) Start(s service.Service) error {
	_, err := s.Status()
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for {
			time.Sleep(time.Second * 1)
		}
	}()
	return nil
}
func (t *Server) Stop(s service.Service) error {
	return nil
}

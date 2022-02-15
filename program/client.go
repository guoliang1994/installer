package program

import (
	"fmt"
	"github.com/kardianos/service"
	"time"
)

type Client struct {
}

func (t *Client) Start(s service.Service) error {
	status, err := s.Status()
	fmt.Println(err)
	fmt.Println(status)
	go t.run()
	return nil
}
func (t *Client) run() {
	for {
		fmt.Println("nice to meet you")
		time.Sleep(time.Second * 1)
	}
}

func (t *Client) Stop(s service.Service) error {
	return nil
}

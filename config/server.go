package config

import "fmt"

type Server struct {
	Port       int `envconfig:"port" required:"true"`
	TargetPort int `envconfig:"target_port" required:"true"`
}

func (s Server) BindingAddress() string {
	return fmt.Sprintf(":%d", s.Port)
}

package config

import "fmt"

type Server struct {
	Host       string `envconfig:"host" default:""`
	Port       int    `envconfig:"port" required:"true"`
	TargetPort int    `envconfig:"target_port" required:"true"`
}

func (s Server) BindingAddress() string {
	return fmt.Sprintf(":%d", s.Port)
}

func (s Server) ServiceAddress() string {
	return fmt.Sprintf("%s:%d", s.Host, s.TargetPort)
}

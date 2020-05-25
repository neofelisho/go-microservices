package config

import "fmt"

type Database struct {
	Host     string `envconfig:"host" required:"true"`
	Port     int    `envconfig:"port" required:"true"`
	User     string `envconfig:"user" required:"true"`
	Password string `envconfig:"password" required:"true"`
}

func (db Database) URI() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d", db.User, db.Password, db.Host, db.Port)
}

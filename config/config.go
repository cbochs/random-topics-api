package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbName string
	dbHost string
	dbPort string
	dbUser string
	dbPswd string
}

func Get() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.dbName, "dbname", os.Getenv("POSTGRES_DB"), "database name")
	flag.StringVar(&cfg.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "database host")
	flag.StringVar(&cfg.dbPort, "dbport", os.Getenv("POSTGRES_PORT"), "database port")
	flag.StringVar(&cfg.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "database user")
	flag.StringVar(&cfg.dbPswd, "dbpswd", os.Getenv("POSTGRES_PASSWORD"), "database password")

	flag.Parse()

	return cfg
}

func (c *Config) GetDSNStr() string {
	return c.getDSNStr(c.dbHost, c.dbName)
}

func (c *Config) getDSNStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"dbname=%s host=%s port=%s user=%s password=%s sslmode=disable",
		dbname,
		dbhost,
		c.dbPort,
		c.dbUser,
		c.dbPswd,
	)
}

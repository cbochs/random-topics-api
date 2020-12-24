package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbName     string
	dbHost     string
	dbPort     string
	dbUser     string
	dbPswd     string
	dbTestName string
	dbTestHost string
	apiPort    string
	migrate    string
}

func Get() *Config {
	cfg := &Config{}

	flag.StringVar(&cfg.dbName, "dbname", os.Getenv("POSTGRES_DB"), "database name")
	flag.StringVar(&cfg.dbHost, "dbhost", os.Getenv("POSTGRES_HOST"), "database host")
	flag.StringVar(&cfg.dbPort, "dbport", os.Getenv("POSTGRES_PORT"), "database port")
	flag.StringVar(&cfg.dbUser, "dbuser", os.Getenv("POSTGRES_USER"), "database user")
	flag.StringVar(&cfg.dbPswd, "dbpswd", os.Getenv("POSTGRES_PASSWORD"), "database password")
	flag.StringVar(&cfg.dbTestName, "dbtestname", os.Getenv("TEST_DB_NAME"), "database name")
	flag.StringVar(&cfg.dbTestHost, "dbtesthost", os.Getenv("TEST_DB_HOST"), "database host")
	flag.StringVar(&cfg.apiPort, "apiport", os.Getenv("API_PORT"), "API Port")
	flag.StringVar(&cfg.migrate, "migrate", "up", "specify database migration direction {up, down}")

	flag.Parse()

	return cfg
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) GetTestDBConnStr() string {
	return c.getDBConnStr(c.dbTestHost, c.dbTestName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPswd,
		dbhost,
		c.dbPort,
		dbname,
	)
}

func (c *Config) GetDSNStr() string {
	return c.getDSNStr(c.dbHost, c.dbName)
}

func (c *Config) GetTestDSNStr() string {
	return c.getDSNStr(c.dbTestHost, c.dbTestName)
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

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}

func (c *Config) GetMigration() string {
	return c.migrate
}

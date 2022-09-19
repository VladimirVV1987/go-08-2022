/*
Пример конфигурации
port: 8080
db_url: postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable
jaeger_url: http://jaeger:16686
sentry_url: http://sentry:9000
KafkaBroker: kafka:9092
SomeAppID: testid
some_app_key: testkey
*/

package config

import (
	"fmt"
	"os"
	"regexp"
)

type SrvConfigs struct {
	AppPort     string
	DBUrl       string
	JaegerUrl   string
	SentryUrl   string
	KafkaBroker string
	SomeAppID   string
	SomeAppKey  string
}

func (c *SrvConfigs) DefaultConfigs() {
	matchAppPort, _ := regexp.MatchString(`[0-9]*`, c.AppPort)
	if c.AppPort == "" || !matchAppPort {
		c.AppPort = "8080"
	}
	matchDBUrl, _ := regexp.MatchString(`^[0-9a-z_]+:\/\/[0-9a-z_]+:[0-9]+.*`, c.DBUrl)
	if c.DBUrl == "" || !matchDBUrl {
		c.DBUrl = "postgres://db-user:db-password@petstore-db:5432/petstore?sslmode=disable"
	}
	matchJaegerUrl, _ := regexp.MatchString(`^[0-9a-z_]+:\/\/[0-9a-z_]+:[0-9]+.*`, c.JaegerUrl)
	if c.JaegerUrl == "" || !matchJaegerUrl {
		c.JaegerUrl = "http://jaeger:16686"
	}
	matchSentryUrl, _ := regexp.MatchString(`^[0-9a-z_]+:\/\/[0-9a-z_]+:[0-9]+.*`, c.SentryUrl)
	if c.SentryUrl == "" || !matchSentryUrl {
		c.SentryUrl = "http://sentry:9000"
	}

}

func (c *SrvConfigs) ReadSrvConfigs() {
	*c = SrvConfigs{
		AppPort:     os.Getenv("APP_PORT"),
		DBUrl:       os.Getenv("DB_URL"),
		JaegerUrl:   os.Getenv("JAEGER_URL"),
		SentryUrl:   os.Getenv("Sentry_Url"),
		KafkaBroker: os.Getenv("KAFKA_BROKER"),
		SomeAppID:   os.Getenv("SOME_APP_ID"),
		SomeAppkEY:  os.Getenv("SOME_APP_KEY"),
	}

}

func (c *SrvConfigs) AddSrvConfigs() {
	c.ReadSrvConfigs()
	c.DefaultConfigs()
	fmt.Println(*c)
}

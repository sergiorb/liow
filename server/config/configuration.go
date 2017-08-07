package config

import (
    "fmt"
    "os"
    "encoding/json"
)

const (
  API_VERSION = "0.1"
)

type Server struct {
  Host      string `json: "host"`
  Port      string `json: "port"`
  UrlPrefix string `json: "urlPrefix"`
}

type Api struct {
  Version       string `json: "version"`
  ApiTokenName  string `json: "apiTokenName"`
}

type Database struct {
  Host  string `json: "host"`
  Port  string `json: "port"`
  User  string `json: "user"`
  Pass  string `json: "pass"`
  Name  string `json: "name"`
}

type Logger struct {
  Level string `json: "level"`
}

type Configuration struct {
    Server    Server    `json: "server"`
    Api       Api       `json: "api"`
    Database  Database  `json: "database"`
    Logger    Logger    `json: "logger"`
}

func (c *Configuration) GetFullApiPrefix() string {

  return fmt.Sprintf("%v/v%v", c.Server.UrlPrefix, c.Api.Version)
}

func (d *Database) GetUrl() string {

  return fmt.Sprintf("mongodb://%v:%v", d.Host, d.Port)
}

func Load() *Configuration {

  file, _ := os.Open("config.json")
  decoder := json.NewDecoder(file)
  conf := &Configuration{}
  err := decoder.Decode(&conf)

  conf.Api.Version = API_VERSION

  if err != nil { fmt.Printf("%v\n", conf.Server.Host) }

  return conf
}

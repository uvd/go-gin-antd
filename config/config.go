package config

import "errors"
import "io/ioutil"
import "log"
import "encoding/json"

var (
	ErrNotFind = errors.New("Not Find")
)

type MysqlConfig struct {
	User     string
	Password string
	Port     string
	Host     string
	Database string
}
type RedisConfig struct {
}
type MongodbConfig struct {
}

type WebConfig struct {
	Port string
}

type Config struct {
	Mysql   MysqlConfig
	Redis   RedisConfig
	Mongodb MongodbConfig
	Web     WebConfig
}

func New() *Config {
	filename := "./config/config.json"
	con, err := readFile(filename)
	if err != nil {
		panic(err)
	}
	return con
}

func readFile(filename string) (*Config, error) {
	nbytes, err := ioutil.ReadFile(filename)

	config := new(Config)
	if err != nil {
		log.Fatal(err)
		return config, err
	}

	if err := json.Unmarshal(nbytes, config); err != nil {
		log.Fatal(err)
		return config, err
	}

	return config, nil

}

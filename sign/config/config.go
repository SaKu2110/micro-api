package config

import(
	"os"
	"fmt"
)

type dataBaseConfig struct {
	User string
	Pass string
	IP   string
	Port string
	Name string
}

var c dataBaseConfig

const accessTokenTemplate = "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local"

func init() {
	c = dataBaseConfig {
		User: os.Getenv("DB_USER"),
		Pass: os.Getenv("DB_PASS"),
		IP	: os.Getenv("DB_IP"),
		Port: os.Getenv("DB_PORT"),
		Name: os.Getenv("DB_NAME"),
	}

	err := checkElements(c)
	if err != nil {
		// TODO: Faild to get env Value 
	}
}

func checkElements(c dataBaseConfig) error {
	if c.User == "" {
		return fmt.Errorf("")
	}
	if c.Pass == "" {
		return fmt.Errorf("")
	}
	if c.IP == "" {
		return fmt.Errorf("")
	}
	if c.Port == "" {
		return fmt.Errorf("")
	}
	if c.Name == "" {
		return fmt.Errorf("")
	}
	return nil
}

func GetConnectionToken() string {
	return fmt.Sprintf(accessTokenTemplate, c.User, c.Pass, c.IP, c.Port, c.Name)
}
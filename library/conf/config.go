package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type DBConf struct {
	Host string `yaml:"DBHost"`
}

func GetConf() {
	var conf DBConf
	config, err := ioutil.ReadFile("../../conf/app.yaml")

	if err != nil {
		fmt.Print(err)

	}

	yaml.Unmarshal(config, &conf)

	fmt.Println(conf.Host)
}

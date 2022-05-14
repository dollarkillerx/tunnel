package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type conf struct {
	ProxyItem []proxyItem `json:"proxy_item"`
}

type proxyItem struct {
	ExternalAddress string `json:"external_address"`
	ProxyAddress    string `json:"proxy_address"`
}

var Conf *conf

func InitConf() {
	file, err := ioutil.ReadFile("tunnel.json")
	if err != nil {
		var c = conf{
			ProxyItem: []proxyItem{
				{
					ExternalAddress: "",
					ProxyAddress:    "",
				},
			},
		}
		indent, _ := json.MarshalIndent(c, "", "    ")
		err := ioutil.WriteFile("tunnel.json", indent, 00666)
		if err != nil {
			log.Fatalln("Default Generate Config Error: ", err)
		}
		log.Fatalln("Please fill in the configuration file")
	}

	var c conf
	err = json.Unmarshal(file, &c)
	if err != nil {
		log.Fatalln(err)
	}

	Conf = &c
}

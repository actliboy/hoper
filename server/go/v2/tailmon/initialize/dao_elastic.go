package initialize

import (
	"github.com/liov/hoper/go/v2/utils/log"
	"github.com/olivere/elastic"
)

type ElasticConfig struct {
	Host string
	Port int32
}

func (conf *ElasticConfig) generate() *elastic.Client {
	client, err := elastic.NewClient()
	if err != nil {
		log.Error(err)
	}
	//closes = append(closes,client.Stop)
	return client
}

func (conf *ElasticConfig) Generate() interface{} {
	return conf.generate()
}

package initialize

import (
	"github.com/olivere/elastic/v7"
	"mall.com/config/global"
)

func Elastic() {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(global.Config.ElasticSearch.Url))
	if err != nil {
		panic(err)
	}
	global.Es = client
}

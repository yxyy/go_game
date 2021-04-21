package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/olivere/elastic/v7"
	"lhc.go.game.center/libs/es"
	"lhc.go.game.center/logs"
)

type NetbianImg struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Alt  string `json:"alt" form:"alt"`
	Details string `json:"details"`
	Src  string `json:"src"`
	DownloadsUrl string `json:"downloads_url"`
	Sort int `json:"sort"`
	Mark string `json:"mark"`
	Type string `json:"type"`
	Origin string `json:"origin"`
	TickTime string `json:"tick_time"`
	Lists []string `json:"lists"`
}

func NewNetbianImg() *NetbianImg {
	return &NetbianImg{}
}

func (this *NetbianImg) GetList(param Params)(data []*NetbianImg,total int64,err error)  {
	client := es.Client.Search("gallery")
	if this.Alt!="" {
		query := elastic.NewQueryStringQuery("alt:"+this.Alt)
		client = client.Query(query)
	}
	result, err := client.From(param.Satrt).Size(param.Length).Do(context.Background())

	if err!=nil {
		return nil,0,err
	}
	total = result.Hits.TotalHits.Value
	if total > 0 {
		for _,v := range result.Hits.Hits {
			var tmp NetbianImg
			err := json.Unmarshal(v.Source, &tmp)
			if err != nil {
				logs.Error.Error(err)
				return nil,0,err
			}
			tmp.Id = v.Id
			data = append(data, &tmp)
		}
	}
	return
}

func (this *NetbianImg) GetOneById()(err error)  {
	if this.Id=="" {
		return errors.New("id 不能为空")
	}
	result, err := es.Client.Index().Index("gallery").Id(this.Id).Do(context.Background())


	fmt.Printf("%#v\n",result.Shards.Failures)

	return
}
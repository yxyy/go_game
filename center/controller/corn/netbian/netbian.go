package netbian

import (
	"context"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/olivere/elastic/v7"
	"github.com/spf13/viper"
	"lhc.go.game.center/controller/common"
	"lhc.go.game.center/libs/es"
	"lhc.go.game.center/logs"
	"lhc.go.game.center/model"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var count int

type Netbian struct {
	Mark string
	Url string
	Domain string
	Origin string
	Type string
	Category []*model.Category
}

func NewNetbian() *Netbian {
	return &Netbian{
		Origin:viper.GetString("web.netbian.name"),
		Domain:viper.GetString("web.netbian.domain"),
	}
}

func (this *Netbian) Run()  {
	fmt.Println("爬虫已启动.....")
	category := model.NewCategory()
	category.Mark=viper.GetString("web.netbian.mark")
	if err := category.GetOneCategoryByMark();err!=nil{
		logs.Corn.WithField("web","netbian").Info(err)
		return
	}
	categorys, err := category.GetChilendCategoryById()
	if err!=nil{
		logs.Corn.WithField("web","netbian").Info(err)
		return
	}
	if len(categorys) <=0 {
		logs.Corn.WithField("web","netbian").Info("无可爬数据")
		return
	}
	fmt.Println(categorys)

	for _,v := range categorys {
		var i = 1	//页数
		this.Mark=v.Title
		this.Type = v.Mark
		for  {
			fmt.Println("正在爬"+this.Mark+"取第"+strconv.Itoa(i)+"页......")
			if i==1 {
				this.Url = v.Url
			}else {
				this.Url =v.Url + "/index_"+strconv.Itoa(i)+ ".html"
			}
			if err := this.PageCralwer();err!=nil{
				fmt.Println(err)
				break
			}
			time.Sleep(1*time.Second)
			i++
		}
	}

}

func (this *Netbian) PageCralwer() error {
	if this.Url==""{
		return errors.New("url 不能为空")
	}
	resp, err := http.Get(this.Url)
	if err!=nil {
		logs.Error.Info(err)
		return err
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != 200 {
		return errors.New(this.Mark+"爬取完成")
	}

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err!=nil {
		return err
	}
	//批量操作es
	bulk := es.Client.Bulk()

	document.Find(".slist>ul>li").Each(func(i int, selection *goquery.Selection) {
		var data model.NetbianImg
		name :=selection.Find("a>b").Text()
		img_alt, _ := selection.Find("a>img").Attr("alt")
		data.Name = common.Gbktoutf8(name)
		data.Alt = common.Gbktoutf8(img_alt)
		data.Src, _ = selection.Find("a>img").Attr("src")
		data.Details, _ = selection.Find("a").Attr("href")
		count++
		data.Sort = count
		data.Mark=this.Mark
		data.Type=this.Type
		data.Origin=this.Origin
		data.TickTime=time.Now().Local().Format("2006-01-02 15:04:05")

		starIndex := strings.LastIndex(data.Details,"/")
		lastIndex := strings.LastIndex(data.Details,".")
		id := "netbian_"+data.Details[starIndex+1:lastIndex]
		parseSrc, _ := url.Parse(data.Src)
		if !parseSrc.IsAbs() {
			tmpUrl := strings.TrimLeft(data.Src, "/")
			data.Src = this.Domain+tmpUrl
		}
		parseDetails, _ := url.Parse(data.Details)
		if !parseDetails.IsAbs() {
			tmpUrl := strings.TrimLeft(data.Details, "/")
			data.Details = this.Domain+tmpUrl
		}

		//保存目录
		dir := "netbian/" + data.Type
		//下载图片
		data.DownloadsUrl = common.Downloads(data.Src, dir)

		request := elastic.NewBulkIndexRequest()
		request.Index("gallery").Id(id).Doc(data)
		bulk = bulk.Add(request)

	})

	_, err = bulk.Do(context.Background())

	time.Sleep(100*time.Microsecond)

	return err
}


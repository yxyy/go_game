package model

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"lhc.go.game.sdk/libs/mongo"
	"lhc.go.game.sdk/libs/mysql"
	"lhc.go.game.sdk/libs/redis"
	"log"
	"context"
)

type ReturnDatas struct {
	Zone map[string]interface{} `json:"zone"`
	Server []*ServerData     	`json:"server"`
	RecentlyServer interface{}  `json:"recently_server"`
}

type ServerData struct {
	Default      int    `bson:"default" json:"default"`
	IsShow       int    `bson:"is_show" json:"is_show"`
	GroupId      int    `bson:"group_id" json:"group_id"`
	IsCreate     int    `bson:"is_create" json:"is_create,omitempty"`
	OpenTime     string `bson:"open_time" json:"open_time"`
	CombineTime  string `bson:"combine_time" json:"combine_time,omitempty"`
	ServerId     string `bson:"server_id" son:"server_id,omitempty"`
	Name         string `bson:"name" json:"name"`
	Flag         string `bson:"flag" json:"flag"`
	Sort         int    `bson:"sort" json:"sort"`
	Status       int    `bson:"status" json:"status"`
	Zone         int    `bson:"zone,,int" json:"zone,int"`
	Md5Key       string `bson:"md_5_key" json:"md_5_key,omitempty"`
	Domain       string `bson:"domain" json:"domain,omitempty"`
	ApiUrl       string `bson:"api_url" json:"api_url,omitempty"`
	MaxOnline    int    `bson:"max_online" json:"max_online"`
	ServerAddr   string `bson:"server_addr" json:"server_addr"`
	ServerPort   int    `bson:"server_port" son:"server_port"`
	ChatAddr     string `bson:"chat_addr" json:"chat_addr"`
	ChatPort     int    `bson:"chat_port" son:"chat_port"`
	ChannelFlag  string `bson:"channel_flag" json:"channel_flag,omitempty"`
	PackageFlags string `bson:"package_flags" json:"package_flags,omitempty"`
	CreatedTime  string `bson:"created_time" json:"created_time,omitempty"`
	Text         string `json:"test"`
	Sign         string `json:"sign"`
	Tick         int64  `json:"tick"`
}


func (this *Server) getList() (data []*ServerData, err error) {

	if this.Param.Channel=="" {
		return nil,errors.New("channel is invalid")
	}
	val := redis.Redis.Get(this.Param.Channel).Val()
	if val=="" {
		return this.getServerByMysql()
	}else if val =="nil"{
		return nil,nil
	}

	if err = json.Unmarshal([]byte(val), &data);err!=nil{
		return nil,err
	}

	return data, nil

}

func (this *Server) getServerByMysql() (data []*ServerData, err error)  {
	this.mu.Lock()
	defer this.mu.Unlock()
	server:=redis.Redis.Get(this.Param.Channel).Val()
	if server!= "" {
		if err=json.Unmarshal([]byte(server),&data);err!=nil{
			return nil,err
		}
		return data,nil
	}

	err = mysql.MysqlConnet.Select("sb.default, sb.is_show, sb.group_id, sb.name, sb.flag," +
		"sb.is_create,sb.open_time, sb.combine_time,sb.server_id,sb.sort, sb.status, sb.zone," +
		"sc.md5_key,  sc.domain,sc.api_url, sc.max_online,sc.server_addr, sc.server_port, sc.chat_addr," +
		" sc.chat_port,scp.channel_flag, scp.package_flags").Table("xxwd_server as sb").Joins("left join xxwd_server_config as sc on sb.cflag=sc.flag").Joins("inner join xxwd_server_channel_package as scp on scp.server_flag = sb.flag").Where( "scp.channel_flag = ?",this.Param.Channel).Find(&data).Error
	if err!=nil{
		return nil,err
	}
	bytes, err := json.Marshal(&data)
	if err !=nil{
		return nil,err
	}
	redis.Redis.Set(this.Param.Channel,string(bytes),-1)
	return data ,err

}

func (this *Server) getListByMongo() ([]*ServerData, error) {
	mongo := mongo.Mongo.Database("sycenter").Collection("server_list")
	//var mogoData MongoServer
	var mogoData = make(map[string]interface{})
	filter := bson.D{{"channel", this.Param.Channel}}
	err := mongo.FindOne(context.TODO(), filter).Decode(&mogoData)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	fmt.Printf("%#v\n",mogoData)
	return nil, nil

}

func (this *Server) recentlyServer()  {

	this.ReturnData.RecentlyServer = redis.Redis.ZRange("recentServer:" + this.Param.Account,0,-1).Val()
}



package main

import (
	"encoding/json"
	"log"
	"net"
	"strings"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/oschwald/geoip2-golang"
	"github.com/tidwall/gjson"
)

type GeoIP struct {
	IP   string      `json:"ip"`
	Data interface{} `json:"data"`
}

func Handle(event json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	names := AssetNames()
	log.Printf("AssetNames: %v", names)
	data, err := Asset("data/GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}
	db, err := geoip2.FromBytes(data)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	ev, err := json.Marshal(event)
	if err != nil {
		return "", err
	}
	qs := gjson.GetBytes(ev, "params.querystring.ip").String()
	ips := strings.Split(qs, ",")
	res := []*GeoIP{}
	for _, ip := range ips {
		record, err := db.City(net.ParseIP(ip))
		if err != nil {
			return "", err
		}
		res = append(res, &GeoIP{
			IP:   ip,
			Data: record,
		})
	}
	return res, err
}

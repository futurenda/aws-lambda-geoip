package main

import (
	"encoding/json"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/oschwald/geoip2-golang"
	"github.com/tidwall/gjson"
	"log"
	"net"
	"strings"
)

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
	res := map[string]interface{}{}
	for _, ip := range ips {
		record, err := db.City(net.ParseIP(ip))
		if err != nil {
			return "", err
		}
		res[ip] = record
	}
	return res, err
}

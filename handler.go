package main

import (
	"encoding/json"
	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/oschwald/geoip2-golang"
	"log"
	"net"
)

func Handle(evt interface{}, ctx *runtime.Context) (interface{}, error) {
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
	_, err = json.Marshal(evt)
	if err != nil {
		return "", err
	}
	ip := net.ParseIP("8.8.8.8")
	record, err := db.City(ip)
	if err != nil {
		return "", err
	}
	res := map[string]interface{}{}
	res[ip.String()] = record
	return res, err
}

func main() {
	res, _ := Handle(nil, nil)
	log.Printf("Res: %s", res)
}

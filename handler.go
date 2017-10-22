package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"strings"

	"github.com/eawsy/aws-lambda-go-core/service/lambda/runtime"
	"github.com/futurenda/aws-utils-go/lambda/proxy-integration"
	"github.com/oschwald/geoip2-golang"
)

type GeoIP struct {
	IP   string      `json:"ip"`
	Data interface{} `json:"data"`
}

func Query(ips []string) ([]*GeoIP, error) {
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

	res := []*GeoIP{}
	for _, input := range ips {
		ip := net.ParseIP(input)
		if ip == nil {
			continue
		}
		record, err := db.City(ip)
		if err != nil {
			return nil, err
		}
		res = append(res, &GeoIP{
			IP:   input,
			Data: record,
		})
	}

	return res, nil
}

func HandleHTTPRequest(req *http.Request, w http.ResponseWriter) error {
	qs := req.URL.Query().Get("ip")
	ips := strings.Split(qs, ",")
	res, err := Query(ips)
	if err != nil {
		return err
	}
	data, err := json.Marshal(res)
	if err != nil {
		return err
	}
	w.Header().Add("GeoIP-Service-Version", "v2.0.1")
	_, err = w.Write(data)
	return err
}

func Handle(event json.RawMessage, ctx *runtime.Context) (interface{}, error) {
	req, err := proxyIntegration.NewRequest(event)
	if err != nil {
		return nil, err
	}
	w := proxyIntegration.NewResponseWriter()
	err = HandleHTTPRequest(req, w)
	if err != nil {
		return nil, err
	}
	return w.Response(), nil
}

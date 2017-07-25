# aws-lambda-geoip

AWS Lambda GeoIP Service based on:

- https://github.com/eawsy/aws-lambda-go-shim
- https://github.com/oschwald/geoip2-golang

## Data

http://dev.maxmind.com/geoip/geoip2/geolite2/

```
tree data
data
└── GeoLite2-City.mmdb
```

## Build

```bash
go get -u github.com/jteeuwen/go-bindata/...
go-bindata -nocompress -o data.go data/
```

```bash
docker pull eawsy/aws-lambda-go-shim:latest
go get -u -d github.com/eawsy/aws-lambda-go-core/...
make
```

## Deploy

- Runtime: python2.7
- Handler: handler.Handle

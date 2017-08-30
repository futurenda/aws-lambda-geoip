# aws-lambda-geoip

AWS Lambda GeoIP Service based on:

- https://github.com/eawsy/aws-lambda-go-shim
- https://github.com/oschwald/geoip2-golang

## Usage

```bash
curl "https://YOUR_API_GATEWAY.execute-api.us-west-2.amazonaws.com/prod?ip=1.2.3.4"
```

```bash
curl "https://YOUR_API_GATEWAY.execute-api.us-west-2.amazonaws.com/prod?ip=1.2.3.4,1.2.3.5"
```

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
dep ensure -v
make
```

## Deploy

### AWS Lambda

- Runtime: python2.7
- Handler: handler.Handle

#### Test Event

```json
{
   "params":{
      "querystring":{
         "ip":"1.2.3.4,2.3.4.5"
      }
   }
}
```

### API Gateway

#### Method Request

##### URL Query String Parameters

- `ip`

#### Body Mapping Template

- `Content-Type`: `application/json`
- `Template`: `Method Request passthrough`

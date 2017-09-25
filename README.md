# serverless-geoip

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
dep ensure -v
make
```

## Deploy

### AWS Lambda

- Runtime: python2.7
- Handler: handler.Handle

### API Gateway Quick Start Guide

1. Create Method - Any
2. Use Lambda Proxy integration - True
3. Choose Lambda Function

### Example Response

```json
[
    {
        "data": {
            "City": {
                "GeoNameID": 5804306,
                "Names": {
                    "en": "Mukilteo",
                    "ja": "\u30e0\u30ad\u30eb\u30c6\u30aa",
                    "zh-CN": "\u9a6c\u79d1\u5c14\u8482\u5965"
                }
            },
            "Continent": {
                "Code": "NA",
                "GeoNameID": 6255149,
                "Names": {
                    "de": "Nordamerika",
                    "en": "North America",
                    "es": "Norteam\u00e9rica",
                    "fr": "Am\u00e9rique du Nord",
                    "ja": "\u5317\u30a2\u30e1\u30ea\u30ab",
                    "pt-BR": "Am\u00e9rica do Norte",
                    "ru": "\u0421\u0435\u0432\u0435\u0440\u043d\u0430\u044f \u0410\u043c\u0435\u0440\u0438\u043a\u0430",
                    "zh-CN": "\u5317\u7f8e\u6d32"
                }
            },
            "Country": {
                "GeoNameID": 6252001,
                "IsoCode": "US",
                "Names": {
                    "de": "USA",
                    "en": "United States",
                    "es": "Estados Unidos",
                    "fr": "\u00c9tats-Unis",
                    "ja": "\u30a2\u30e1\u30ea\u30ab\u5408\u8846\u56fd",
                    "pt-BR": "Estados Unidos",
                    "ru": "\u0421\u0428\u0410",
                    "zh-CN": "\u7f8e\u56fd"
                }
            },
            "Location": {
                "AccuracyRadius": 1000,
                "Latitude": 47.913,
                "Longitude": -122.3042,
                "MetroCode": 819,
                "TimeZone": "America/Los_Angeles"
            },
            "Postal": {
                "Code": "98275"
            },
            "RegisteredCountry": {
                "GeoNameID": 2077456,
                "IsoCode": "AU",
                "Names": {
                    "de": "Australien",
                    "en": "Australia",
                    "es": "Australia",
                    "fr": "Australie",
                    "ja": "\u30aa\u30fc\u30b9\u30c8\u30e9\u30ea\u30a2",
                    "pt-BR": "Austr\u00e1lia",
                    "ru": "\u0410\u0432\u0441\u0442\u0440\u0430\u043b\u0438\u044f",
                    "zh-CN": "\u6fb3\u5927\u5229\u4e9a"
                }
            },
            "RepresentedCountry": {
                "GeoNameID": 0,
                "IsoCode": "",
                "Names": null,
                "Type": ""
            },
            "Subdivisions": [
                {
                    "GeoNameID": 5815135,
                    "IsoCode": "WA",
                    "Names": {
                        "en": "Washington",
                        "es": "Washington",
                        "fr": "Washington",
                        "ja": "\u30ef\u30b7\u30f3\u30c8\u30f3\u5dde",
                        "ru": "\u0412\u0430\u0448\u0438\u043d\u0433\u0442\u043e\u043d",
                        "zh-CN": "\u534e\u76db\u987f\u5dde"
                    }
                }
            ],
            "Traits": {
                "IsAnonymousProxy": false,
                "IsSatelliteProvider": false
            }
        },
        "ip": "1.2.3.4"
    },
    {
        "data": {
            "City": {
                "GeoNameID": 5804306,
                "Names": {
                    "en": "Mukilteo",
                    "ja": "\u30e0\u30ad\u30eb\u30c6\u30aa",
                    "zh-CN": "\u9a6c\u79d1\u5c14\u8482\u5965"
                }
            },
            "Continent": {
                "Code": "NA",
                "GeoNameID": 6255149,
                "Names": {
                    "de": "Nordamerika",
                    "en": "North America",
                    "es": "Norteam\u00e9rica",
                    "fr": "Am\u00e9rique du Nord",
                    "ja": "\u5317\u30a2\u30e1\u30ea\u30ab",
                    "pt-BR": "Am\u00e9rica do Norte",
                    "ru": "\u0421\u0435\u0432\u0435\u0440\u043d\u0430\u044f \u0410\u043c\u0435\u0440\u0438\u043a\u0430",
                    "zh-CN": "\u5317\u7f8e\u6d32"
                }
            },
            "Country": {
                "GeoNameID": 6252001,
                "IsoCode": "US",
                "Names": {
                    "de": "USA",
                    "en": "United States",
                    "es": "Estados Unidos",
                    "fr": "\u00c9tats-Unis",
                    "ja": "\u30a2\u30e1\u30ea\u30ab\u5408\u8846\u56fd",
                    "pt-BR": "Estados Unidos",
                    "ru": "\u0421\u0428\u0410",
                    "zh-CN": "\u7f8e\u56fd"
                }
            },
            "Location": {
                "AccuracyRadius": 1000,
                "Latitude": 47.913,
                "Longitude": -122.3042,
                "MetroCode": 819,
                "TimeZone": "America/Los_Angeles"
            },
            "Postal": {
                "Code": "98275"
            },
            "RegisteredCountry": {
                "GeoNameID": 2077456,
                "IsoCode": "AU",
                "Names": {
                    "de": "Australien",
                    "en": "Australia",
                    "es": "Australia",
                    "fr": "Australie",
                    "ja": "\u30aa\u30fc\u30b9\u30c8\u30e9\u30ea\u30a2",
                    "pt-BR": "Austr\u00e1lia",
                    "ru": "\u0410\u0432\u0441\u0442\u0440\u0430\u043b\u0438\u044f",
                    "zh-CN": "\u6fb3\u5927\u5229\u4e9a"
                }
            },
            "RepresentedCountry": {
                "GeoNameID": 0,
                "IsoCode": "",
                "Names": null,
                "Type": ""
            },
            "Subdivisions": [
                {
                    "GeoNameID": 5815135,
                    "IsoCode": "WA",
                    "Names": {
                        "en": "Washington",
                        "es": "Washington",
                        "fr": "Washington",
                        "ja": "\u30ef\u30b7\u30f3\u30c8\u30f3\u5dde",
                        "ru": "\u0412\u0430\u0448\u0438\u043d\u0433\u0442\u043e\u043d",
                        "zh-CN": "\u534e\u76db\u987f\u5dde"
                    }
                }
            ],
            "Traits": {
                "IsAnonymousProxy": false,
                "IsSatelliteProvider": false
            }
        },
        "ip": "1.2.3.5"
    }
]
```

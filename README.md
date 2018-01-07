# go-geoip

A simple CLI tool that fetches geographical information about an IP or a hostname.
The data is fetched by querying [Freegeoip.net](https://freegeoip.net/).

## Usage

The IP/hostname has to be provided using the `-target` option:

```
> go-geoip -target github.com

IP: 192.30.253.112
City: San Francisco (CA)
Country: United States (US)
```

Or

```
> go-geoip -target 192.30.253.112

IP: 192.30.253.112
City: San Francisco (CA)
Country: United States (US)
```

# Onion Scraper
A tor onion scraper/generator service written in go

[![License MIT](https://img.shields.io/badge/license-MIT-lightgrey.svg?style=flat)](LICENSE)

# How it works
* Generates random v2 or v3 onion link
* Proxy tor call
* Scrapes title and description of onion website
* Saves online onion link, title and description into CSV

# Usage

### Building
```sh
go build
```

### Running Tests
```sh
go test -v
```

### Configuration
!!! Lookup and copy enviroment file `.env.example` to `.env`

| Key          | Value       | Default | Description | 
| -------------|-------------|---------|-------------|
| DEBUG | bool | true | Use debug logger |
| TOR_PROXY | string | socks5://127.0.0.1:9050 | Tor Proxy |
| HTTP_TIMEOUT | int | 5 | Seconds to wait for HTTP request |
| ONION_VERSION | int | 2 | Onion links to generate - 2 or 3 |
| CSV_FILE | string | temp.csv | CSV file to write - required |

### Running
```sh
./onion-scraper
```

# License
MIT
 
# Elastic [![License](https://img.shields.io/:license-apache-blue.svg)](https://opensource.org/licenses/Apache-2.0) [![Go Report Card](https://goreportcard.com/badge/Rakanixu/elastic)](https://goreportcard.com/report/github.com/Rakanixu/elastic)

Elastic service with fqdn go.micro.srv.elastic

Elastic API with fqdn go.micro.api.elastic

Perform agnostic CRUD, search and queryDSL operation within elasticsearch.

## Getting Started

### Prerequisites
Get Micro
[Micro](https://github.com/micro)
```
go get github.com/micro
```

Install Consul
[https://www.consul.io/intro/getting-started/install.html](https://www.consul.io/intro/getting-started/install.html)

Run Consul
```
$ consul agent -dev -advertise=127.0.0.1
```

### Run Service manually

```
$ go run srv/main.go
```

### Run API manually

```
$ go run api/main.go
```


### Run docker containers
Compile Go binaries and build docker image. 
```
make 
```

Run docker container:
```
docker-compose -f docker-compose-build.yml up
```


## Usage
[API](https://github.com/Rakanixu/elastic/tree/master/api)

[Microservice](https://github.com/Rakanixu/elastic/tree/master/srv)

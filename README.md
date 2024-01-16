# Go Boilerplate

An API boilerplate written in Golang with Gin Framework and Gorm

### Motivation

Write restful API with fast development and developer friendly

## Table of Contents

- [Motivation](#motivation)
- [Configuration Manage](#configuration-manage)
  - [ENV Manage](#env-manage)
  - [Server Configuration](#server-configuration)
- [Installation](#installation)
  - [Local Setup Instruction](#local-setup-instruction)
  - [Develop Application in Docker with Live Reload](#develop-application-in-docker-with-live-reload)
- [Middlewares](#middlewares)
- [Boilerplate Structure](#boilerplate-structure)
- [Deployment](#deployment)
  - [Container Development Build](#container-development-build)
  - [Container Production Build and Up](#container-production-build-and-up)
- [Useful Commands](#useful-commands)
- [ENV YAML Configure](#env-yaml-configure)
- [Use Packages](#use-packages)

### Configuration Manage

#### ENV Manage

- Default ENV Configuration Manage from `.env`. sample file `.env.example`

```text
# Server Configuration
SECRET=h9wt*pasj6796j##w(w8=xaje8tpi6h*r&hzgrz065u&ed+k2)
DEBUG=True # `False` in Production
ALLOWED_HOSTS=0.0.0.0
SERVER_HOST=0.0.0.0
SERVER_PORT=8000

PODCAST_SERVICE_URL=https://601f1754b5a0e9001706a292.mockapi.io
```

- Server `DEBUG` set `False` in Production
- If ENV Manage from YAML file add a config.yml file and configuration [db.go](config/db.go) and [server.go](config/server.go). See More [ENV YAML Configure](#env-yaml-configure)

#### Server Configuration

- Use [Gin](https://github.com/gin-gonic/gin) Web Framework

### Installation

#### Local Setup Instruction

Follow these steps:

- Copy [.env.example](.env.example) as `.env` and configure necessary values
- To add all dependencies for a package in your module `go get .` in the current directory
- Locally run `go run main.go` or `go build main.go` and run `./main`
- Check Application health available on [0.0.0.0:8000/health](http://0.0.0.0:8000/health)

#### Develop Application in Docker with Live Reload

Follow these steps:

- Make sure install the latest version of docker and docker-compose
- Docker Installation for your desire OS https://docs.docker.com/engine/install/ubuntu/
- Docker Composer Installation https://docs.docker.com/compose/install/
- Run and Develop `make dev`
- Check Application heartbeat available on [0.0.0.0:8000/health](http://0.0.0.0:8000/heartbeat)

### Middlewares

- Use Gin CORSMiddleware

```go
router := gin.New()
router.Use(gin.Logger())
router.Use(gin.Recovery())
router.Use(middlewares.CORSMiddleware())
```

- Use Gin RateLimitMiddleware

```go
router := gin.New()
limiter := rate.NewLimiter(1, 5)
router.Use(middlewares.RateLimitMiddleware(limiter))
```

### Boilerplate Structure

<pre>├── <font color="#3465A4"><b>internal</b></font>
│   ├── <font color="#3465A4"><b>adapters</b></font>
│   ├── <font color="#3465A4"><b>app</b></font>
│   │   ├── <font color="#3465A4"><b>controllers</b></font>
│   │   │   └── podcast_controller.go
│   │   ├── <font color="#3465A4"><b>graphql</b></font>
│   │   │   ├── handler.go
│   │   │   └── schema.go
│   │   ├── <font color="#3465A4"><b>middlewares</b></font>
│   │   │   ├── cors.go
│   │   │   └── rate_limit.go
│   │   ├── <font color="#3465A4"><b>routers</b></font>
│   │   │   ├── index.go
│   │   │   └── router.go
│   └── <font color="#3465A4"><b>domain</b></font>
│   │   ├── <font color="#3465A4"><b>models</b></font>
│   │   ├── <font color="#3465A4"><b>repositories</b></font>
│   │   ├── <font color="#3465A4"><b>services</b></font>
│   │   │   ├── http_service.go
│   │   │   └── podcast_service.go
├── <font color="#3465A4"><b>pkg</b></font>
│   ├── <font color="#3465A4"><b>config</b></font>
│   │   ├── config.go
│   │   └── server.go
│   ├── <font color="#3465A4"><b>constants</b></font>
│   │   └── constants.go
│   ├── <font color="#3465A4"><b>logger</b></font>
│   │   └── logger.go
│   ├── <font color="#3465A4"><b>types</b></font>
│   │   └── GetPodcasts.go
│   ├── <font color="#3465A4"><b>utils</b></font>
├── docker-compose-dev.yml
├── docker-compose-prod.yml
├── Dockerfile
├── Dockerfile-dev
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── Makefile
</pre>

### Deployment

#### Container Development Build

- Run `make build`

#### Container Production Build and Up

- Run `make production`

#### ENV Yaml Configure

```yaml
server:
  host: "0.0.0.0"
  port: "8000"
  secret: "secret"
  allow_hosts: "localhost"
  debug: false #use `false` in production
  request:
    timeout: 100
```

- [Server Config](config/server.go)

```go
func ServerConfig() string {
viper.SetDefault("server.host", "0.0.0.0")
viper.SetDefault("server.port", "8000")
appServer := fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port"))
return appServer
}
```

### Useful Commands

- `make dev`: make dev for development work
- `make build`: make build container
- `make production`: docker production build and up
- `clean`: clean for all clear docker images

### Use Packages

- [Viper](https://github.com/spf13/viper) - Go configuration with fangs.
- [Gorm](https://github.com/go-gorm/gorm) - The fantastic ORM library for Golang
- [Logger](https://github.com/sirupsen/logrus) - Structured, pluggable logging for Go.
- [Air](https://github.com/cosmtrek/air) - Live reload for Go apps (Docker Development)

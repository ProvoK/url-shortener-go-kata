# URL Shortener Go Kata

URL Shortener example service written in Go.
The application should be able to create shortlink that redirects to the original url.

##  Running

The web application as the following flags:
```
Usage of ./web:
  -host string
    	public host where the application is located (default "localhost")
  -https
    	is application under https
  -port string
    	application's port (default "8080")
```

### Local

The project uses [go mod](https://blog.golang.org/using-go-modules).

Run with
```
go build -o web ./cmd/web
./web
```
The service starts at port 8080.


### Docker

TODO

## Usage

With the service running, you can create new shortlinks by doing a PUT against `/urls` route.

Eg:
```
~ $ http put :8080/urls url=https://www.google.com
HTTP/1.1 200 OK
Content-Length: 36
Content-Type: text/plain; charset=utf-8
Date: Fri, 01 May 2020 13:59:33 GMT

{
    "id": "some-id",
	"shortlink": "http://localhost:8080/r/some-id"

}
```

The shortlink url will redirect (301 Moved Permanently) to the original url (`https://www.google.com`).


## TODOs

Must
* Unit tests using [testify](https://github.com/stretchr/testify)
* Log with [zerolog](https://github.com/rs/zerolog)
* Dockerize application

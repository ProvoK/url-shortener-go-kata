# URL Shortner Go Kata

URL Shortner example service written in Go.
The application should be able to create shortlink that redirects to the original url.

##  Installation

### Local

The project uses [go mod]().

Run with
```
go run ./cmd/web/main.go
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

module bitmex-go

go 1.14

require (
	github.com/go-resty/resty v1.12.0 // indirect
	github.com/satori/go.uuid v1.2.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	golang.org/x/oauth2 v0.0.0-20201109201403-9fd604954f58
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.12.0

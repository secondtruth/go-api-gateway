# Go API Gateway

This project is a simple, but flexible and extensible API gateway written in Go. It features a variety of features that make it easy to manage API routes.

## Features

- **Custom Handlers:** Use middlewares and other handlers of your choice to handle incoming requests.
- **Response Formatting:** Format your responses in the way that best suits your API's needs. Currently, JSON and plain text are supported.
- **Authentication:** Authenticate your users with basic or bearer token authentication.
- **Logging:** Keep track of your API's usage with access logs.

This project is still under development, and more features are planned for the future. Contributions are welcome!

## Installation

To install `go-api-gateway`, use the following command:

	go get -u github.com/secondtruth/go-api-gateway

## Usage

To use `go-api-gateway`, you first need to create a `services.Context` for storing your services (e.g. responder, loggers, etc.).

For convenience, it's recommended to create a `HandlerFactory` then. This allows you to generate the handlers you wish to use,
such as reverse proxies, preconfigured with your desired services and settings.

Now, you can create a `Gateway` and add your entrypoints to it. Finally, you can start the gateway by calling `ListenAndServe` on it.

```go
package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/secondtruth/go-api-gateway/auth"
	"github.com/secondtruth/go-api-gateway/response/format"
	"github.com/secondtruth/go-api-gateway/gateway"
)

func main() {
	formatter := format.NewJsonFormatter()

	sc := services.NewContextWithDefaults()
	sc.Responder.SetFormatter(formatter)

	f := gateway.NewHandlerFactory(sc)
	f.RequestHeader = http.Header{
	 	"X-Foo": []string{"Bar"},
	}
	f.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	rp, err := f.MakeReverseProxy("http://my-backend:8000")
	if err != nil {
		log.Fatal(err)
	}

	rp.PassPath("*", "/")
	rp.PassPaths("HEAD|GET|POST", "/api/version", "/api/posts")
	rp.RewritePath("HEAD|GET|POST", "/posts", "/api/posts")

	g := gateway.NewFromServiceContext(sc)
	g.HandleHost("*.api.example.com", rp)
	
	log.Fatal(g.ListenAndServe(":8080"))
}
```

## Glossary

**Entrypoint**: An entrypoint is a specific host or path that the `Gateway` can handle. It is associated with a specific `http.Handler` that will process the incoming HTTP requests for that entrypoint.

**Service**: A service is a reusable component that provides certain functionality to the Gateway and its parts. Services can include things like responders and loggers. They are stored in a `services.Context` and can be accessed from there as needed.

**Responder**: The responder is a service that sends HTTP responses to the client. It can be configured with a specific response formatter, such as JSON or plain text.

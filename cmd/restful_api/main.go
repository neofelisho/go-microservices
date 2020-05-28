package main

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/neofelisho/go-microservices/config"
	"github.com/neofelisho/go-microservices/proto"
	"google.golang.org/grpc"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/:name", getGreeting)
	e.POST("/", echoRequest)

	// Start server
	e.Logger.Fatal(e.Start(config.MustLoad().API.BindingAddress()))
}

func echoRequest(context echo.Context) error {
	body, err := ioutil.ReadAll(context.Request().Body)
	if err != nil {
		return err
	}
	return context.String(http.StatusOK, string(body))
}

func getGreeting(context echo.Context) error {
	name := context.Param("name")
	greeting, err := sayHello(name)
	if err != nil {
		context.Logger().Errorf("occurred problem when communicating with gRPC server: %v", err)
		return context.String(http.StatusInternalServerError, "")
	}
	return context.String(http.StatusOK, greeting)
}

func sayHello(name string) (string, error) {
	cfg := config.MustLoad().GRPC
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second)
	defer cancelFunc()
	conn, err := grpc.DialContext(ctx, cfg.ServiceAddress(), grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return "", err
	}

	client := proto.NewGreeterClient(conn)
	helloReply, err := client.SayHello(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}
	err = conn.Close()
	if err != nil {
		return "", err
	}

	return helloReply.Message, nil
}

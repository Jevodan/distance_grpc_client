package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	api "github.com/Jevodan/proto/distance"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	A            float64 = 1
	B            float64 = 5
	DEFAULT_ADDR string  = "31.28.14.238:50052"
)

var (
	addr = flag.String("addr", DEFAULT_ADDR, "the address to connect to")
	ax   = flag.Float64("AX", A, "point A , value X")
	ay   = flag.Float64("AY", A, "point A , value Y")
	bx   = flag.Float64("BX", B, "point B , value X")
	by   = flag.Float64("BY", B, "point B , value Y")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := api.NewDistanceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &api.Points{
		A: &api.Point{X: *ax, Y: *ay},
		B: &api.Point{X: *bx, Y: *by},
	}
	res, err := c.GetDistance(ctx, req)
	if err != nil {
		log.Fatalf("Error calling GetDistance: %v", err)
	}
	fmt.Println("Результат: ", res.GetResult())
}

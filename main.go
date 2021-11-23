package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "qaecli/pb/gen/app"

	"google.golang.org/grpc"
)

var (
	num = flag.Int64("number", 42, "double input numer")
)

func main() {
	addr := ":8000"
	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("dial failed", addr, err)
	}
	defer conn.Close()
	c := pb.NewAppServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Double(ctx, &pb.Number{Value: *num})
	if err != nil {
		log.Fatalln("grpc app get failed", err)
	}
	log.Printf("double %v = %v\n", *num, r.Value)
}

package qgrpc

import (
	"qaecli/config"
	pb "qaecli/pb/gen/app"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func NewAppClient() (pb.AppServiceClient, func()) {
	conn, err := grpc.Dial(config.Server, grpc.WithInsecure())
	if err != nil {
		logrus.Fatalln("dail failed", err)
	}

	c := pb.NewAppServiceClient(conn)

	return c, func() {
		conn.Close()
	}
}

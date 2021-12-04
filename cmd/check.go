package cmd

import (
	"context"
	"log"
	"qaecli/config"
	pb "qaecli/pb/gen/app"
	"time"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	checkCmd *cobra.Command
)

func init() {
	checkCmd = &cobra.Command{
		Use: "check",
		Run: checkMain,
	}
	checkCmd.PersistentFlags().Int64("num", 42, "Output NUMBER * 3 Result")
}

func checkMain(cmd *cobra.Command, args []string) {
	addr := config.Server
	num, err := cmd.PersistentFlags().GetInt64("num")

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("dial failed", addr, err)
	}
	defer conn.Close()
	c := pb.NewAppServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Triple(ctx, &pb.Number{Value: num})
	if err != nil {
		log.Fatalln("grpc app get failed", err)
	}
	log.Printf("%v * 3 = %v\n", num, r.Value)
}

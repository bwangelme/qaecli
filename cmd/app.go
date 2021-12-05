package cmd

import (
	"context"
	"errors"
	"log"
	"os"
	"qaecli/config"
	pb "qaecli/pb/gen/app"
	"regexp"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var appCmd *cobra.Command

func init() {
	appCmd = &cobra.Command{
		Use:     "app",
		Aliases: []string{"a"},
		Run:     appMain,
	}

	appCmd.AddCommand(initCreateCmd())
}

func appMain(cmd *cobra.Command, args []string) {

}

func initCreateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Aliases: []string{"c"},
		Long:    "create app",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("require a app name argument")
			}

			return nil
		},
		Run: createMain,
	}

	cmd.PersistentFlags().String("repo", "", `Github Repo HTTP URL:
For Example: https://github.com/bwangelme/qae
`)
	return cmd
}

func createMain(cmd *cobra.Command, args []string) {
	repo, err := cmd.PersistentFlags().GetString("repo")
	if err != nil {
		logrus.Errorf("Get Repo failed %v/%v", repo, err)
		os.Exit(1)
	}

	if !isValidURL(repo) {
		logrus.Warningf("Invalid repo url `%v`", repo)
		os.Exit(1)
	}

	addr := config.Server
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("dial failed", addr, err)
	}
	defer conn.Close()
	c := pb.NewAppServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.CreateReq{
		Name: args[0],
		Repo: repo,
	}
	r, err := c.Create(ctx, req)
	if err != nil {
		log.Fatalln("grpc app get failed", err)
	}

	if r.Err != "" {
		log.Printf("Create app failed %v\n", r.Err)
	}

	log.Printf("Create app %v success\n", r.App.Id)
}

func isValidURL(repo string) bool {
	re := regexp.MustCompile(`https?://(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	return re.MatchString(repo)
}

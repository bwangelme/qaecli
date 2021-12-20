package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"qaecli/helper"
	pb "qaecli/pb/gen/app"
	"qaecli/qgrpc"
	"qaecli/render"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var appCmd *cobra.Command

func init() {
	appCmd = &cobra.Command{
		Use:     "app",
		Aliases: []string{"a"},
		Run:     appMain,
	}

	appCmd.AddCommand(initCreateCmd())
	appCmd.AddCommand(initListCmd())
	appCmd.AddCommand(initDeleteCmd())
	appCmd.AddCommand(initInfoCmd())
}

func appMain(cmd *cobra.Command, args []string) {
}

func initInfoCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "info",
		Aliases: []string{"i"},
		Long:    "Get App Info",
		Run:     infoMain,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("require id argument")
			}
			return nil
		},
	}

	return cmd
}

func infoMain(cmd *cobra.Command, args []string) {
	id, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		logrus.Fatalln("Invalid id", args[0], err)
	}
	c, df := qgrpc.NewAppClient()
	defer df()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.GetReq{
		Value: id,
	}

	resp, err := c.Get(ctx, req)
	if err != nil {
		logrus.Fatalln("grpc get app failed", err)
	}

	if resp.Err != "" {
		logrus.Fatalln("get app error, server return", resp.Err)
	}

	render.AppInfo(resp, os.Stdout)
}

func initDeleteCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete",
		Aliases: []string{"d"},
		Long:    "Delete App",
		Run:     deleteMain,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("require id argument")
			}
			return nil
		},
	}

	return cmd
}

func deleteMain(cmd *cobra.Command, args []string) {
	ids, err := helper.StrSliceToInt64Slice(strings.Split(args[0], ","))
	if err != nil {
		logrus.Fatalln("Invalid Arg Ids")
	}
	logrus.Infof("Start to delete app %v\n", ids)

	c, df := qgrpc.NewAppClient()
	defer df()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.DeleteReq{
		Ids: ids,
	}

	resp, err := c.Delete(ctx, req)
	if err != nil {
		logrus.Fatalln("grpc app delete failed", err)
	}

	if resp.Err != "" {
		logrus.Fatalln("Delete app failed %v\n", resp.Err)
	}

	logrus.Infof("Delete %v apps\n", resp.Cnt)
}

func initListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "list",
		Aliases: []string{"l"},
		Long:    "List App",
		Run:     listMain,
	}

	cmd.PersistentFlags().Int64("start", 0, `Start`)
	cmd.PersistentFlags().Int64("limit", 10, `Limit`)
	return cmd
}

func listMain(cmd *cobra.Command, args []string) {
	c, df := qgrpc.NewAppClient()
	defer df()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	start, _ := cmd.PersistentFlags().GetInt64("start")
	limit, _ := cmd.PersistentFlags().GetInt64("limit")
	req := &pb.ListReq{
		Start: start,
		Limit: limit,
	}
	r, err := c.List(ctx, req)
	if err != nil {
		logrus.Fatalln("grpc app list failed", err)
	}

	render.AppList(r, os.Stdout)
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

	if !helper.IsValidURL(repo) {
		logrus.Warningf("Invalid repo url `%v`", repo)
		os.Exit(1)
	}

	c, df := qgrpc.NewAppClient()
	defer df()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	req := &pb.CreateReq{
		Name: args[0],
		Repo: repo,
	}
	r, err := c.Create(ctx, req)
	if err != nil {
		logrus.Fatalln("grpc app get failed", err)
	}

	if r.Err != "" {
		logrus.Fatalln("Create app failed %v\n", r.Err)
	}

	logrus.Infof("Create app %v success\n", r.App.Id)
}

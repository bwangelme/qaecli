package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"qaecli/git"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var commitCmd *cobra.Command

func init() {
	commitCmd = &cobra.Command{
		Use:     "commit",
		Aliases: []string{"c"},
	}

	commitCmd.AddCommand(initAddCmd())
}

func initAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "add",
		Aliases: []string{"a"},
		Long:    "Add Commit",
		Run:     CommitAddMain,
	}

	return cmd
}

func CommitAddMain(cmd *cobra.Command, args []string) {
	gitroot, err := os.Getwd()
	if err != nil {
		logrus.Fatalln("Getwd failed", err)
	}

	// 检查当前仓库是否有未提交的修改
	if !checkGitStatus(gitroot) {
		logrus.Fatalln("当前仓库有未提交到 Git 的修改，请提交后再 Commit")
	}

	// 获取 git 版本
	head, err := git.RevParse(gitroot, "HEAD")
	if err != nil {
		logrus.Fatalln("git rev-parse head failed", err)
	}

	// 获取 app.yaml
	fmt.Println(filepath.Join(gitroot, "app.yaml"), head)

	// 发送请求，创建 commit
}

//checkGitStatus 检查当前仓库中是否有未提交的修改
func checkGitStatus(gitroot string) bool {
	fileStatus, err := git.Status(gitroot)
	if err != nil {
		logrus.Fatalln("Git status failed", err)
	}

	return len(fileStatus.Modified) == 0 && len(fileStatus.StagedModified) == 0
}

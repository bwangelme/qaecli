package render

import (
	"fmt"
	"io"
	"qaecli/helper"
	pb "qaecli/pb/gen/app"

	"github.com/jedib0t/go-pretty/v6/table"
)

func AppList(resp *pb.ListResp, w io.Writer) {
	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleRounded)

	t.AppendHeader(table.Row{"Id", "Name", "Owner", "Repo", "创建时间"})
	t.AppendSeparator()
	t.AppendRow(table.Row{fmt.Sprintf("Total %v", resp.Total), fmt.Sprintf("Start %v", resp.Start), fmt.Sprintf("Limit %v", resp.Limit)})
	t.AppendSeparator()
	for _, app := range resp.Apps {
		t.AppendRow(table.Row{app.Id, app.Name, app.Owner, app.Repo, helper.FmtTimeByProtoTime(app.CreatedTime, helper.DefaultTimeFmt)})
	}

	t.Render()
}

func AppInfo(resp *pb.GetResp, w io.Writer) {
	t := table.NewWriter()
	t.SetOutputMirror(w)
	t.SetStyle(table.StyleColoredBright)

	t.AppendHeader(table.Row{"字段", "值"})

	t.AppendRow(table.Row{"Name", resp.App.Name})
	t.AppendRow(table.Row{"Github Repo", resp.App.Repo})

	t.Render()
}

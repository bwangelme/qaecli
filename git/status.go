package git

import (
	"bufio"
	"bytes"
	"io"
	"os/exec"

	"github.com/sirupsen/logrus"
)

type FileStatus struct {
	StagedModified []string
	Modified       []string
	Untracked      []string
}

func Status(gitroot string) (*FileStatus, error) {
	cmd := exec.Command("git", "-C", gitroot, "status", "-s", "--porcelain=v1")
	stdout, err := cmd.Output()

	if err != nil {
		return nil, err
	}

	return parseStatusOut(bytes.NewReader(stdout))
}

func parseStatusOut(reader io.Reader) (*FileStatus, error) {
	var res = &FileStatus{
		StagedModified: make([]string, 0),
		Modified:       make([]string, 0),
		Untracked:      make([]string, 0),
	}
	s := bufio.NewScanner(reader)

	for s.Scan() {
		// skip empty line
		if len(s.Text()) < 1 {
			continue
		}

		stats := s.Text()[:2]
		filename := s.Text()[3:]
		if stats == "M " {
			res.StagedModified = append(res.StagedModified, filename)
		} else if stats == " M" {
			res.Modified = append(res.Modified, filename)
		} else if stats == "AM" {
			res.StagedModified = append(res.StagedModified, filename)
			res.Modified = append(res.Modified, filename)
		} else if stats == "A " {
			res.StagedModified = append(res.StagedModified, filename)
		} else if stats == "??" {
			res.Untracked = append(res.Untracked, filename)
		} else {
			logrus.Fatalf("Unknown stats %v on file %v in %v", stats, filename, s.Text())
		}
	}

	return res, nil
}

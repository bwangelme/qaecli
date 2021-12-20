package git

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseStatusOut(t *testing.T) {
	out := ` M .gitignore
A  cmd/commit.go
A  git/status.go
AM git/status_test.go
 M go.mod
M  cmd/app.go
?? file`
	s, err := parseStatusOut(strings.NewReader(out))
	assert.Nil(t, err)
	assert.Equal(t, len(s.Modified), 3)
	assert.Equal(t, len(s.StagedModified), 4)
	assert.Equal(t, len(s.Untracked), 1)
}

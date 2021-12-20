package git

import "os/exec"

func RevParse(gitroot, revision string) (string, error) {
	cmd := exec.Command("git", "-C", gitroot, "rev-parse", revision)
	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}
	return string(stdout), nil
}

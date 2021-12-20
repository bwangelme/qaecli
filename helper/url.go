package helper

import "regexp"

func IsValidURL(repo string) bool {
	re := regexp.MustCompile(`https?://(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)`)
	return re.MatchString(repo)
}

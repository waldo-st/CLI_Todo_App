package todo

import (
	"fmt"
	"strings"
)

func Collect_input(prompt string) (string, error) {
	fmt.Printf("\x1b[42m\x1b[1m\x1b[30m%s\x1b[0m ", prompt)
	saisie, err := Read.ReadString('\n')
	if err != nil {
		return "", err
	}
	saisie = strings.TrimSpace(saisie)
	return saisie, nil
}

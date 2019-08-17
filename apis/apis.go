package apis

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var previousConfig string

func LoadConfig(configPath string) (string, error) {
	file, err := os.Open(configPath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	first := true
	for scanner.Scan() {
		if first {
			first = false
			continue
		}
		lines = append(lines, scanner.Text())
	}
	return strings.Join(lines, "\n"), scanner.Err()
}


func SaveConfig(configPath string, content string) error {

	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	fmt.Fprintln(w, "{ -block }")
	lines := strings.Split(content, "\n")
	sort.Strings(lines)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func GetLogs(logFile string, numLastLines int) {
	if numLastLines <= 0 {
		numLastLines = 10
	}
}

func RestartPrivoxy() error {
	return nil
}
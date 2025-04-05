package prompt

import (
	"bufio"
	"fmt"
	"ollama-benchmark/internal/i18n"
	"os"
	"strings"
)

func ReadPromptsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf(i18n.T("err_file_open"), err)
	}
	defer file.Close()

	var prompts []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			prompts = append(prompts, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf(i18n.T("err_file_read"), err)
	}

	if len(prompts) == 0 {
		return nil, fmt.Errorf(i18n.T("err_file_empty"))
	}

	return prompts, nil
}

package cleaner

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Converte string de tamanho (ex: "3Mb", "500Kb", "10b") para bytes
func ParseSize(sizeStr string) (int64, error) {
	sizeStr = strings.ToUpper(strings.TrimSpace(sizeStr))

	var multiplier int64 = 1
	var numPart string

	if strings.HasSuffix(sizeStr, "KB") {
		multiplier = 1024
		numPart = strings.TrimSuffix(sizeStr, "KB")
	} else if strings.HasSuffix(sizeStr, "MB") {
		multiplier = 1024 * 1024
		numPart = strings.TrimSuffix(sizeStr, "MB")
	} else if strings.HasSuffix(sizeStr, "GB") {
		multiplier = 1024 * 1024 * 1024
		numPart = strings.TrimSuffix(sizeStr, "GB")
	} else if strings.HasSuffix(sizeStr, "B") {
		numPart = strings.TrimSuffix(sizeStr, "B")
	} else {
		numPart = sizeStr
	}

	val, err := strconv.ParseInt(strings.TrimSpace(numPart), 10, 64)
	if err != nil {
		return 0, err
	}

	return val * multiplier, nil
}

// Verifica se o arquivo atende aos critérios de limpeza e retorna o tamanho
func CheckMatch(path string, cleanFile bool, cleanFolder bool, content string) (bool, int64, error) {
	var targetSize int64 = -1
	isCheckingEmpty := (strings.ToLower(content) == "empty")

	if !isCheckingEmpty && content != "" {
		parsed, err := ParseSize(content)
		if err != nil {
			return false, 0, fmt.Errorf("formato de tamanho inválido: %s (use ex: 3Mb, 500Kb)", content)
		}
		targetSize = parsed
	}

	info, err := os.Stat(path)
	if err != nil {
		return false, 0, err
	}

	if cleanFile && !info.IsDir() {
		if isCheckingEmpty && info.Size() == 0 {
			return true, info.Size(), nil
		} else if targetSize != -1 && info.Size() == targetSize {
			return true, info.Size(), nil
		}
	}

	return false, 0, nil
}

package cleaner

import (
	"strconv"
	"strings"
)

// Converte string de tamanho (ex: "0b", "3Mb", "500Kb", "empty") para bytes
func ParseSize(sizeStr string) (int64, error) {
	sizeStr = strings.ToUpper(strings.TrimSpace(sizeStr))

	if sizeStr == "EMPTY" {
		return 0, nil
	}

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

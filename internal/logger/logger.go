package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// LogAction registra uma ação realizada no arquivo housekeeper.log do diretório alvo
func LogAction(targetPath string, message string, disabled bool) error {
	if disabled {
		return nil
	}

	logFilePath := filepath.Join(targetPath, "housekeeper.log")
	f, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	timestamp := time.Now().Format("02/01/2006 15:04:05")
	logEntry := fmt.Sprintf("[%s] %s\n", timestamp, message)

	_, err = f.WriteString(logEntry)
	return err
}

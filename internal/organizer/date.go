package organizer

import (
	"time"
)

// GetDateFolder formata a data do arquivo para o formato de pasta (ex: "2026-10")
func GetDateFolder(t time.Time) string {
	return t.Format("2006-01")
}

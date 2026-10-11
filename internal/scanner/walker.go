package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// Lê o arquivo .hkignore na raiz do diretório alvo, se existir
func LoadIgnorePatterns(rootPath string) ([]string, error) {
	ignorePath := filepath.Join(rootPath, ".hkignore")
	var patterns []string

	file, err := os.Open(ignorePath)
	if err != nil {
		if os.IsNotExist(err) {
			return patterns, nil // Se não existir, retorna lista vazia sem erro
		}
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" && !strings.HasPrefix(line, "#") {
			patterns = append(patterns, line)
		}
	}
	return patterns, scanner.Err()
}

// Verifica se um caminho deve ser ignorado com base nos padrões
func ShouldIgnore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		// Verifica se o nome do arquivo/pasta contém o padrão ignorado
		if strings.Contains(path, pattern) {
			return true
		}
	}
	return false
}

// Função principal de varredura recursiva
func WalkDir(rootPath string, callback func(path string, info os.FileInfo, err error) error) error {
	patterns, err := LoadIgnorePatterns(rootPath)
	if err != nil {
		return err
	}

	return filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignora o próprio arquivo .hkignore e pastas/arquivos listados nele
		if info.Name() == ".hkignore" || ShouldIgnore(path, patterns) {
			if info.IsDir() {
				return filepath.SkipDir // Pula a pasta inteira se ela estiver no ignore
			}
			return nil
		}

		return callback(path, info, nil)
	})
}

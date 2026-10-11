package scanner

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// LoadIgnoreRules busca e lê arquivos .hkignore subindo na árvore de diretórios até a raiz
func LoadIgnoreRules(targetPath string) []string {
	var rules []string
	curr, err := filepath.Abs(targetPath)
	if err != nil {
		curr = targetPath
	}

	// Percorre para cima procurando .hkignore em diretórios pais
	for {
		ignoreFilePath := filepath.Join(curr, ".hkignore")
		if file, err := os.Open(ignoreFilePath); err == nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if line != "" && !strings.HasPrefix(line, "#") {
					rules = append(rules, line)
				}
			}
			file.Close()
		}

		parent := filepath.Dir(curr)
		if parent == curr {
			break
		}
		curr = parent
	}

	return rules
}

// IsIgnored verifica se o arquivo ou pasta deve ser ignorado
func IsIgnored(filePath string, targetPath string, rules []string) bool {
	if len(rules) == 0 {
		return false
	}

	absPath, err := filepath.Abs(filePath)
	if err != nil {
		absPath = filePath
	}

	baseName := filepath.Base(filePath)

	absTarget, err := filepath.Abs(targetPath)
	if err != nil {
		absTarget = targetPath
	}

	relPath, err := filepath.Rel(absTarget, absPath)
	if err != nil {
		relPath = filePath
	}

	for _, rule := range rules {
		rule = strings.TrimSpace(rule)
		if rule == "" {
			continue
		}

		// 1. Compara pelo nome exato do arquivo/pasta (ex: .gitKeep)
		if baseName == rule {
			return true
		}

		// 2. Compara pelo caminho relativo exato
		if relPath == rule || filepath.Clean(relPath) == filepath.Clean(rule) {
			return true
		}

		// 3. Compara se termina com o caminho da regra (ex: 00-inbox/.gitKeep)
		if strings.HasSuffix(absPath, rule) || strings.HasSuffix(relPath, rule) {
			return true
		}

		// 4. Suporte a padrões glob (ex: *.log, .git*)
		if matched, _ := filepath.Match(rule, baseName); matched {
			return true
		}
		if matched, _ := filepath.Match(rule, relPath); matched {
			return true
		}
	}

	return false
}

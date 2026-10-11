package scanner

import (
	"os"
	"path/filepath"
)

// WalkDir percorre o diretório respeitando as regras do .hkignore
func WalkDir(targetPath string, walkFn filepath.WalkFunc) error {
	ignoreRules := LoadIgnoreRules(targetPath)

	return filepath.Walk(targetPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Se o arquivo ou diretório estiver nas regras de ignore, pula ele
		if IsIgnored(path, targetPath, ignoreRules) {
			if info.IsDir() {
				return filepath.SkipDir // Pula a pasta inteira e seu conteúdo
			}
			return nil // Pula o arquivo individual
		}

		return walkFn(path, info, nil)
	})
}

# 🗺️ Housekeeper (`hk`) - Plano de Projeto e Roadmap

Este documento detalha o planejamento, as diretrizes de segurança, a arquitetura e o roadmap de desenvolvimento da ferramenta CLI **Housekeeper**.

---

## 🎯 Visão Geral e Arquitetura

O projeto foi construído em **Go** utilizando a biblioteca **Cobra** para gerenciamento de comandos e flags. A estrutura de pastas segue os padrões limpos de desenvolvimento em Go:

- `main.go`: Ponto de entrada leve que chama o pacote de comandos.
- `cmd/`: Contém a definição dos subcomandos (`clean`, `organize`) e suas flags.
- `internal/`: Contém a lógica de negócio isolada (scanner, cleaner, organizador, gerenciador de logs/undo).

---

## 🛣️ Roadmap de Desenvolvimento

### Fase 1: Fundação e Ambiente ✅
- [x] Configuração do ambiente Go e inicialização do módulo (`go mod init housekeeper`).
- [x] Criação do esqueleto base do CLI com Cobra (`hk clean` e `hk organize`).

### Fase 2: Varredura e Respeito ao `.hkignore` ✅
- [x] Implementação do leitor de arquivos `.hkignore`.
- [x] Varredura recursiva de diretórios ignorando padrões indesejados.

### Fase 3: Funcionalidades de Limpeza (Em Andamento)
- [x] Detecção de arquivos vazios (0 bytes) com exibição de tamanho.
- [x] Parser de tamanhos flexíveis (bytes, KB, MB, GB).
- [ ] Implementação efetiva da remoção de arquivos e pastas vazias/por tamanho.
- [ ] Detecção e remoção de arquivos duplicados via Checksum (SHA-256).

### Fase 4: Funcionalidades de Organização
- [ ] Organização por extensão de arquivos.
- [ ] Organização por alfabeto (`--name alphabet`).
- [ ] Organização por data de criação (`--name createat`).

### Fase 5: Segurança, Logs e Experiência do Usuário (UX)
- [ ] **Modo Dry-Run (`--dry-run`):** Exibição detalhada de árvores de diretórios e ações simuladas.
- [ ] **Confirmações de Segurança:** Prompt interativo solicitando confirmação do usuário antes de executar ações destrutivas.
- [ ] **Flag `-y` / `--yes`:** Ignora confirmações manuais para automação ou scripts.
- [ ] **Sistema de Logs (`--no-log`):** Registro de auditoria das ações realizadas.
- [ ] **Sistema de Undo (`--no-undo`):** Geração de arquivo de reversão e comando `hk undo` para restaurar arquivos.
- [ ] **Barras de Progresso:** Feedback visual em tempo real para operações demoradas (como checksums em arquivos grandes).

### Fase 6: Distribuição Multiplataforma
- [ ] Compilação de binários otimizados para Windows, Linux e macOS (`GOOS` e `GOARCH`).
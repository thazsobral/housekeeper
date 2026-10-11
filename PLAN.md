Com base nas tarefas concluídas até agora no roadmap e no progresso das nossas últimas conversas, estamos no final da **Fase 3** e início das definições técnicas da **Fase 5**.

Aqui está o **Roadmap Atualizado**:

---

### Fase 1: Fundação e Ambiente ✅

* [x] Configuração Go, estrutura modular e Cobra CLI.

### Fase 2: Varredura e Respeito ao `.hkignore` ✅

* [x] Leitura de regras com busca hierárquica (subindo até a raiz) e filtragem de arquivos/pastas.

### Fase 3: Funcionalidades de Limpeza 🔄 *(Fase Atual)*

* [x] Parser de tamanhos flexíveis (0b, 3Mb, etc.) com exibição de tamanho e ModTime.
* [x] Remoção efetiva de arquivos por tamanho e pastas vazias (estratégia bottom-up).
* [x] Detecção de duplicados via SHA-256 com preservação automática do arquivo mais antigo (ModTime).

### Fase 4: Funcionalidades de Organização ⏳

* [x] Organização por extensão de arquivos.
* [x] Organização por alfabeto (`--name alphabet`).
* [x] Organização por data de criação (`--date createat`).

### Fase 5: Segurança, Logs e Experiência do Usuário (UX) 🏗️ *(Especificado)*

* [ ] **Modo Dry-Run (`--dry-run`):** Simulação detalhada exibindo o que acontecerá sem alterar arquivos.
* [ ] **Flag `-y` / `--yes`:** Confirmação automática para automação e scripts.
* [ ] **Controle de Logs (`--no-log`):** Opção para desativar o registro em log do que foi executado.
* [ ] **Informações de Progresso:** Feedback em tempo real durante operações de varredura/checksum.

### Fase 6: Distribuição Multiplataforma ⏳

* [ ] Compilação de binários otimizados para Windows, Linux e macOS (`GOOS` e `GOARCH`).
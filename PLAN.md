Com base nas tarefas concluídas até agora no roadmap e no progresso das nossas últimas conversas, estamos no final da **Fase 3** e início das definições técnicas da **Fase 5**.

Aqui está o **Roadmap Atualizado**:

---

### Fase 1: Fundação e Ambiente ✅

* [x] Configuração do ambiente Go e inicialização do módulo (`go mod init housekeeper`).
* [x] Criação do esqueleto base do CLI com Cobra (`hk clean` e `hk organize`).

### Fase 2: Varredura e Respeito ao `.hkignore` ✅

* [x] Implementação do leitor de arquivos `.hkignore`.
* [x] Varredura recursiva de diretórios ignorando padrões indesejados.

### Fase 3: Funcionalidades de Limpeza 🔄 *(Fase Atual)*

* [x] Detecção de arquivos vazios (0 bytes) com exibição de tamanho.
* [x] Parser de tamanhos flexíveis (bytes, KB, MB, GB).
* [ ] Implementação efetiva da remoção de arquivos e pastas vazias/por tamanho.
* [ ] Detecção e remoção de arquivos duplicados via Checksum (SHA-256).

### Fase 4: Funcionalidades de Organização ⏳

* [ ] Organização por extensão de arquivos.
* [ ] Organização por alfabeto (`--name alphabet`).
* [ ] Organização por data de criação (`--name createat`).

### Fase 5: Segurança, Logs e Experiência do Usuário (UX) 🏗️ *(Especificado)*

* [ ] **Modo Dry-Run (`--dry-run`):** Simulação detalhada exibindo o que acontecerá sem alterar arquivos.
* [ ] **Flag `-y` / `--yes`:** Confirmação automática para automação e scripts.
* [ ] **Controle de Logs (`--no-log`):** Opção para desativar o registro em log do que foi executado.
* [ ] **Controle de Undo (`--no-undo`):** Opção para desativar a geração do arquivo de reversão.
* [ ] **Comando Undo (`hk undo`):** Leitura do arquivo de reversão para restaurar ações passadas.
* [ ] **Informações de Progresso:** Feedback em tempo real durante operações de varredura/checksum.

### Fase 6: Distribuição Multiplataforma ⏳

* [ ] Compilação de binários otimizados para Windows, Linux e macOS (`GOOS` e `GOARCH`).

---
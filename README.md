# Housekeeper (hk) 🧹

O **Housekeeper** (`hk`) é uma ferramenta CLI moderna e eficiente escrita em **Go** para organizar, limpar e gerenciar arquivos e diretórios de forma recursiva, segura e personalizável.

## 🚀 Funcionalidades Principais

- **Limpeza Inteligente (`clean`):**
  - Remoção de arquivos ou pastas vazias (`--content empty`).
  - Remoção por tamanho específico ou limite (ex: `3Mb`, `500Kb`, `10b`).
  - Detecção e remoção de arquivos duplicados via checksum (SHA-256).
- **Organização de Arquivos (`organize`):**
  - Organização por extensão de arquivo (`--extension`).
  - Organização por nome/alfabeto (`--name alphabet`).
  - Organização por data de criação (`--name createat`).
- **Segurança e Controle:**
  - **Modo Dry-Run (`--dry-run`):** Simula todas as ações e exibe a hierarquia/alterações sem modificar o disco.
  - **Confirmações Interativas:** Exige confirmação antes de executar ações destrutivas (com suporte a `-y` / `--yes` para pular).
  - **Sistema de Logs e Reversão:** Registra as ações executadas (`--no-log` para desativar log e `--no-undo` para desativar arquivo de reversão).
- **Filtros Personalizados (`.hkignore`):**
  - Suporte nativo a um arquivo `.hkignore` para ignorar pastas e arquivos específicos (como `node_modules`, `.git`, etc.) de forma similar ao `.gitignore`.

---

## 📦 Instalação e Execução

Certifique-se de ter o [Go instalado](https://go.dev/) em sua máquina.

1. Clone ou acesse a pasta do repositório.
2. Execute o projeto usando o Go:
   ```bash
   go run main.go --help
   ```
3. Para compilar o binário executável localmente:
   ```bash
   go build -o hk main.go
   ```

---

## 💡 Exemplos de Uso

### Limpeza
- Simular a remoção de arquivos vazios na pasta atual:
  ```bash
  hk clean --file --content empty --dry-run .
  ```
- Remover arquivos com tamanho exato de 3 Megabytes:
  ```bash
  hk clean --file --content 3Mb .
  ```

### Organização
- Simular a organização por extensão:
  ```bash
  hk organize --file --extension --dry-run .
  ```
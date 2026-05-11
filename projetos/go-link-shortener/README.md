# 🔗 Shortify API

Uma API de encurtamento de URLs rápida, eficiente e escalável construída com Go.  
O projeto segue uma arquitetura limpa e modular, focada em separação de responsabilidades e facilidade de manutenção.

O encurtamento das URLs utiliza geração de hash com **SHA-256 + Base64**, garantindo códigos curtos consistentes e seguros.

---

# 🚀 Tecnologias Utilizadas

- **Linguagem:** Go
- **Roteador HTTP:** Chi
- **Banco de Dados:** PostgreSQL
- **Driver PostgreSQL:** pgx/v5
- **Containerização:** Docker + Docker Compose


# 📌 Endpoints da API

---

## 🔗 Encurtar URL

Gera um código curto para a URL enviada.

### Rota

```http
POST /api/shorten
```

### Body

```json
{
  "url": "https://go.dev/doc/effective_go"
}
```

### Response — `201 Created`

```json
{
  "original_url": "https://go.dev/doc/effective_go",
  "short_code": "k9zQwA",
  "created_at": "2026-05-11T17:00:00Z"
}
```

---

## 🌍 Buscar URL Original

Retorna a URL original associada ao código curto.

### Rota

```http
GET /{code}
```

### Response — `200 OK`

```json
{
  "original_url": "https://go.dev/doc/effective_go"
}
```

---

# 🧠 Arquitetura

A aplicação segue o fluxo:

```txt
HTTP Request
    ↓
Handler
    ↓
Use Case
    ↓
Repository
    ↓
PostgreSQL
```

Essa separação facilita:

- manutenção
- testes
- escalabilidade
- reutilização de código

---

# 🔐 Estratégia de Encurtamento

O código curto é gerado utilizando:

- SHA-256
- Base64 URL Safe
- truncamento do hash

Exemplo:

```txt
https://go.dev/doc/effective_go
↓
k9zQwA
```

---

# 📈 Funcionalidades Futuras

- [ ] Contador de acessos
- [ ] Expiração de links
- [ ] Autenticação JWT
- [ ] Links customizados
- [ ] Analytics
- [ ] Rate Limiting
- [ ] Redis Cache
- [ ] QR Code

---

# 🧪 Testes

Executar todos os testes:

```bash
go test ./...
```
---

# 📄 Licença

Este projeto está sob a licença MIT.


# REST API 2

API REST simples desenvolvida em Go utilizando:

- Go
- Gin
- PostgreSQL
- Docker

## Funcionalidades

- Criar produtos
- Listar produtos
- Buscar produto por ID
- Atualizar produto
- Deletar produto

---

# Rotas

## Buscar produtos

```http
GET /products
```

---

## Buscar produto por ID

```http
GET /product/:productId
```

---

## Criar produto

```http
POST /product
```

### Body

```json
{
  "name": "Teclado",
  "price": 150
}
```

---

## Atualizar produto

```http
PUT /product/:productId
```

### Body

```json
{
  "name": "Mouse",
  "price": 120
}
```

---

## Deletar produto

```http
DELETE /product/:productId
```

---

# Rodando o projeto

## Clone o repositório

```bash
git clone <url-do-repositorio>
```

---

## Instale as dependências

```bash
go mod tidy
```

---

## Execute a aplicação

```bash
go run ./cmd/api
```

---

# Docker

## Build da imagem

```bash
docker build -t rest-api-2 .
```

---

## Rodar container

```bash
docker run -p 8000:8000 rest-api-2
```
# Projetos

Aplicações completas ou estudos maiores em Go.

## Índice

| Projeto | Descrição |
| --- | --- |
| [`REST-API-1`](REST-API-1) | Primeira API REST para praticar organização básica. |
| [`REST-API-2`](REST-API-2) | API REST de produtos com Gin, PostgreSQL e Docker. |
| [`go-link-shortener`](go-link-shortener) | Encurtador de URLs com Chi, PostgreSQL, Redis e Docker. |

## Padrão recomendado

Cada projeto deve conter:

- `README.md` com descrição, rotas e instruções de execução.
- `go.mod` próprio.
- `cmd/` para pontos de entrada.
- `internal/` para código da aplicação.
- `docker-compose.yml` quando depender de serviços externos.

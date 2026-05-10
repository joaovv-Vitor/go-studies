FROM golang:1.26

# Diretório de trabalho
WORKDIR /go/src/app

# Copia os arquivos
COPY . .

# Porta da API
EXPOSE 8000

# Build da aplicação
RUN go build -o main ./cmd/api

# Executa o binário
CMD ["./main"]
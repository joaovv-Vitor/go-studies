-- Criação da tabela de URLs
CREATE TABLE IF NOT EXISTS urls (
    id SERIAL PRIMARY KEY,
    original_url TEXT NOT NULL,
    short_code VARCHAR(20) NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

-- Índice para acelerar a busca pelo código curto na hora do redirecionamento
CREATE INDEX IF NOT EXISTS idx_short_code ON urls(short_code);
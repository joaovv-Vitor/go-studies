package hash

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

// GenerateShortCode cria um hash da string e retorna os primeiros 7 caracteres
func GenerateShortCode(input string) string {
	hasher := sha256.New()
	hasher.Write([]byte(input))
	hashBytes := hasher.Sum(nil)

	// Usa o Base64 seguro para URLs
	encoded := base64.URLEncoding.EncodeToString(hashBytes)

	// Remove caracteres que podem não ser visualmente agradáveis em uma URL curta
	encoded = strings.ReplaceAll(encoded, "-", "")
	encoded = strings.ReplaceAll(encoded, "_", "")

	// Trunca para 7 caracteres (tamanho padrão de encurtadores)
	if len(encoded) > 7 {
		return encoded[:7]
	}

	return encoded
}
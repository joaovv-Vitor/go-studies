package main

import (
	"fmt"
	"os"
)

func salvarDiario(caminho, conteudo string) error {
	// Salve o conteúdo com os.WriteFile e permissão 0644.
	// Se falhar, retorne o erro com contexto usando fmt.Errorf e %w.
	err := os.WriteFile(caminho, []byte(conteudo), 0644)
	if err != nil {
		return fmt.Errorf("Salvar diário: %w", err)
	}
	return nil
}

func carregarDiario(caminho string) (string, error) {
	// Leia o arquivo com os.ReadFile.
	// Se falhar, retorne string vazia e erro contextualizado.
	// Se funcionar, converta []byte para string.
	conteudo, err := os.ReadFile(caminho)
	if err != nil {
		return "", fmt.Errorf("erro ao carregar: %w", err)
	}

	return string(conteudo), nil
}

func main() {
	caminho := "diario.txt"

	conteudo := `Missão: Aurora
				Destino: Europa
				Status: concluída`

	if err := salvarDiario(caminho, conteudo); err != nil {
		fmt.Println("Erro:", err)
		return
	}

	diario, err := carregarDiario(caminho)
	if err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Diário carregado:")
	fmt.Println(diario)
}

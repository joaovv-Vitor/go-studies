#!/bin/bash

# Este script testa a performance da rota de GET (Postgres vs Redis)
# O servidor precisa estar rodando para este script funcionar (go run cmd/api/main.go)

PORT=8080
API_URL="http://localhost:$PORT"

echo "============================================="
echo "🧪 Teste de Cache (PostgreSQL vs Redis)"
echo "============================================="

echo -e "\n1. Criando uma URL encurtada de teste..."
# Gera uma URL aleatória para não pegar cache antigo
RANDOM_URL="https://example.com/test-$(date +%s)"

RESPONSE=$(curl -s -X POST "$API_URL/api/shorten" \
  -H "Content-Type: application/json" \
  -d "{\"url\": \"$RANDOM_URL\"}")

# Extrai o short_code usando grep e sed (funciona sem precisar instalar jq)
SHORT_CODE=$(echo "$RESPONSE" | grep -o '"short_code":"[^"]*' | grep -o '[^"]*$')

if [ -z "$SHORT_CODE" ]; then
    echo "❌ Erro ao criar a URL. O servidor está rodando?"
    echo "Resposta recebida: $RESPONSE"
    exit 1
fi

echo "✅ URL criada com sucesso!"
echo "URL Original: $RANDOM_URL"
echo "Short Code gerado: $SHORT_CODE"

echo -e "\n---------------------------------------------"
echo "2. Primeira requisição (CACHE MISS - vai no banco de dados)"
echo "---------------------------------------------"

# O parâmetro -w '%{time_total}' do curl mede o tempo total da requisição em segundos
TIME1=$(curl -s -o /dev/null -w "%{time_total}s\n" "$API_URL/$SHORT_CODE")
echo "Tempo de resposta: $TIME1"

echo -e "\n---------------------------------------------"
echo "3. Próximas requisições (CACHE HIT - vai direto no Redis)"
echo "---------------------------------------------"

for i in {1..5}
do
    TIME=$(curl -s -o /dev/null -w "%{time_total}s\n" "$API_URL/$SHORT_CODE")
    echo "Requisição $i: Tempo de resposta: $TIME"
    sleep 0.2
done

echo -e "\n============================================="
echo "✨ Conclusão:"
echo "A primeira requisição demorou $TIME1 porque precisou ir no PostgreSQL e gravar no Redis."
echo "As requisições seguintes foram quase instantâneas porque leram direto da memória (Redis)!"
echo "============================================="

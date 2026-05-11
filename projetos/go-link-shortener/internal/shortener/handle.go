package shortener

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Handler agrupa as dependências para as rotas HTTP.
type Handler struct {
	useCase UseCase
}

// NewHandler cria uma nova instância do controlador.
func NewHandler(uc UseCase) *Handler {
	return &Handler{useCase: uc}
}

// CreateShortURL recebe o JSON com a URL original e devolve a entidade URL criada.
func (h *Handler) CreateShortURL(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest

	// Decodifica o corpo da requisição para a nossa struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validação básica
	if req.URL == "" {
		http.Error(w, "O campo 'url' é obrigatório", http.StatusBadRequest)
		return
	}

	// Passa a responsabilidade para o UseCase (regra de negócio)
	url, err := h.useCase.CreateShortURL(r.Context(), req.URL)
	if err != nil {
		http.Error(w, "Erro interno ao processar a URL", http.StatusInternalServerError)
		return
	}

	// Retorna a entidade criada com status 201 (Created)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)
}

// GetOriginalURL busca a URL original pelo código e a retorna em formato JSON.
func (h *Handler) GetOriginalURL(w http.ResponseWriter, r *http.Request) {
	// Pega o código da URL (ex: k9zQw)
	code := chi.URLParam(r, "code")

	if code == "" {
		http.Error(w, "Código não informado", http.StatusBadRequest)
		return
	}

	// Busca a URL original no UseCase
	originalURL, err := h.useCase.GetOriginalURL(r.Context(), code)
	if err != nil {
		if errors.Is(err, ErrURLNotFound) {
			http.Error(w, "URL não encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, "Erro interno ao buscar a URL", http.StatusInternalServerError)
		return
	}

	// Monta a resposta em formato de mapa (que será convertido para JSON)
	response := map[string]string{
		"original_url": originalURL,
	}

	// Retorna o status 200 OK e o JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

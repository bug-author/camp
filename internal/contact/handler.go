package contact

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Handler struct {
	repo *Repository
}

func NewHandler(repo *Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var contact Contact

	if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
		fmt.Println(err)
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	// check if email already exists
	existingContact, err := h.repo.GetByEmail(contact.Email)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")

	if existingContact.Email != "" {
		w.WriteHeader(http.StatusConflict)

		resp := map[string]string{
			"message": "email already exists.",
		}

		json.NewEncoder(w).Encode(resp)
		return
	}

	createdId, err := h.repo.Create(&contact)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusCreated)

	resp := map[string]int64{
		"id": createdId,
	}
	json.NewEncoder(w).Encode(resp)
}

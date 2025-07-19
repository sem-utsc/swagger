package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"swagger-go/models"
	"swagger-go/storage"
)

func GetCharactersHandler(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))
	limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
	data, paginaActual, numTotalPaginas, totalPersonajes := storage.GetCharacters(page, limit)
	resp := map[string]interface{}{
		"PaginaActual": paginaActual,
		"NumTotalPaginas": numTotalPaginas,
		"PersonajesPorPagina": limit,
		"TotalPersonajes": totalPersonajes,
		"data": data,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func CreateCharacterHandler(w http.ResponseWriter, r *http.Request) {
	var c models.Character
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Error de parametros", "status": 400})
		return
	}
	if c.Nombre == "" || c.Calidad == "" || c.Clase == "" || c.Salud == 0 || c.Ataque == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Campos requeridos faltantes", "status": 400})
		return
	}
	created := storage.AddCharacter(c)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(created)
}

func DeleteCharacterHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/characters/"):] // asumiendo ruta /characters/{id}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "ID inválido", "status": 400})
		return
	}
	if storage.DeleteCharacter(id) {
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "ok"})
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]interface{}{"message": "Not delete character", "status": 404})
	}
}

package models

type Character struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Calidad   string `json:"calidad"`
	Clase     string `json:"clase"`
	Salud     int    `json:"salud"`
	Ataque    int    `json:"ataque"`
	FechaCreacion string `json:"fechaCreacion"`
}

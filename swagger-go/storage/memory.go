package storage

import (
	"sync"
	"time"
	"swagger-go/models"
)

// MemoryStore almacena los personajes en memoria
var (
	characters []models.Character
	mu sync.Mutex
	lastID int
)

func GetCharacters(page, limit int) ([]models.Character, int, int, int) {
	mu.Lock()
	defer mu.Unlock()
	total := len(characters)
	if limit <= 0 {
		limit = 10
	}
	if page <= 0 {
		page = 1
	}
	start := (page - 1) * limit
	end := start + limit
	if start > total {
		return []models.Character{}, page, 0, total
	}
	if end > total {
		end = total
	}
	return characters[start:end], page, (total+limit-1)/limit, total
}

func AddCharacter(c models.Character) models.Character {
	mu.Lock()
	defer mu.Unlock()
	lastID++
	c.ID = lastID
	c.FechaCreacion = time.Now().Format(time.RFC3339)
	characters = append(characters, c)
	return c
}

func DeleteCharacter(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, c := range characters {
		if c.ID == id {
			characters = append(characters[:i], characters[i+1:]...)
			return true
		}
	}
	return false
}

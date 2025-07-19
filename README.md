# Characters API - Swagger Go

Este proyecto es una API REST para consultar, crear y eliminar personajes de Clash Royale, desarrollada en Go 1.24 usando la librería estándar.

## Requisitos previos

- Go 1.24 instalado en tu sistema.
- Acceso a la terminal (Linux recomendado).

## Ejecución del programa

1. **Clona el repositorio o navega a la carpeta del proyecto:**
   ```bash
   cd /home/sem-hernandez/cursos/swagger/swagger-go
   ```

2. **Ejecuta el servidor:**
   ```bash
   go run main.go
   ```
   El servidor estará disponible en [http://localhost:3000](http://localhost:3000).

## Agregar información inicial (30 personajes)

Puedes agregar 30 personajes de ejemplo usando el siguiente script en tu terminal:

```bash
for i in {1..30}; do
  calidades=("Común" "Rara" "Épica" "Legendaria")
  clases=("Tanque" "Apoyo" "Ofensivo" "Defensivo")
  calidad=${calidades[$(( (i-1)%4 ))]}
  clase=${clases[$(( (i-1)%4 ))]}
  curl -X POST http://localhost:3000/characters \
    -H "Content-Type: application/json" \
    -d "{\"nombre\":\"Personaje $i\",\"calidad\":\"$calidad\",\"clase\":\"$clase\",\"salud\":$((500 + i * 10)),\"ataque\":$((100 + i * 5))}"
done
```

Este script crea 30 personajes con datos variados usando el endpoint `/characters`.

## Probar los endpoints

- **Listar personajes (GET):**
  ```bash
  curl http://localhost:3000/characters
  ```

- **Crear personaje (POST):**
  ```bash
  curl -X POST http://localhost:3000/characters \
    -H "Content-Type: application/json" \
    -d '{"nombre":"Ejemplo","calidad":"Épica","clase":"Tanque","salud":1000,"ataque":200}'
  ```

- **Eliminar personaje (DELETE):**
  ```bash
  curl -X DELETE http://localhost:3000/characters/1
  ```

## Notas

- El almacenamiento es en memoria, por lo que los datos se perderán al reiniciar el servidor.
- Puedes modificar los scripts para agregar más personajes o cambiar los
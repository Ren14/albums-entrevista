## Proyecto para entrevista técnica con Pedidos YA

### Autor
Renzo Mauro Ontivero
[renn.carp@gmail.com]()

### Idea
A partir del tutorial denominado Tutorial: Developing a RESTful API with Go and Gin (https://go.dev/doc/tutorial/web-service-gin), me propuse tomar como base el CRUD que se explica, y continuar aplicando conceptos básicos de CleanCode, SOLID y arquitectura de puertos y adaptadores. 

### ¿Qué hace esta API ?
Permite insertar, consultar todos los albums de música almacenados o a partir de su ID. Dichos albums tienen la siguiente estructura:
```
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}
```

### Scaffolding

La API está dividida en 3 grandes secciones. 

El paquete `rest` es la capa que contiene la rutas y los handlers. 

El paquete `domain` es la capa que contiene las reglas de negocio y el dominio de la aplicación.

El paquete `repository` es la capa que contiene la lógica para comunicarse con la infraestructura necesaria para ejecutar la aplicación (Ejemplo: persistencia, conexión con otras APIs, publishers, servicio de bloqueo distribuido, entreo otros)

Adicionalmente el paquete `configs` contiene una estructura que permite mappear el contenido del archivo `config/local.yml`, que permite configurar la aplicación según el entorno donde se ejecute.

Por último el bootstraping de la aplicación está desarrollado en el archivo `cmd/main.go`. Allí se puede observar cómo se gestiona la inyección de dependencias del proyecto.

### ¿ Cómo utilizarla ?
1. Revisar pre requisitos https://go.dev/doc/tutorial/web-service-gin
2. Instalar dependencias `go mod tidy`
3. Ejecutar servidor `go run cmd/main.go`

### Collection para interactuar
Listar todos los albums almacenados
```
curl --location 'localhost:8080/albums'
```

Agregar un album nuevo
```
curl --location 'localhost:8080/albums' \
--header 'Content-Type: application/json' \
--data '{
    "id": "3",
    "title": "Gulp!",
    "artist": "Patricio Rey y sus Redonditos de Ricota",
    "price": 17.99
}'
```

Obtener un album a partir de su ID
```
curl --location 'localhost:8080/albums/2'
```

Listar todos los albums con precio en Dólares
```
localhost:8080/albums/usd
```

### Sitios web consultados
1. https://go.dev/doc/tutorial/web-service-gin
2. https://go.dev/doc/effective_go
3. https://github.com/Streelzz/car-pooling
4. https://github.com/uber-go/mock
# API REST Backend SADU - Eventos Deportivos

 Backend en Go para gestionar eventos, torneos y atletas en sistema universitario.
 API REST escalable con Gin + GORM + SQLite3.

**Características**
```text
 - CRUD completo: Eventos, Torneos, Atletas...

 - Hot-reload: Cambios Go en tiempo real con Air

 - SQLite3: Base embebida (sin servidor externo)

 - Gin + GORM: Rutas optimizadas y ORM robusto

 - Backend: Go >1.23 + Gin Framework + GORM

   ```
   **📁 Estructura del proyecto**

```text
├─SERVICIO-SADU-BACK/
├── ./cmd/                # main.go
├── ./config/             # configuracion de BD y variables de entorno
├── ./helpers/            # Funciones reutilizables (utils)
├── ./internal/
├─────./internal/handlers # Controladores hechos con GIN
├─────./internal/middlewares 
├─────./routes            # Rutas y endpoint
├─────./services          # Logica de negocio y metodos HTTP
├── ./schemas/            # Schemas GORM
├── ./seed/               # Pruebas con datos falsos
├── ./tmp/                # Build temporal (.gitignore)
├── .air.toml             # Multiplataforma
├── .air.linux.toml       # Multiplataforma
└── Makefile              # Comandos
  ```
## Run Locally

Clone the project

```bash
git clone https://github.com/SortexGuy/servicio-sadu-back.git
```

Go to the project directory

```bash
cd servicio-sadu-back

```

```bash
go mod tidy

```

Si estas en windows puedes usar:

```bash
Air
```

o
```bash
make run
```
Si estas en linux y quieres utilizar el archivo .air.toml debes cambiar la configuracion dentro del archivo a las siguientes:

```toml
cmd = "go build -o tmp/main ./cmd"
bin = "./tmp/main"
```

```bash
air
```

o
```bash
air -c .air.linux.toml
```


En caso contrario al Air, puedes usar:

o
```bash
make run
```


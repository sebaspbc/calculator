# Sezzle Calculator

*[English version below / Versión en español más abajo]*

---

## English

Full-stack calculator application built as a technical assessment. React
(TypeScript) frontend consuming a Go REST API backend, with basic and advanced
arithmetic operations, input validation, unit tests, and Docker support.

**Live demo:** [frontend URL] · **API:** [backend URL]

### Tech stack

- **Backend:** Go, [`chi`](https://github.com/go-chi/chi) router
- **Frontend:** React 18, TypeScript, Vite
- **Testing:** Go's `testing` package (table-driven tests), Vitest + React Testing Library
- **Deployment:** Render (backend, Docker) + Netlify (frontend)

### Project structure

\`\`\`
Calculator/
├── backend/
│   ├── cmd/                  # entry point (main.go)
│   ├── internal/
│   │   ├── calculator/       # domain logic (pure functions, no HTTP)
│   │   ├── service/          # service layer (interface for DI/testability)
│   │   └── handlers/         # HTTP handlers
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── components/       # Calculator UI
│   │   ├── services/         # API client
│   │   └── types/            # shared TypeScript types
│   └── Dockerfile
├── docker-compose.yml
├── PROMPTS.md                # AI usage log
└── README.md
\`\`\`

### Setup & running locally

**Prerequisites:** Go 1.22+, Node.js 20+, Docker (optional).

**Backend:**
\`\`\`bash
cd backend
go mod tidy
go run ./cmd
\`\`\`
Server starts on `http://localhost:8080`. Health check: `GET /health`.

**Frontend:**
\`\`\`bash
cd frontend
cp .env.example .env   # sets VITE_API_URL=http://localhost:8080
npm install
npm run dev
\`\`\`
App starts on `http://localhost:5173`.

**Docker Compose (both services):**
\`\`\`bash
docker compose up --build
\`\`\`
Frontend: `http://localhost:3000` · Backend: `http://localhost:8080`

### Running tests

\`\`\`bash
# Backend, with coverage
cd backend && go test ./... -v -cover

# Frontend, with coverage
cd frontend && npm run test
\`\`\`

### API reference

**`POST /api/v1/calculate`**

| Field       | Type   | Required | Notes                                      |
|-------------|--------|----------|---------------------------------------------|
| `operation` | string | yes      | `add`, `subtract`, `multiply`, `divide`, `power`, `sqrt`, `percentage` |
| `a`         | number | yes      | First operand                               |
| `b`         | number | conditional | Required for binary operations. Not used for `sqrt`/`percentage`. |

Success — `200 OK`: `{ "result": 8 }`
Error — `400`/`422`: `{ "error": "division by zero is not allowed" }`

**Examples:**
\`\`\`bash
curl -X POST http://localhost:8080/api/v1/calculate \\
  -H "Content-Type: application/json" \\
  -d '{"operation":"add","a":5,"b":3}'
# → {"result":8}

curl -X POST http://localhost:8080/api/v1/calculate \\
  -H "Content-Type: application/json" \\
  -d '{"operation":"sqrt","a":16}'
# → {"result":4}

curl -X POST http://localhost:8080/api/v1/calculate \\
  -H "Content-Type: application/json" \\
  -d '{"operation":"divide","a":10,"b":0}'
# → {"error":"division by zero is not allowed"}
\`\`\`

### Design decisions & assumptions

- **Go was chosen for the backend** per the assignment's stated preference.
- **Layered backend architecture** (`calculator` → `service` → `handlers`): the
  service layer exists as an interface consumed by the HTTP handler, enabling
  dependency injection and isolated handler tests via a mock — without a
  `Repository`/`Model` layer, since this project has no persistence.
- **`b` as `*float64` (pointer)** in the domain layer distinguishes "operand not
  provided" (`nil`) from "operand is zero" (`&0.0`) — meaningful for division and
  for validating unary vs. binary operations.
- **Percentage** is implemented as a unary operation (`a / 100`), since the
  assignment didn't specify two-operand percentage semantics.
- **CORS** is restricted to a single configurable origin (`FRONTEND_ORIGIN` env
  var) rather than a wildcard.
- **Chaining behavior** (`5 + 3 + 2` without pressing `=`) mirrors standard
  physical calculator UX.
- **Operation history** (last 10 operations, client-side only) was added as a UX
  enhancement beyond the core requirements.

### Deployment

- **Backend:** [Render](https://render.com), Docker-based web service using
  `backend/Dockerfile`. Env vars: `FRONTEND_ORIGIN` (set to the Netlify URL), `PORT`
  (set automatically by Render).
- **Frontend:** [Netlify](https://netlify.com), built from `frontend/` with
  `VITE_API_URL` set to the Render backend URL at build time.

### AI usage

Built with assistance from Claude (Anthropic). See [`PROMPTS.md`](./PROMPTS.md) for
a summary of how AI was used throughout development.

---

## Español

Aplicación de calculadora full-stack construida como prueba técnica. Frontend en
React (TypeScript) que consume una API REST en Go, con operaciones aritméticas
básicas y avanzadas, validación de entrada, tests unitarios y soporte Docker.

**Demo en vivo:** [URL del frontend] · **API:** [URL del backend]

### Stack tecnológico

- **Backend:** Go, router [`chi`](https://github.com/go-chi/chi)
- **Frontend:** React 18, TypeScript, Vite
- **Tests:** paquete `testing` de Go (tabla-driven), Vitest + React Testing Library
- **Despliegue:** Render (backend, Docker) + Netlify (frontend)

### Estructura del proyecto

\`\`\`
Calculator/
├── backend/
│   ├── cmd/                  # punto de entrada (main.go)
│   ├── internal/
│   │   ├── calculator/       # lógica de dominio (funciones puras, sin HTTP)
│   │   ├── service/          # capa de servicio (interfaz para DI/testeo)
│   │   └── handlers/         # handlers HTTP
│   └── Dockerfile
├── frontend/
│   ├── src/
│   │   ├── components/       # UI de la calculadora
│   │   ├── services/         # cliente de la API
│   │   └── types/            # tipos TypeScript compartidos
│   └── Dockerfile
├── docker-compose.yml
├── PROMPTS.md                # log de uso de IA
└── README.md
\`\`\`

### Instalación y ejecución local

**Requisitos:** Go 1.22+, Node.js 20+, Docker (opcional).

**Backend:**
\`\`\`bash
cd backend
go mod tidy
go run ./cmd
\`\`\`
El servidor arranca en `http://localhost:8080`. Health check: `GET /health`.

**Frontend:**
\`\`\`bash
cd frontend
cp .env.example .env   # define VITE_API_URL=http://localhost:8080
npm install
npm run dev
\`\`\`
La app arranca en `http://localhost:5173`.

**Docker Compose (ambos servicios):**
\`\`\`bash
docker compose up --build
\`\`\`
Frontend: `http://localhost:3000` · Backend: `http://localhost:8080`

### Ejecutar tests

\`\`\`bash
# Backend, con cobertura
cd backend && go test ./... -v -cover

# Frontend, con cobertura
cd frontend && npm run test
\`\`\`

### Referencia de la API

**`POST /api/v1/calculate`**

| Campo       | Tipo   | Requerido | Notas                                      |
|-------------|--------|-----------|---------------------------------------------|
| `operation` | string | sí        | `add`, `subtract`, `multiply`, `divide`, `power`, `sqrt`, `percentage` |
| `a`         | number | sí        | Primer operando                             |
| `b`         | number | condicional | Requerido para operaciones binarias. No se usa en `sqrt`/`percentage`. |

Éxito — `200 OK`: `{ "result": 8 }`
Error — `400`/`422`: `{ "error": "division by zero is not allowed" }`

**Ejemplos:**
\`\`\`bash
curl -X POST http://localhost:8080/api/v1/calculate \\
  -H "Content-Type: application/json" \\
  -d '{"operation":"add","a":5,"b":3}'
# → {"result":8}

curl -X POST http://localhost:8080/api/v1/calculate \\
  -H "Content-Type: application/json" \\
  -d '{"operation":"sqrt","a":16}'
# → {"result":4}

curl -X POST http://localhost:8080/api/v1/calculate \\
  -H "Content-Type: application/json" \\
  -d '{"operation":"divide","a":10,"b":0}'
# → {"error":"division by zero is not allowed"}
\`\`\`

### Decisiones de diseño y supuestos

- **Se eligió Go para el backend** según la preferencia indicada en el enunciado.
- **Arquitectura en capas** (`calculator` → `service` → `handlers`): la capa de
  servicio existe como interfaz consumida por el handler HTTP, permitiendo
  inyección de dependencias y tests aislados del handler vía un mock — sin capa de
  `Repository/Model`, ya que este proyecto no tiene persistencia.
- **`b` como `*float64` (puntero)** en la capa de dominio distingue "operando no
  provisto" (`nil`) de "operando es cero" (`&0.0`) — relevante para división y para
  validar operaciones unarias vs. binarias.
- **Percentage** se implementó como operación unaria (`a / 100`), ya que el
  enunciado no especificó semántica de porcentaje con dos operandos.
- **CORS** restringido a un único origen configurable (variable `FRONTEND_ORIGIN`)
  en vez de un wildcard.
- **Encadenamiento de operaciones** (`5 + 3 + 2` sin presionar `=`) imita el
  comportamiento estándar de una calculadora física.
- **Historial de operaciones** (últimas 10, solo del lado del cliente) se agregó
  como mejora de UX más allá de los requisitos base.

### Despliegue

- **Backend:** [Render](https://render.com), servicio web basado en Docker usando
  `backend/Dockerfile`. Variables de entorno: `FRONTEND_ORIGIN` (URL de Netlify),
  `PORT` (definida automáticamente por Render).
- **Frontend:** [Netlify](https://netlify.com), build desde `frontend/` con
  `VITE_API_URL` apuntando a la URL del backend en Render al momento del build.

### Uso de IA

Construido con asistencia de Claude (Anthropic). Ver [`PROMPTS.md`](./PROMPTS.md)
para un resumen de cómo se usó la IA durante el desarrollo.
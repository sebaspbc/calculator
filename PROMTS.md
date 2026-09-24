# AI Usage Log

*[English version below / Versión en español más abajo]*

---

## English

Claude (Anthropic) was used as a development assistant throughout the project, in
an iterative flow of design → implementation → debugging → refinement. Below is an
organized summary, by stage, of the prompts used and how each response was applied.

### 1. Architecture and design decisions

Consulted on folder structure for a full-stack project with a Go backend (as
required by the assignment) and a React/TypeScript frontend, deployed on two
separate free-tier services.

Also evaluated whether to introduce a Java-style "manager" service layer
(`Controller → Service → Repository → Model`). The recommendation was to use Go's
idiomatic pattern instead: a service interface (`CalculatorService`) enabling
dependency injection and isolated handler testing, without adding a
`Repository`/`Model` layer — since this project has no persistence, that layer
would have added complexity without real benefit.

### 2. Backend implementation (Go)

Generated the domain layer (`internal/calculator`) with arithmetic operations and
error handling via sentinel error variables (`ErrDivisionByZero`, etc.), the service
layer (`internal/service`) as an interface-based wrapper, and the HTTP handlers
(`internal/handlers`) with input validation and error-to-status-code mapping.

Also requested a table-driven unit testing strategy (Go's standard pattern) covering
edge cases: division by zero, negative square root, missing operand, unknown
operation.

### 3. Frontend implementation (React + TypeScript)

Built the `Calculator` component with state management for chained operations (e.g.
`5 + 3 + 2` without pressing `=` in between), a separate HTTP service layer
(`services/api.ts`), and shared types (`Operation`, `CalculateRequest`) to keep
consistency with the backend contract.

Consulted on testing with Vitest + Testing Library, mocking the network layer to
isolate component tests from backend availability.

### 4. Real debugging during development

Several errors were resolved with Claude's direct help, identifying the root cause
from actual error messages:

- Empty `tsconfig.json` causing Vite to fail on startup.
- `r.use` vs `r.Use` typo (Go is case-sensitive) breaking the backend build.
- CORS error: the default `FRONTEND_ORIGIN` value pointed to `localhost:3000`
  instead of `localhost:5173` (Vite's actual port), causing the `OPTIONS` preflight
  to return 200 without the `Access-Control-Allow-Origin` header.
- TypeScript error in Docker (`Property 'env' does not exist on type 'ImportMeta'`)
  that didn't show up in dev mode because `vite dev` doesn't run strict type
  checking, but `tsc -b` does during the production build — fixed by adding
  `vite-env.d.ts` with the `vite/client` type reference.

### 5. Docker and deployment

Requested help designing multi-stage Dockerfiles: backend (build in `golang:alpine`
+ runtime in plain `alpine`, static binary via `CGO_ENABLED=0`) and frontend (build
in `node:alpine`, served with `nginx:alpine`), along with a `docker-compose.yml` to
orchestrate both services locally with a single command.

For the actual deployment, consulted on free hosting options for a Go backend and a
React frontend deployed separately, resulting in Render (backend, using the
Dockerfile directly) and Netlify (frontend, native Vite build).

### 6. Visual design and UX

Iterated on the calculator's appearance: global style reset (the body had the
browser's default margin), reordering the button grid to the standard calculator
layout, an operation history panel with an entrance animation, and error handling
as a separate banner (instead of replacing the number the user had on screen).

### Note on supervision

All generated code was reviewed, run, and validated manually before being
incorporated: tests were run (`go test ./...`, `npm run test`), behavior was
verified in the browser after each change, and real configuration/typo errors were
fixed as they appeared during local development.

---

## Español

Se usó Claude (Anthropic) como asistente de desarrollo a lo largo de todo el
proyecto, en un flujo iterativo de diseño → implementación → debugging →
refinamiento. A continuación, un resumen organizado por etapa de los prompts
usados y cómo se aplicó cada respuesta.

### 1. Arquitectura y decisiones de diseño

Se consultó sobre estructura de carpetas para un proyecto full-stack con backend en
Go (requerido por el enunciado) y frontend en React/TypeScript, con despliegue en
dos servicios gratuitos separados.

También se evaluó si convenía introducir una capa de servicio al estilo "manager"
de Java (`Controller → Service → Repository → Model`). La recomendación fue usar el
patrón idiomático de Go: una interfaz de servicio (`CalculatorService`) para
permitir inyección de dependencias y testeo aislado del handler HTTP, sin agregar
una capa de `Repository/Model` — al no haber persistencia en este proyecto, esa
capa habría sido complejidad sin propósito real.

### 2. Implementación del backend (Go)

Se generó la capa de dominio (`internal/calculator`) con las operaciones
aritméticas y manejo de errores mediante variables sentinel (`ErrDivisionByZero`,
etc.), la capa de servicio (`internal/service`) como wrapper con interfaz, y los
handlers HTTP (`internal/handlers`) con validación de entrada y mapeo de errores a
status codes.

Se pidió también una estrategia de tests unitarios tabla-driven (patrón estándar en
Go) cubriendo casos borde: división por cero, raíz de número negativo, operando
faltante, operación desconocida.

### 3. Implementación del frontend (React + TypeScript)

Se construyó el componente `Calculator` con manejo de estado para operaciones
encadenadas (ej. `5 + 3 + 2` sin presionar `=` en medio), una capa de servicio HTTP
separada de la UI (`services/api.ts`), y tipos compartidos (`Operation`,
`CalculateRequest`) para mantener consistencia con el contrato del backend.

Se consultó sobre tests con Vitest + Testing Library, mockeando la capa de red para
aislar los tests de componente de la disponibilidad real del backend.

### 4. Debugging real durante el desarrollo

Varios errores se resolvieron con ayuda directa de Claude, identificando la causa a
partir de mensajes de error reales:

- `tsconfig.json` vacío causando fallo de Vite al arrancar.
- Typo `r.use` vs `r.Use` (Go es case-sensitive) rompiendo la compilación del
  backend.
- Error de CORS: el valor por defecto de `FRONTEND_ORIGIN` apuntaba a
  `localhost:3000` en vez de `localhost:5173` (puerto real de Vite), causando que
  el preflight `OPTIONS` respondiera 200 sin el header `Access-Control-Allow-Origin`.
- Error de TypeScript en Docker (`Property 'env' does not exist on type
  'ImportMeta'`) que no aparecía en modo desarrollo porque `vite dev` no corre
  type-checking estricto, pero sí `tsc -b` en el build de producción — solucionado
  agregando `vite-env.d.ts` con la referencia a los tipos de `vite/client`.

### 5. Docker y despliegue

Se pidió ayuda para diseñar Dockerfiles multi-stage: backend (build en
`golang:alpine` + runtime en `alpine` puro, binario estático con `CGO_ENABLED=0`) y
frontend (build en `node:alpine` + servido con `nginx:alpine`), junto con
`docker-compose.yml` para orquestar ambos servicios localmente con un solo comando.

Para el despliegue real se consultó sobre opciones gratuitas para alojar un backend
en Go y un frontend en React por separado, resultando en la elección de Render
(backend, usando el Dockerfile directamente) y Netlify (frontend, con build nativo
de Vite).

### 6. Diseño visual y UX

Se iteró sobre la apariencia de la calculadora: reset de estilos globales (el body
tenía margin por defecto del navegador), reordenamiento del grid de botones a la
disposición estándar de calculadora, un historial de operaciones con animación de
entrada, y manejo de errores como un banner separado del display (en vez de
reemplazar el número que el usuario tenía en pantalla).

### Nota sobre supervisión

Todo el código generado fue revisado, ejecutado y validado manualmente antes de
incorporarse: se corrieron los tests (`go test ./...`, `npm run test`), se
verificó el comportamiento en el navegador tras cada cambio, y se corrigieron
errores reales de configuración y tipeo a medida que aparecían durante el
desarrollo local.
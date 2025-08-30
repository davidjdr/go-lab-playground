# go-lab-playground
go-lab-playground es mi espacio de pruebas en Go, un laboratorio backend donde voy integrando tecnologías, patrones y experimentos

## Roadmap de Funcionalidades

### 🏗️ **1. Servicio HTTP base + Clean/Hexagonal**
- [x] Framework HTTP (Gin/Chi)
- [ ] Arquitectura Hexagonal (ports & adapters)
- [ ] Configuración con Viper + .env
- [ ] Inyección de dependencias (Wire/Fx)
- [x] Logs estructurados (Zap/Zerolog)
- [ ] Validación (go-playground/validator)
- [ ] Manejo de errores con wrapping
- [ ] Endpoints básicos:
  - [x] `/health`
  - [ ] `/metrics`
  - [ ] `/version`

### 🗄️ **2. Persistencia con Postgres + Migraciones**
- [ ] Driver/ORM configurado (sqlc/ent/gorm)
- [ ] Migraciones automáticas (golang-migrate/goose)
- [ ] Patrón Repository implementado
- [ ] CRUD de entidades demo ("Items")
- [ ] Unit of Work (opcional)

### ⚡ **3. Caché con Redis**
- [ ] Cliente Redis configurado (go-redis/v9)
- [ ] Patrón cache-aside para lecturas
- [ ] Patrón write-through para escrituras
- [ ] Invalidación en updates
- [ ] Headers `X-Cache: HIT/MISS` en endpoints
- [ ] TTL configurables

### 📬 **4. Colas/Mensajería y Worker Asíncrono**
- [ ] NATS configurado y funcionando
- [ ] Job de ingesta: consume APIs públicas → publica mensajes
- [ ] Worker: consume → transforma → persiste en Postgres
- [ ] Invalidación de caché desde el worker
- [ ] (Futuro) RabbitMQ para comparar patrones

### 🌐 **5. Conexión a APIs Públicas + Resiliencia**
- [ ] Integración con API pública (GitHub/PokéAPI/SpaceX)
- [ ] Reintentos configurables (go-retryablehttp)
- [ ] Circuit breaker (sony/gobreaker)
- [ ] Rate limiting (golang.org/x/time/rate)
- [ ] Endpoint `/external/sync` funcional
- [ ] Endpoint `/reports/top` con datos procesados

### 📊 **6. Observabilidad Completa**
- [ ] OpenTelemetry + Jaeger para tracing
- [ ] Métricas Prometheus en `/metrics`
- [ ] Logs correlacionados con trace IDs
- [ ] Dashboards Grafana básicos
- [ ] Headers `X-Trace-Id` en respuestas

### 🔒 **7. Seguridad y Hardening**
- [ ] Autenticación JWT básica
- [ ] Rate limiting por IP/API key
- [ ] Middlewares CORS y security headers
- [ ] Gestión segura de secrets (.env)
- [ ] (Futuro) OAuth con GitHub

### 📋 **8. API Contracts + Documentación**
- [ ] OpenAPI/Swagger con swaggo
- [ ] Documentación auto-generada
- [ ] (Futuro) gRPC + grpc-gateway
- [ ] Client stubs generados

### ⏰ **9. Tareas Programadas y Sagas**
- [ ] Scheduler con robfig/cron
- [ ] Jobs programados básicos
- [ ] (Futuro) Sagas para transacciones distribuidas

### 🎛️ **10. Feature Flags y A/B Testing**
- [ ] Sistema de feature flags básico
- [ ] Toggle de funcionalidades por env
- [ ] (Futuro) Integración con Unleash
- [ ] Experimento A/B en endpoint demo

### 🚀 **11. CI/CD y Despliegue**
- [ ] GitHub Actions para testing
- [ ] Pipeline de linting (golangci-lint)
- [ ] Build de imágenes Docker multi-arch
- [ ] Deploy automático a Render/Koyeb
- [ ] Health checks configurados

### 🐳 **12. Dockerización y Orquestación**
- [ ] Dockerfile multi-stage optimizado
- [ ] docker-compose.yml completo para dev
- [ ] Hot-reload con air (opcional)
- [ ] (Futuro) Kubernetes + Helm charts

## Quickstart

1. Clona el repo y crea tu `.env` desde `example.env`.
2. `make dev-up` para levantar todo el stack local.
3. Prueba:
   - `GET http://localhost:8080/health` → OK.
   - `GET http://localhost:8080/metrics` → métricas Prometheus.
   - `POST http://localhost:8080/external/sync?source=github` → publica mensajes; el worker los procesa y guarda en Postgres.
4. Abre Jaeger en `http://localhost:16686` y busca trazas del servicio `api`.

## Scripts útiles

- `make dev-up` / `make dev-down` / `make dev-logs`
- `make sqlc`
- `make migrate-up` / `make migrate-down`

## Despliegue

### Render
- [ ] Web Service configurado con deploy automático desde GitHub
- [ ] Postgres administrado en plan Free
- [ ] Variables de entorno configuradas: `DATABASE_URL`, `REDIS_URL`, `NATS_URL`
- [ ] Health checks funcionando

**Build**: Dockerfile multi-stage  
**Variables requeridas**: `DATABASE_URL`, `REDIS_URL`, `NATS_URL`, `OTEL_EXPORTER_OTLP_ENDPOINT` (opcional), `APP_ENV`

## Stack Tecnológico

**Core**: Go 1.24+, Gin/Chi, Clean Architecture  
**Persistencia**: Postgres + sqlc/ent/gorm, golang-migrate  
**Caché**: Redis (go-redis/v9)  
**Mensajería**: NATS (futuro: RabbitMQ para comparar)  
**Observabilidad**: OpenTelemetry, Jaeger, Prometheus, Grafana  
**Resiliencia**: retryablehttp, gobreaker, rate limiting  
**Configuración**: Viper, dotenv  
**Testing**: testify, mockery  
**CI/CD**: GitHub Actions, golangci-lint
arca los checkboxes conforme vayas completando cada funcionalidad. El proyecto está diseñado para construirse incrementalmente.
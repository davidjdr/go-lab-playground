FROM golang:1.24
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Compila dos binarios: api y worker
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/api ./cmd/api \
&& CGO_ENABLED=0 GOOS=linux go build -o /out/worker ./cmd/worker

# ---- Run (distroless) ----
FROM gcr.io/distroless/base-debian12 AS run
WORKDIR /
COPY --from=build /out/api /api
COPY --from=build /out/worker /worker
COPY example.env /.env
ENV APP_ENV=production
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/api"]
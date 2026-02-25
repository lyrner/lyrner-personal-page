# syntax=docker/dockerfile:1
# Multi-stage Go build.
# The binary embeds all templates and static assets via //go:embed,
# so the final image only needs the single compiled binary.

# ---- Stage 1: Build ----
FROM golang:1.23-alpine AS build
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /lyrner .

# ---- Stage 2: Run ----
FROM alpine:3.21 AS runtime
WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=build /lyrner .

USER appuser

EXPOSE 8080

CMD ["/app/lyrner"]

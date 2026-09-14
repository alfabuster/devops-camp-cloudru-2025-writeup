# ---- Этап 1: сборка ----
FROM golang:1.23-alpine AS builder
WORKDIR /src
COPY go.mod ./
COPY main.go ./
# CGO_ENABLED=0 — статическая линковка, чтобы бинарник
# жил в scratch без единой системной библиотеки.
# -ldflags="-s -w" срезает отладочные символы (меньше вес).
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /echo-server .
RUN echo 'appuser:x:10001:10001::/:' > /etc/passwd

# ---- Этап 2: финальный образ ----
FROM scratch
COPY --from=builder /echo-server /echo-server
COPY --from=builder /etc/passwd /etc/passwd
USER appuser
EXPOSE 8000
ENV AUTHOR="alfabuster"
ENTRYPOINT ["/echo-server"]
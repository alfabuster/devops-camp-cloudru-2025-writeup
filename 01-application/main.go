package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
)

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// IP контейнера в его собственном network namespace.
func getContainerIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "unknown"
	}
	defer conn.Close()
	return conn.LocalAddr().(*net.UDPAddr).IP.String()
}

func handler(w http.ResponseWriter, r *http.Request) {
	containerHost, _ := os.Hostname()   // имя из UTS namespace = ID контейнера
	containerIP := getContainerIP()     // адрес из network namespace

	// Данные ХОСТА не добываются — они приходят снаружи.
	hostHost := envOr("HOST_HOSTNAME", "(not injected)")
	hostIP := envOr("HOST_IP", "(not injected)")
	author := envOr("AUTHOR", "anonymous")

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<!DOCTYPE html>
<html><head><title>Echo Server</title></head>
<body style="font-family:monospace">
  <h1>Echo Server</h1>
  <ul>
    <li><b>Container hostname:</b> %s</li>
    <li><b>Container IP:</b> %s</li>
    <li><b>Host hostname:</b> %s</li>
    <li><b>Host IP:</b> %s</li>
    <li><b>Author:</b> %s</li>
  </ul>
</body></html>`, containerHost, containerIP, hostHost, hostIP, author)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/healthz", healthz)
	http.ListenAndServe(":8000", nil)
}
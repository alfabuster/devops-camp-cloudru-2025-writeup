#!/usr/bin/env python3
import os
import socket
from http.server import BaseHTTPRequestHandler, ThreadingHTTPServer


def env_or(key, fallback):
    return os.getenv(key) or fallback


def get_container_ip():
    s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    try:
        s.connect(("8.8.8.8", 80))
        return s.getsockname()[0]
    except Exception:
        return "unknown"
    finally:
        s.close()


class Handler(BaseHTTPRequestHandler):
    def do_GET(self):
        if self.path == "/healthz":
            self.send_response(200)
            self.send_header("Content-Type", "text/plain")
            self.end_headers()
            self.wfile.write(b"ok\n")
            return

        container_host = socket.gethostname()      # ID контейнера
        container_ip = get_container_ip()
        host_host = env_or("HOST_HOSTNAME", "(not injected)")
        host_ip = env_or("HOST_IP", "(not injected)")
        author = env_or("AUTHOR", "anonymous")

        body = f"""<!DOCTYPE html>
<html><head><title>Echo Server</title></head>
<body style="font-family:monospace">
  <h1>Echo Server (Python)</h1>
  <ul>
    <li><b>Container hostname:</b> {container_host}</li>
    <li><b>Container IP:</b> {container_ip}</li>
    <li><b>Host hostname:</b> {host_host}</li>
    <li><b>Host IP:</b> {host_ip}</li>
    <li><b>Author:</b> {author}</li>
  </ul>
</body></html>"""

        self.send_response(200)
        self.send_header("Content-Type", "text/html; charset=utf-8")
        self.end_headers()
        self.wfile.write(body.encode("utf-8"))

    def log_message(self, *args):
        pass


if __name__ == "__main__":
    ThreadingHTTPServer(("0.0.0.0", 8003), Handler).serve_forever()
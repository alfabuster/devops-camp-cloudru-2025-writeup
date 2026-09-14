```bash
# 1. Логинимся (введёт логин и Access Token — не пароль, пароли уже немодно)
docker login -u <ваш_логин>

# 2. Собираем с правильным тегом: <логин>/<репозиторий>:<версия>
docker build -t <ваш_логин>/echo-server:1.0 .

# 3. Проверяем локально, что оно вообще дышит
docker run --rm -p 8001:8000 \
  -e AUTHOR="alfabuster" \
  -e HOST_HOSTNAME="$(hostname)" \
  -e HOST_IP="$(ip -4 route get 1.1.1.1 | grep -oP 'src \K\S+')" \
  <ваш_логин>/echo-server:1.0
#    открываем http://localhost:8001 — должны увидеть hostname/IP/author

# 4. Пушим в приватный регистри
docker push <ваш_логин>/echo-server:1.0
```
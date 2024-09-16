go test ./...

curl -s -o /dev/null -w "%{http_code}" http://localhost:8080/ping
if [ $? -eq 0 ]; then
  echo "Сервер доступен"
else
  echo "Сервер недоступен"
  exit 1
fi

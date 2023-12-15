### Сборка docker образа

```
docker build -t test-docker-service:1.0 .
```

### Запуск контейнера

```
docker run -p 8000:8000 test-docker-service:1.0
```
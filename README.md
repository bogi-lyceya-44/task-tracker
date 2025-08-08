# Task Tracker

Бэкенд для таск трекера. Таски, топики, борды - все для канбанщины.

---

Как запустить:
1. Установите себе docker и docker-compose

2. Сгенерируйте PAT (personal access token) в настройках гитхаба
    - Settings -> в самом низу Developer settings -> Personal access tokens
    - Выбирайте Tokens (classic), а затем Generate new token (classic)
    - Возьмите нужные разрешения (если не знаете какие, то возьмите все xd)
3. Полученный токен сохраните

4. В терминале введите
```bash
# вместо <token> вставьте свой personal access token
# вместо <username> вставьте свой юзернейм на гитхабе
$ echo "<token>" | docker login ghcr.io --username <username> --password-stdin
```

5. Теперь, чтоб запустить таск трекер, введите
```bash
$ docker-compose up -d
# или
$ docker compose up -d
```
(тут как получится)

6. Установите себе goose для применения миграций
```bash
$ brew install goose
# или
$ go install github.com/pressly/goose/v3/cmd/goose@latest
```

7. Чтоб применить миграции, введите
```bash
$ goose -dir migrations postgres "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable" up
```
Ну или поменяйте под себя DSN, если вы трогали `docker-compose.yaml`.

8. По дефолту сервис должен был запуститься на `localhost:7000` для HTTP и `localhost:7001` для gRPC.
По адресу `localhost:7000/docs` или `localhost:7000/swagger/` можно найти сваггер док.

9. Роняем контейнеры:
```bash
$ docker-compose down # роняет все контейнеры без волюмов (т.е. в нашем случае постгрес сохранит информацию)
$ docker-compose down -v # роняет все контейнеры с волюмами (очищаем постгрес)
$ docker-compose down -v postgres # роняет только постгрес с волюмом
```
и так далее.

10. Логи контейнеров:
```bash
$ docker-compose logs -f
```

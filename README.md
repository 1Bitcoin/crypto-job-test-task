### Запуск приложения с дефолтными параметрами

```bash
git clone ...
```

```bash
docker-compose up -d
```

### Запуск приложения с передачей флагов

Вы также можете передать параметры подключения через флаги:

- `-host` - хост базы данных
- `-port` - порт базы данных
- `-user` - пользователь базы данных
- `-password` - пароль базы данных
- `-name` - имя базы данных
- `-sslmode` - режим SSL


```bash
git clone ...
```

```bash
docker-compose run --rm app ./app -host=db -port=5432 -user=myuser -password=mypassword -name=postgres -sslmode=disable
```

### Тестовое задание

https://gist.github.com/nanaban/27e482f75357e53c2014beab6cea498b
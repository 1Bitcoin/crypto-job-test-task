### Запуск приложения с дефолтными параметрами

```bash
git clone https://github.com/1Bitcoin/crypto-job-test-task.git
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
git clone https://github.com/1Bitcoin/crypto-job-test-task.git
```

```bash
docker-compose run --rm app ./app -host=db -port=5432 -user=myuser -password=mypassword -name=postgres -sslmode=disable
```

### Методы микросервиса

Имеется два метода

1) Healthcheck - без параметров
2) GetRates - принимает в качестве параметра marketID
Для демонстрации работы можно использовать "marketID": "usdtrub"


### Тестовое задание

https://gist.github.com/nanaban/27e482f75357e53c2014beab6cea498b
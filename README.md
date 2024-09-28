### Флаги командной строки

Вы также можете передать параметры подключения через флаги:

- `-host` - хост базы данных
- `-port` - порт базы данных
- `-user` - пользователь базы данных
- `-password` - пароль базы данных
- `-name` - имя базы данных
- `-sslmode` - режим SSL

#### Пример запуска приложения с флагами:

```bash
go run cmd/service/main.go -host=localhost -port=5432 -user=myuser -password=mypassword -name=postgres -sslmode=disable

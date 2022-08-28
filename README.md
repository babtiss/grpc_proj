[![Typing SVG](https://readme-typing-svg.herokuapp.com?color=%2336BCF7&lines=Тестовое+задание+Golang)](https://git.io/typing-svg)

### Структура сервиса.
- gRPC сервис построен на основе proto файла на Go.
- Брокер сообщений: Kafka
- Кеширование: Redis
- СУБД: PostgreSQL, ClickHouse

### Описание ручек:
- `Add` : добавить пользователя в бд + сделать лог в clickhouse через kafka; кешируется в redis.
- `Delete`: удалить пользователя из бд; кешируется в redis.
- `GetClients`: получить список всех пользователей из бд.
- `GetClientsFromRedis`: получить список всех пользователей из redis за последнюю минуту.

### Как запустить у себя
- Клонируем репозиторий: `git clone https://github.com/babtiss/grpc_proj`
- Устанавливаем окружение. Поднимаем нужные ресурсы. [Тут пример как это делал я](https://github.com/babtiss/grpc_proj/tree/master/assets)
- Настраиваем конфиг для postgres в environments (`host=;port=;user=;password=;dbname=;sslmode=`)
- Запускаем `cmd/server/main.go` с нашими настройками.
- Для тестирования работы ручек используем утилиту `evans` из корня проекта:
```
evans .\proto\clientbase.proto -p 8000
```
- Дергаем ручки командой `call < имя ручки > `
- Смотрим в логи терминала и развернутых серверов что всё работает успешно.

### TODO (как можно улучшить):
- собрать по контейнерам в docker/docker-compose
- написать скрипт(батник или баш) для быстрого развертывания

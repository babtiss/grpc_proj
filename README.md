[![Typing SVG](https://readme-typing-svg.herokuapp.com?color=%2336BCF7&lines=Тестовое+задание+Golang)](https://git.io/typing-svg)

### Структура сервиса.
- gRPC сервис построен на основе proto файла на Go.
- Брокер сообщений: Kafka
- Кеширование: Redis
- СУБД: PostgreSQL, ClickHouse

### Запуск:

1. Через `Docker`
- `git clone https://github.com/babtiss/grpc_proj`
- Настраиваем `docker-compose.yml` и `configs/config.yml` под себя.
- `docker-compose up --build grpc_proj`
- Создать таблицы для СУБД (PostgreSQL, ClickHouse).

2. Локально руками
- Поднимаем окружение. Как это сделать: [тут](https://github.com/babtiss/grpc_proj/tree/master/assets).
- Запускаем `main.go` с env: `RUN_MODE='local'`

> TODO: докеризировать init database && table
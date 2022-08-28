# Используемые ресурсы и их функционал
> TODO: докеризировать сборку

## Protobuf

#### Генерация pb файлов в наш проект (Windows).
Из корня проекта выполнить:
```
protoc -I proto --go_out=pkg/api --go-grpc_out=pkg/api clientbase.proto
```

## kafka + zookeeper (на виртуальную Ubuntu):
[Устанавливаем](https://www.digitalocean.com/community/tutorials/how-to-install-apache-kafka-on-ubuntu-20-04)

Список команд:
#### Запуск zookeeper:
```
~/kafka/bin/zookeeper-server-start.sh ~/kafka/config/zookeeper.properties
```
#### Запуск kafka:
```
~/kafka/bin/kafka-server-start.sh ~/kafka/config/server.properties
```
#### Создать тему (необяз - команда чтоб потестить)
```
~/kafka/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test
```
#### Запустить продюсера (необяз - команда чтоб потестить)
```
~/kafka/bin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic test
```

#### Запустить потребителя (необяз - команда чтоб потестить)
```
~/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test
```


## ClickHouse (на виртуальную Ubuntu):

[Качаем и собираем файл по туториалу](https://clickhouse.com/docs/en/quick-start)
[Или по другому туториалу](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-clickhouse-on-debian-10-ru)

## СУБД
В postgres и clickhouse создаем бд и таблицы:

Для sql использвал схему:
```
CREATE TABLE IF NOT EXISTS Client
(
name String
)
```
Для clickhouse добавляем Engine
```
CREATE TABLE IF NOT EXISTS Client
(
    name String
) ENGINE = Kafka(`localhost:9092`, `test`, 'test', 'JSONEachRow')
```
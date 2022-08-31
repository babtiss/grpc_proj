# Окружение

## Kafka + Zookeeper (на виртуальную Ubuntu):
[Устанавливаем](https://www.digitalocean.com/community/tutorials/how-to-install-apache-kafka-on-ubuntu-20-04)

#### Запуск zookeeper:
```
~/kafka/bin/zookeeper-server-start.sh ~/kafka/config/zookeeper.properties
```
#### Запуск kafka:
```
~/kafka/bin/kafka-server-start.sh ~/kafka/config/server.properties

#### Если нужно потестить в работу в cmd:
```
Создать тему
~/kafka/bin/kafka-topics.sh --create --bootstrap-server localhost:9092 --replication-factor 1 --partitions 1 --topic test

Запустить продюсера
~/kafka/bin/kafka-console-producer.sh --bootstrap-server localhost:9092 --topic test

Запустить потребителя
~/kafka/bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic test
```

## ClickHouse (на виртуальную Ubuntu):
[Устанавливаем](https://clickhouse.com/docs/en/quick-start)
#### Запуск ClickHouse:
```
cd Clickhouse/build/programs/
./clickhouse server
```

## Redis (Windows):
[Устанавливаем](https://redis.io/docs/getting-started/installation/install-redis-on-windows/)

## СУБД
В postgres и clickhouse создаем бд и таблицы:

Для sql схема:
```
CREATE TABLE IF NOT EXISTS Client
(
    name String
)
```

Для clickhouse добавляем Engine:
```
CREATE TABLE IF NOT EXISTS Client
(
    name String
) ENGINE = Kafka(`localhost:9092`, `test`, 'test', 'JSONEachRow')
```

## Protobuf

#### Генерация pb файлов в наш проект если меняли апи (Windows).
```
protoc -I proto --go_out=pkg/api --go-grpc_out=pkg/api clientbase.proto
```

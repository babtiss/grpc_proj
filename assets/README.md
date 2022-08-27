## Всё что может понадобится для работы


### Protobuf

#### Генерация pb файлов в наш проект (Windows).
Из корня проекта выполнить:
```
protoc -I proto --go_out=pkg/api --go-grpc_out=pkg/api clientbase.proto
```

### kafka + zookeeper (на виртуальную Ubuntu):
#### Установка:
```
wget https://archive.apache.org/dist/kafka/3.0.0/kafka_2.13-3.0.0.tgz
tar xzf kafka_2.13-3.0.0.tgz
mv kafka_2.13-3.0.0 ~
mv kafka_2.13-3.0.0 kafka
```
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


### ClickHouse (на виртуальную Ubuntu):

[качаем и собираем файл по туториалу](https://clickhouse.com/docs/en/quick-start)

#### clickhouse run:
```
cd Clickhouse/build
./clickhouse server
```
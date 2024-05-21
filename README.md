SQL-запросы для ClickHouse из 1 задачи описаны в файле task_1.sql

Вставка тестовых данных в таблицу events и вывод событий по заданному eventType 
и временному диапазону из 2 задачи выполнены в файле task_2.go

Пример файла конфигурации находится по пути config/config.yml.example

Запустить сервис:

docker-compose up (-d)

Провести миграции:

docker compose run --rm migrations

Создание миграции (создает сразу пару up и down):
migrate create -ext sql -dir migrations {migration_name}

Применение миграции:
migrate -verbose -source=file://migrations -database 'postgres://telegrambot:botpassword@localhost:5432/telegrambot?sslmode=disable&search_path=public&x-migrations-table=schema_migrations' up

Получение версии схемы миграции:
migrate -verbose -source=file://migrations -database 'postgres://telegrambot:botpassword@localhost:5432/telegrambot?sslmode=disable&search_path=public&x-migrations-table=schema_migrations' version

Откатить миграции:
migrate -verbose -source=file://migrations -database 'postgres://telegrambot:botpassword@localhost:5432/telegrambot?sslmode=disable&search_path=public&x-migrations-table=schema_migrations' down

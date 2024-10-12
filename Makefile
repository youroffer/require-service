run:
	docker compose up --build
	
# Команда для генерации дерева зависимостей
deep:
	dep-tree entropy cmd/main.go

# Команда для генерации Swagger документации
.PHONY: swag
swag:
	swag init --generalInfo cmd/main.go



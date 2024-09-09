.PHONY: run seed reset-db

# Jalankan aplikasi
run:
	go run cmd/web/main.go

# Seed data dummy
seed:
	go run cmd/web/main.go seed

# Reset database
reset-db:
	go run cmd/web/main.go reset-db
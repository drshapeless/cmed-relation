include .env

demo:
	go run cmd/demo/main.go

parser:
	go run cmd/parser/main.go

templ:
	templ generate -path internal/templates/

sql_reset:
	$(MAKE) -C migrations reset

include .env

run:
	go run main.go

sql_reset:
	$(MAKE) -C migrations reset

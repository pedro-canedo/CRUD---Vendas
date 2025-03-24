.PHONY: run swagger test

# Executa a aplicação
run:
	go run cmd/main.go

# Gera a documentação Swagger
swagger:
	~/go/bin/swag init -g cmd/main.go

# Executa os testes
test:
	go test ./...

# Gera a documentação e executa a aplicação
dev: swagger run

# Limpa arquivos temporários
clean:
	rm -rf docs/docs.go docs/swagger.json docs/swagger.yaml 
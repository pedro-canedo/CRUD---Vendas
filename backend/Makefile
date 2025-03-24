.PHONY: run swagger test seed dev clean

# Executa a aplicação
run:
	go run cmd/main.go

# Gera a documentação Swagger
swagger:
	~/go/bin/swag init -g cmd/main.go

# Executa os testes
test:
	go test ./...

# Carrega dados iniciais
seed:
	go run cmd/seed/main.go

# Gera a documentação e executa a aplicação
dev: swagger seed run

# Limpa arquivos temporários
clean:
	rm -rf docs/docs.go docs/swagger.json docs/swagger.yaml
	rm -f vendas.db 
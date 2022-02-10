PACKAGE=github.com/MarceloRezin/go-env
VERSAO=v0.2.0

test:
	@ echo "Executando testes de unidade"
	go test ./...

deploy:
	@ echo "Realizando deploy"
	git add .
	git commit -m "Mudanças para versão $(VERSAO)"
	git tag $(VERSAO)
	git push origin $(VERSAO)
	GOPROXY=proxy.golang.org go list -m $(PACKAGE)@$(VERSAO)
	git push
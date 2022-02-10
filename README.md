# go-env
Utilitário em golang para ler varáveis de ambiente do arquivo .env

### Arquivo env válido:
```
minhaVar=meu valor
meuTeste=meu dado
```
### Utilização:
```
envMap, err := goenv.Read("dir/para/envfile")
value := envMap["minhaVar"]
```

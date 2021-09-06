# abc_brasil
Lib de integração com a API do ACB Brasil Banco

## Baixando
```shell script
go get github.com/Fix-Pay/abc-brasil
```

````shell script
go env -w GOPRIVATE=github.com/Fix-Pay/abc-brasil

git config --global url."https://${username}:${access_token}@github.com".insteadOf / "https://github.com"
````

### Sonarqube
Chave: dbda89652dd9749ee3505f238296bc091dea1efe

```
sonar-scanner \
  -Dsonar.projectKey=abc-brasil \
  -Dsonar.sources=. \
  -Dsonar.host.url=http://sonarqube.fixpay.com.br:9001 \
  -Dsonar.login=dbda89652dd9749ee3505f238296bc091dea1efe
```
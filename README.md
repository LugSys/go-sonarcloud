## SonarQube com Golang

### Executar container Golang
```
docker run -it --rm -v ./go/:/app golang:1.19 bash
```

### Gerar arquivo de cobertura de testes
```
go test -coverprofile=coverage.out
```

### Executar o SonarQube

```
docker-compose up -d
```

### Executar o sonar scanner

Neste projeto optei por não utilizar a criação do arquivo `sonar-project.properties` e os parâmetros estão sendo passados via `SONARQUBE_SCANNER_PARAMS` (parâmetros em formato JSON) no comando `docker run`, incluso no arquivo `scan.sh` na pasta `./sonarqube`.

Para executar o `sonar scanner` entre na pasta `./sonarqube`, execute o comando `chmod +x scan.sh` para conceder permissão de execução em seguida `./scan.sh` e a análise será executada. Nota: O `SonarQube` container deve estar em execução, veja o comando acima para executá-lo.

> Nota: Neste exemplo o Golang, SonarQube e scanner estão todos executados em containers.

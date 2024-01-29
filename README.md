# Go Backend Template

## Create docker images

```
docker build -t go-starter:email -f ./build/Dockerfile.email .
docker build -t go-starter:server -f ./build/Dockerfile.server .
docker build -t go-starter:worker -f ./build/Dockerfile.worker .
```

## Check for updates

```
go list -m -u all
```

## Run tests

```
go test -v -cover ./...
```

## Run security checks

```
gosec ./...
```

## Lint

```
staticcheck ./...
```

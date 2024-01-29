```
docker build -t go-starter:cli -f ./build/Dockerfile.cli .
docker build -t go-starter:server -f ./build/Dockerfile.server .
docker build -t go-starter:worker -f ./build/Dockerfile.worker .
```

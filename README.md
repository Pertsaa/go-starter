```
docker build -t go-starter:email -f ./build/Dockerfile.email .
docker build -t go-starter:server -f ./build/Dockerfile.server .
docker build -t go-starter:worker -f ./build/Dockerfile.worker .
```

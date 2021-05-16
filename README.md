# Kratos Project Template

## Install Kratos
```
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# create a template project
kratos new helloworld

cd helloworld
# Add a proto template
kratos proto add api/helloworld/helloworld.proto
# Generate the source code of service by proto file
kratos proto server api/helloworld/helloworld.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/helloworld -conf ./configs
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

# ent
ent init User
go generate ./ent
```

# deploy
- docker run -p 6379:6379 --name redis -d redis
- docker run -p 3306:3306 --name mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.7
echo -------- Build Doc --------
swag init -g app/cmd/main.go

echo -------- Build Server --------

mkdir build
cp -r conf build/conf
GOOS=linux go build -o build/wusthelper-todolist-linux-amd64 wusthelper-todolist-service/app/cmd
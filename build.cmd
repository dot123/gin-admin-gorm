swag fmt -g .\cmd\GameAdmin\main.go --exclude ./pkg/rabbitMQ
swag init -g .\cmd\GameAdmin\main.go
cd .\cmd\GameAdmin
wire
pause
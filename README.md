# todo

ENV DATABASE_URL="postgresql://postgres:123456@localhost:5432/postgres"
ENV HTTP_ADDRESS="0.0.0.0:3000" 

Локальный запуск
1) go get -u
2) HTTP_ADDRESS="0.0.0.0:3000" DATABASE_URL="postgresql://postgres:123456@localhost:5432/postgres" go run ./main.go
3) Проверить работоспособность сервиса отправив запрос

Запуск docker file
1) Указать ENV в dockerFile
2) docker build -t todo .
3) docker run -p 3000:3000 todo
4) Проверить работоспособность сервиса
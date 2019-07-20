##Run DB
docker run -d -p 5433:5432 \
--name postgres_todo \
-e POSTGRES_PASSWORD=mysecretpassword \
-e POSTGRES_USER=postgres \
-e POSTGRES_DB=todo \
postgres:9.6.1
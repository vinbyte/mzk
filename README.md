# Prerequisites
- Golang installed
- Postgres installed
- [DBMate](https://github.com/amacneil/dbmate) installed. Or you can use dbmate via docker instead (using `make` command)
- Docker (optional, if you want to run via docker instead)

# How to run

## Without docker
1. Clone the repo
1. Copy and paste `.env.sample` to `.env`. And the config in `.env` file.
1. Make sure a database in ready in your Postgres
1. Run a migrations using dbmate to create a table and seed data. `dbmate -u "<database_url>" up`. Example : `dbmate -u "postgres://postgres:@localhost:5432/postgres" up`. Make sure the database is same as defined in `.env` file
1. Run `make dev` to run the server. It will running on port `6060` by default.
![make_dev img](/assets/make_dev.png)
1. The API docs is available via Swagger. You can open it in http://localhost:6060/swagger/index.html
![swagger img](/assets/swagger.png)
1. You can try directly in Swagger.
![swagger-response img](/assets/swagger-response.png)

## With docker

1. Clone the repo
1. Run your docker
1. Run `make compose-up`. It will prepare all you need to run the app. Spin up a Postgres container, build the app, and do a db migration.
![docker img](/assets/docker-compose-up.png)
1. The API docs is available via Swagger. You can open it in http://localhost:6060/swagger/index.html
![swagger img](/assets/swagger.png)
1. You can try directly in Swagger.
![swagger-response img](/assets/swagger-response.png)
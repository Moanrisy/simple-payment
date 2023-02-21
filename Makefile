API_URL=127.0.0.1

DB_DRIVER=postgres
# localhostr
DB_HOST=127.0.0.1
DB_PORT=5432
DB_USER=postgres
DB_PASS=password
DB_NAME=simple_payment
DB_SSL_MODE=disable
POSTGRESQL_URL='postgres://postgres:password@127.0.0.1:5432/tabungin_aja?sslmode=disable'

run:
	@echo "Running go gin..."
	swag init
	@env API_URL=${API_URL} DB_HOST=${DB_HOST} DB_PORT=${DB_PORT} DB_USER=${DB_USER} DB_PASS=${DB_PASS} DB_NAME=${DB_NAME} DB_DRIVER=${DB_DRIVER} DB_SSL_MODE=${DB_SSL_MODE} \
	air
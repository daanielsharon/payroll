module auth

go 1.24.1

require github.com/go-chi/chi/v5 v5.2.1

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/golang-jwt/jwt/v4 v4.5.2 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	gorm.io/datatypes v1.2.5 // indirect
	gorm.io/driver/mysql v1.5.6 // indirect
	gorm.io/gorm v1.30.0 // indirect
)

require (
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/text v0.21.0 // indirect
	shared v0.0.0
)

replace shared => ../shared

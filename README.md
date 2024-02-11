## Setup

### Step 1. Copy .env.example as .env

### Step 2. Edit DB/DB_URL fields in .env file

#### Sqlite

``` ini
DB=sqlite
DB_URL=data.db
```

#### Postgres

``` ini
DB=postgres
DB_URL=postgresql://user:pass@host:5432/dbname
```

### Step 3. Start Server

``` sh
go run .
```

## API

```
POST /api/auth/signup { username, password }
POST /api/auth/signin { username, password }
GET /api/v1/notes?offset=0&limit=20
GET /api/v1/notes/1
POST /api/v1/notes { content, visible(public/private) }
......
```

## Packages

### Gin (Web Framework)

- https://gin-gonic.com/docs/quickstart/
- https://github.com/gin-gonic/gin (74.1k stars)

``` sh
go get -u github.com/gin-gonic/gin
```

### Gorm (ORM)

- https://gorm.io/
- https://github.com/go-gorm/gorm (34.8k stars)

``` sh
go get -u gorm.io/gorm 
```


### godotenv (Configuration)

- https://github.com/joho/godotenv (7.2k stars)

``` sh
go get github.com/joho/godotenv
```


### JWT

- https://github.com/golang-jwt/jwt (6k stars)

``` sh
go get -u github.com/golang-jwt/jwt/v5
```
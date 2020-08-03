## BOOK-LIST

Simple CRUD APIS for Books using POSTGRES SQL and Gorilla Mux

## Installation

### Downloading packages

1. Clone the repository
2. `go get github.com/subosito/gotenv`
2. `go get github.com/gorilla/mux`
4. `go get github.com/lib/pq`

### Setting Up Env and Database

1. Create account on [Elephant Sql](https://customer.elephantsql.com/instance).
2. Create db Instance on above url.
3. Copy the `database url` which will be in the format `postgres://user:password@server:PORT/user`.
4. Create `.env` file locally in the root directory.
5. Add `SQL_URL=database url` in the file.

### Running Locally
1. `go run main.go`
2. Open Postman
3. Import `book-list.postman_collection.json`
4. Enjoy

package fuzion

/*
Make a application that uses docker compose with the local Dockerfile and entrypoint
to build the go application and the mysql image.

GO application:

Make a function that connects and initializes a database connection
in the file (mysql/mysql.go) and import it as (github.com/PashaVDW/mysql : check the docs or go.mod file) and use it as mysql.{FUNC_NAME}

Make a function that can register and login a user, using the database

Usage:
Via docker runnen en fmt.Scan (of andere manier) gebruiken om de username & password (masked input) op te vragen (PASOP: dit is lastig voor strings)

GOOGLE GEBRUIKEN!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!
*/

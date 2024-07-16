package db

import(
		"log"
		"database/sql"
	_ "github.com/lib/pq"
)

func ConectaComDB() *sql.DB {
	conexao := "user= dbname=a password= host=localhost port= sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	return db
}

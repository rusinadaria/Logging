package repository

import (
	_"github.com/lib/pq"
	"database/sql"
	"log"
	"log/slog"
	"os"
	// "golang.org/x/crypto/bcrypt"
)

var database *sql.DB

func ConnectDatabase(logger *slog.Logger) {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Error("Failed connect database")
        panic(err)
    } 
	logger.Debug("Connect database")
	database = db
	
    // defer db.Close()
}

func SaveRefreshToken(guid int, refreshToken string) {
	// hash_token := bcrypt.
	//захэшировать пароль перед добавлением
	rows, err := database.Query("INSERT INTO tokens (user_guid, refresh_token) VALUES ($1, $2)", guid, refreshToken)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
}
package db

import (
    "fmt"
    "os"
    "log"
    "wallet-service/models"
    "github.com/jmoiron/sqlx"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sqlx.DB

func InitDB() {
    _ = godotenv.Load()

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

    var err error
    DB, err = sqlx.Connect("postgres", dsn)
    if err != nil {
        log.Fatalf("DB connection error: %v", err)
    }

    DB.MustExec(`CREATE TABLE IF NOT EXISTS wallets (
        id SERIAL PRIMARY KEY,
        user_id TEXT UNIQUE NOT NULL,
        address TEXT UNIQUE NOT NULL,
        created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
    )`)
}

func SaveWallet(w models.WalletResponse) error {
    _, err := DB.Exec("INSERT INTO wallets (user_id, address) VALUES ($1, $2)", w.UserID, w.Address)
    return err
}

func GetWalletByUserID(userID string) (*models.WalletResponse, error) {
    var w models.WalletResponse
    err := DB.Get(&w, "SELECT user_id, address FROM wallets WHERE user_id = $1", userID)
    if err != nil {
        return nil, err
    }
    return &w, nil
}

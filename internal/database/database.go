package database

import (
	"crypto/md5"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewConnection() error {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "sso_testing"

	var err error
	DB, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	if err := DB.Ping(); err != nil {
		return fmt.Errorf("error connecting to the database: %v", err)
	}

	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}

func InsertSSOIntegrationMapping(ssoUsername string, aclUsername string, uuid string) error {
	token := GenerateToken()
	findUser, err := DB.Query("SELECT * FROM sso_integration_mapping WHERE ssousername = ?", ssoUsername)
	if err != nil {
		return fmt.Errorf("insertSSOIntegrationMapping: %v", err)
	}
	if findUser.Next() {
		_, err := DB.Query("INSERT INTO api_token (`username`, `token`) VALUES (?, ?)", aclUsername, token)
		if err != nil {
			return fmt.Errorf("error inserting api token: %v", err)
		}
		_, err = DB.Query("INSERT INTO acl_qr (`uuid`, `token`, `device_name`) VALUES (?, ?, ?)", uuid, token, uuid)
		if err != nil {
			return fmt.Errorf("error inserting acl qr: %v", err)
		}
	}
	return nil
}

func GenerateToken() string {
	uniqueID1 := generateUniqID()
	hash1 := md5Hash(uniqueID1)

	// Generate second unique ID
	uniqueID2 := generateUniqID()
	hash2 := md5Hash(uniqueID2)

	// Combine both hashes
	return hash1 + hash2
}

func generateUniqID() string {
	// Get current timestamp in microseconds
	now := time.Now().UnixNano() / 1000

	// Generate 8 random bytes
	b := make([]byte, 8)
	rand.Read(b)

	// Combine timestamp and random bytes
	return fmt.Sprintf("%d%x", now, b)
}

func md5Hash(input string) string {
	hash := md5.New()
	hash.Write([]byte(input))
	return hex.EncodeToString(hash.Sum(nil))
}

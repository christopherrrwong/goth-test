package database

import (
	"database/sql"
	"fmt"
	"os"

	"sso-auth/internal/utils"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func NewConnection() error {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = os.Getenv("DBHOST")
	cfg.DBName = os.Getenv("DBNAME")

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

func InsertSSOIntegrationMapping(ssoUsername string, uuid string) error {
	token := utils.GenerateToken()
	findUser, err := DB.Query("SELECT * FROM sso_integration_mapping WHERE ssousername = ?", ssoUsername)
	if err != nil {
		return fmt.Errorf("insertSSOIntegrationMapping: %v", err)
	}
	if findUser.Next() {
		_, err := DB.Query("INSERT INTO api_token (`username`, `token`) VALUES (?, ?)", ssoUsername, token)
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

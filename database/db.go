package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	host     = os.Getenv("PG_HOST")
	user     = os.Getenv("PG_USER")
	password = os.Getenv("PG_PASSWORD")
	port     = os.Getenv("PG_PORT")
	dbname   = os.Getenv("PG_DB_NAME")
	dialect  = "postgres"
)

var (
	db  *sql.DB
	err error
)

func handleDatabaseConnection() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname,
	)

	db, err = sql.Open(dialect, psqlInfo)

	if err != nil {
		log.Panic("error occured while trying to validate database arguments:", err)
	}

	err = db.Ping()

	if err != nil {
		log.Panic("error occured while trying to connect to database:", err)
	}

}

func handleCreateRequiredTables() {
	userTable := `
		CREATE TABLE IF NOT EXISTS "users" (
			id SERIAL PRIMARY KEY,
			username VARCHAR(255) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password TEXT NOT NULL,
			age int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now()
		);
	`

	photoTable := `
		CREATE TABLE IF NOT EXISTS "photo" (
			id SERIAL PRIMARY KEY,
			title VARCHAR(255) NOT NULL,
			caption text,
			photo_url VARCHAR(255) NOT NULL,
			user_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT photos_user_id_fk 
				FOREIGN KEY (user_id)
					REFERENCES users (id)
						ON DELETE CASCADE
		);
	`

	commentTable := `
		CREATE TABLE IF NOT EXISTS "comment" (
			id SERIAL PRIMARY KEY,
			user_id int NOT NULL,
			photo_id int NOT NULL,
			message TEXT NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT comments_user_id_fk 
				FOREIGN KEY (user_id)
					REFERENCES users (id)
						ON DELETE CASCADE,
			CONSTRAINT comments_photo_id_fk
				FOREIGN KEY (photo_id)
					REFERENCES photo (id)
						ON DELETE CASCADE
		);
	`

	socialmediaTable := `
		CREATE TABLE IF NOT EXISTS "socialmedia" (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			social_media_url VARCHAR(255) NOT NULL,
			user_id int NOT NULL,
			created_at timestamptz DEFAULT now(),
			updated_at timestamptz DEFAULT now(),
			CONSTRAINT socialmedia_user_id_fk 
				FOREIGN KEY (user_id)
					REFERENCES users (id)
						ON DELETE CASCADE
		);
	`

	createTableQueries := fmt.Sprintf("%s %s %s %s", userTable, photoTable, socialmediaTable, commentTable)

	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Panic("error occured while trying to create required tables:", err.Error())
	}
}

func InitiliazeDatabase() {
	handleDatabaseConnection()
	handleCreateRequiredTables()
}

func GetDatabaseInstance() *sql.DB {
	return db
}

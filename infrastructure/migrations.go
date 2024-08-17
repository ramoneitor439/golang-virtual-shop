package infrastructure

import (
	"log"
	"os"
	"path"

	"mystore.com/infrastructure/data"
)

const UP_FOLDER = "./infrastructure/migrations/up"
const DOWN_FOLDER = "./infrastructure/migrations/down"

func ApplyMigrations(conn *data.PostgresqlConnection) error {
	files, err := os.ReadDir(UP_FOLDER)
	if err != nil {
		return err
	}

	if connErr := conn.Connect(); connErr != nil {
		return connErr
	}

	defer conn.Close()

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := path.Join(UP_FOLDER, file.Name())
		sqlContent, readErr := os.ReadFile(filePath)
		if readErr != nil {
			return readErr
		}

		log.Printf("Appliying migration [%s]...", file.Name())

		log.Printf("Migration code: \n%s", string(sqlContent))

		_, muteError := conn.Mute(string(sqlContent))
		if muteError != nil {
			return muteError
		}

		log.Printf("Migration [%s] applyied", filePath)
	}

	return nil
}

package ini

import "github.com/rudiath95/movie-API-with-JWT/models"

func SyncDatabases() {
	DB.AutoMigrate(
		models.User{},
		models.UserInfo{},
		models.VoucherList{},
		models.Director{},
		models.Genre{},
		models.Movie{},
	)
}

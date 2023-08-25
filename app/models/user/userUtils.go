package user

import "dyStudyDemo/app/database"

func IsUserExist(username string) bool {
	var count int64
	database.DB.Model(&User{}).Where("username = ?", username).Count(&count)
	return count > 0
}

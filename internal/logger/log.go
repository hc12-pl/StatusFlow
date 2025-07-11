package logger

import (
	"statusFlow/internal/database"

)

func SaveLog(text string) {
	database.InsertLog(text)
	
}

func GetAllLogs() (map[string]string, error) {
	return database.GetAllLogs()
}
package logger

import (
	"statusFlow/internal/database"

)

func SaveLog(text string) {
	database.InsertLog(text)
	
}
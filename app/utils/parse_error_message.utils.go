package utils

import (
	"gin-gorm/app/model"
	"strings"
)

func ParseErrorMessages(errorString string) []model.ApiError {
	// Split the string by newline to separate each error message
	errorMessages := strings.Split(errorString, "\n")

	var parsedErrors []model.ApiError

	for _, errorMessage := range errorMessages {
		if errorMessage == "" {
			continue
		}

		// Split each error message by " Error:" to separate the key and the error description
		parts := strings.Split(errorMessage, " Error:")
		if len(parts) != 2 {
			continue
		}

		// Extract the field name from the key part
		keyPart := strings.TrimSpace(parts[0])
		keyParts := strings.Split(keyPart, "'")
		if len(keyParts) < 2 {
			continue
		}
		fieldName := keyParts[1]

		// Clean up the error description part
		description := strings.TrimSpace(parts[1])

		// Add the field name and description to the map

		parsedErrors = append(parsedErrors, model.ApiError{Field: fieldName, Msg: description})
	}

	return parsedErrors
}

package services

func GetJSONMessage() (map[string]string, error) {
	response := map[string]string{"message": "This is a JSON response"}
	return response, nil
}

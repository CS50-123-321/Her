package api

func GetAiResponse() (string, error) {
	// get random promot
	q := GetRandomQuestion()
	aiRes, err := GenerateText(q)
	return aiRes, err
}

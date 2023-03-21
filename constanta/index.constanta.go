package constanta

const (
	ADMIN       string = "admin"
	USER        string = "user"
	TIME_FORMAT string = "2006-01-02 13:4:4"
)

func IMAGE_VALIDTION() []string {
	return []string{"image/jpg", "image/jpeg", "image/png"}
}

func FILE_VALIDTION() []string {
	return []string{"text/plain", "text/html", "application/octet-stream"}
}

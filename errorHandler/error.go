package errorHandler

type Error struct {
	Message    string `json:"message"`
	Code       int    `json:"code"`
	StatusText string `json:"status_text"`
}

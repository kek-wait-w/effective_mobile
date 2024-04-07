package domain

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Details string `json:"details"`
}

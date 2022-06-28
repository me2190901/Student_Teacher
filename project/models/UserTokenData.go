package models

type UserTokenData struct {
	Email string `json:"email" binding:"required"`
	Verify bool	`json:"verified_email" binding:"required"`
}
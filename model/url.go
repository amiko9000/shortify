package model

type URL struct {
	Model
	Token 		string `gorm:"uniqueIndex;size" json:"token"`
	RedirectURL	string `gorm:"index" json:"redirect_url"`
}
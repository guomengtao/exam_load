package models

// Member defines the model for the member table
type Member struct {
}

func (Member) TableName() string {
	return "member"
}

package models

type UserData struct {
    UserID    string `json:"user_id"`
    UserName  string `json:"user_name"`
    Balance   float64 `json:"balance"`
    // add other relevant fields
}
package entity

import "time"

type UserID = string

type User struct {
	UserID    UserID
	Login     string
	Password  string
	createdAt time.Time
	updatedAt time.Time
}

type SignUpRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LogInRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type LogInResponse struct {
	Success bool `json:"success,omitempty"`
}

type RecordInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}
type GetAllRecordsRequest struct {
}

type GetAllRecordsResponse struct {
	Records []RecordInfo `json:"records"`
}

type GetRecordsByTypeRequest struct {
	Type string `json:"type"`
}

type GetRecordsByTypeResponse struct {
	Records []RecordInfo `json:"records"`
}

type CreateRecordRequest struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type CreateRecordResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type GetRecordRequest struct {
	Id string `json:"id"`
}

type Record struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Data      []byte `json:"data"`
	userID    string
	createdAt time.Time
	updatedAt time.Time
}
type GetRecordResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Data      []byte `json:"data"`
	createdAt time.Time
	updatedAt time.Time
}

type EditRecordRequest struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	Data      []byte `json:"data"`
	createdAt time.Time
	updatedAt time.Time
}

type EditRecordResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Data []byte `json:"data"`
}

type DeleteRecordRequest struct {
	Id string `json:"id"`
}

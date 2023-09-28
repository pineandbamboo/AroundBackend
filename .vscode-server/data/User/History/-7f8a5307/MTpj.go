package model

type Post struct {
    Id      string `json:"id"`
    Username    string `json:"user"`
    Message string `json:"message"`
    Url     string `json:"url"`
    Type    string `json:"type"`
}
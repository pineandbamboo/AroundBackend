package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"

	"around/model"
	"around/service"

	"github.com/pborman/uuid"
)

var (
    mediaTypes = map[string]string{
        ".jpeg": "image",
        ".jpg":  "image",
        ".gif":  "image",
        ".png":  "image",
        ".mov":  "video",
        ".mp4":  "video",
        ".avi":  "video",
        ".flv":  "video",
        ".wmv":  "video",
    }
)


// 1. 
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    // Parse from body of request to get a json object.
    fmt.Println("Received one upload request")

    // 1. process request
    token := 
    claims :=
    username :=

    p := model.Post{
        Id:      uuid.New(),
        // User:    r.FormValue("user"),
        User: username.(string), // 从token里面直接拿过来，不用从前端传
        Message: r.FormValue("message"),
    }

    file, header, err := r.FormFile("media_file")
    if err != nil {
        http.Error(w, "Media file is not available", http.StatusBadRequest)
        fmt.Printf("Media file is not available %v\n", err)
        return
    }

    suffix := filepath.Ext(header.Filename)
    if t, ok := mediaTypes[suffix]; ok {
        p.Type = t
    } else {
        p.Type = "unknown"
    }

    // 2. handle request
    err = service.SavePost(&p, file)
    if err != nil {
        http.Error(w, "Failed to save post to backend", http.StatusInternalServerError)
        fmt.Printf("Failed to save post to backend %v\n", err)
        return
    }
    
    // 3. response
    fmt.Println("Post is saved successfully.")
}


func searchHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Received one request for search")
    w.Header().Set("Content-Type", "application/json")

    user := r.URL.Query().Get("user")
    keywords := r.URL.Query().Get("keywords")

    var posts []model.Post
    var err error
    if user != "" {
        posts, err = service.SearchPostsByUser(user)
    } else {
        posts, err = service.SearchPostsByKeywords(keywords)
    }

    if err != nil {
        http.Error(w, "Failed to read post from backend", http.StatusInternalServerError)
        fmt.Printf("Failed to read post from backend %v.\n", err)
        return
    }

    js, err := json.Marshal(posts)
    if err != nil {
        http.Error(w, "Failed to parse posts into JSON format", http.StatusInternalServerError)
        fmt.Printf("Failed to parse posts into JSON format %v.\n", err)
        return
    }
    w.Write(js)
}

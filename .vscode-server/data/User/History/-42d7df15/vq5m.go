package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
   
    "around/model"
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
    decoder := json.NewDecoder(r.Body)
    var p model.Post
    if err := decoder.Decode(&p); err != nil {
        panic(err)
    }

    fmt.Fprintf(w, "Post received: %s\n", p.Message)
}

// 2. 

// 3. 

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

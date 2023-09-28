package service

import (
	"mime/multipart"
    "reflect"

    "around/backend"
    "around/constants"
    "around/model"

    "github.com/olivere/elastic/v7"
)

func SearchPostsByUser(user string) ([]model.Post, error) {
    query := elastic.NewTermQuery("user", user)
    searchResult, err := backend.ESBackend.ReadFromES(query, constants.POST_INDEX)
    if err != nil {
        return nil, err
    }
    return getPostFromSearchResult(searchResult), nil
}

func SearchPostsByKeywords(keywords string) ([]model.Post, error) {
    query := elastic.NewMatchQuery("message", keywords)
    query.Operator("AND")
    if keywords == "" {
        query.ZeroTermsQuery("all")
    }
    searchResult, err := backend.ESBackend.ReadFromES(query, constants.POST_INDEX)
    if err != nil {
        return nil, err
    }
    return getPostFromSearchResult(searchResult), nil
}

func getPostFromSearchResult(searchResult *elastic.SearchResult) []model.Post {
    var ptype model.Post
    var posts []model.Post

    for _, item := range searchResult.Each(reflect.TypeOf(ptype)) {
        p := item.(model.Post)
        posts = append(posts, p)
    }
    return posts
}

func SavePost(post *model.Post, file miltipart.File) error {
	// 返回error，返回post后之后可以写展示新的post之类的
	// 1. Save to GCS
	url, err := backend.GCSBackend.SavetoGCS(file, post.Id)
	if err != nil {
		return err
	}

	// 1.5 get the URL from GCS, set post.URL
	post.url = url

	// 2 save to ES
	backend.ESBackend.SavetoES(post, constants.POST_INDEX, post.Id)
	// err = backend.ESBackend.SavetoES(post, constants.POST_INDEX, post.Id) {
	// 	if err != nil {
	// 		// 1. retry 3 times

	// 		// 2. delete file from GCS // rollback
			
	// 		// 3. program: check all files in bucket, check whether the url save to ES, if not, delete the file
	// 	}
	// }
}
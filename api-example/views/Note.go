package views

// PostRequestNote -> Request Post
type PostRequestNote struct {
	Title string `json:"title"`
	Desc  string `json:"description"`
}

// GetRequestNote -> Request Post
type GetRequestNote struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Desc  string `json:"description"`
}

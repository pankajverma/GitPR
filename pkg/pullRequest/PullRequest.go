package pullRequest

type AllPr struct {
	Items []Items
}

type Items struct {
	Title       string
	Url         string
	CreatedAt   string      `json:"created_at"`
	PullRequest PullRequest `json:"pull_request"`
}

type PullRequest struct {
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

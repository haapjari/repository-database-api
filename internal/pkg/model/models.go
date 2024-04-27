package model

type Repository struct {
	ID                     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name                   string `json:"name"`
	FullName               string `json:"full_name"`
	CreatedAt              string `json:"created_at"`
	StargazerCount         int    `json:"stargazer_count"`
	Language               string `json:"language"`
	OpenIssues             int    `json:"open_issues"`
	ClosedIssues           int    `json:"closed_issues"`
	OpenPullRequestCount   int    `json:"open_pull_request_count"`
	ClosedPullRequestCount int    `json:"closed_pull_request_count"`
	Forks                  int    `json:"forks"`
	WatcherCount           int    `json:"watcher_count"`
	SubscriberCount        int    `json:"subscriber_count"`
	CommitCount            int    `json:"commit_count"`
	NetworkCount           int    `json:"network_count"`
	LatestRelease          string `json:"latest_release"`
	TotalReleasesCount     int    `json:"total_releases_count"`
	ContributorCount       int    `json:"contributor_count"`
	ThirdPartyLOC          int    `json:"third_party_loc"`
	SelfWrittenLOC         int    `json:"self_written_loc"`
}

type NormalizedRepository struct {
	ID                     uint    `gorm:"primaryKey" json:"id"`
	Name                   string  `json:"name"`
	FullName               string  `json:"full_name"`
	Language               string  `json:"language"`
	LatestRelease          float64 `json:"latest_release"`
	CreatedAt              float64 `json:"created_at"`
	StargazerCount         float64 `json:"stargazer_count"`
	OpenIssues             float64 `json:"open_issues"`
	ClosedIssues           float64 `json:"closed_issues"`
	OpenPullRequestCount   float64 `json:"open_pull_request_count"`
	ClosedPullRequestCount float64 `json:"closed_pull_request_count"`
	Forks                  float64 `json:"forks"`
	WatcherCount           float64 `json:"watcher_count"`
	SubscriberCount        float64 `json:"subscriber_count"`
	CommitCount            float64 `json:"commit_count"`
	NetworkCount           float64 `json:"network_count"`
	TotalReleasesCount     float64 `json:"total_releases_count"`
	ContributorCount       float64 `json:"contributor_count"`
	ThirdPartyLOC          float64 `json:"third_party_loc"`
	SelfWrittenLOC         float64 `json:"self_written_loc"`
	Popularity             float64 `json:"popularity"`
	Activity               float64 `json:"activity"`
	Maturity               float64 `json:"maturity"`
}

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
	CommitCount            int    `json:"commit_count"`
	NetworkCount           int    `json:"network_count"`
	LatestRelease          string `json:"latest_release"`
	TotalReleasesCount     int    `json:"total_releases_count"`
	ContributorCount       int    `json:"contributor_count"`
	ThirdPartyLOC          int    `json:"third_party_loc"`
	SelfWrittenLOC         int    `json:"self_written_loc"`
}

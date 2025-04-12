package env

type Build struct {
	Repository    string `config:"GITHUB_REPOSITORY" default:""`
	RepositoryUrl string `config:"GITHUB_REPOSITORY_URL" default:""`
	Commit        string `config:"GITHUB_SHA" default:""`

	AppVersion string `config:"APP_VERSION" default:""`
}

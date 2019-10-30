package command

var (
	// version used for log and info
	appVersion = "unknown"
	// sha1 from git, output of $(git rev-parse HEAD)
	gitCommit = "$Format:%H$"
	// build date in ISO8601 format, output of $(date -u +'%Y-%m-%dT%H:%M:%SZ')
	buildDate = "1970-01-01T00:00:00Z"
)

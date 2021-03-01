package build

import "strings"

const (
	// ShortVersion 短版本号
	ShortVersion = "1.2.0"
)

// The value of variables come form `gb build -ldflags '-X "build.Build=xxxxx" -X "build.CommitID=xxxx"' `
var (
	// Build build time
	Build string
	// Branch current git branch
	Branch string
	// Commit git commit id
	Commit string
)

// Version 生成版本信息
func Version() string {
	var buf strings.Builder
	buf.WriteString(ShortVersion)
	buf.WriteString(" Modified @2021-02-28 17:05:38 by Zachary")

	if Build != "" {
		buf.WriteByte('\n')
		buf.WriteString("build: ")
		buf.WriteString(Build)
	}
	if Branch != "" {
		buf.WriteByte('\n')
		buf.WriteString("branch: ")
		buf.WriteString(Branch)
	}
	if Commit != "" {
		buf.WriteByte('\n')
		buf.WriteString("commit: ")
		buf.WriteString(Commit)
	}
	return buf.String()
}

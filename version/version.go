package version

import (
	"fmt"
	"runtime"
)

// Version - version
const Version = "0.0.1"

var (
	// NetworkType is set by build flags
	NetworkType = ""

	// BuildTags is set by build flags
	BuildTags = ""
)

type VersionInfo struct {
	Name        string `json:"name" yaml:"name"`
	ServerName  string `json:"server_name" yaml:"server_name"`
	ClientName  string `json:"client_name" yaml:"client_name"`
	Version     string `json:"version" yaml:"version"`
	NetworkType string `json:"network_type" yaml:"network_type"`
	// GitCommit  string `json:"commit" yaml:"commit"`
	BuildTags string `json:"build_tags" yaml:"build_tags"`
	GoVersion string `json:"go_version" yaml:"go_version"`
}

func NewVersionInfo() VersionInfo {
	return VersionInfo{
		Name:        "mmx",
		ServerName:  "mmxd",
		ClientName:  "mmxcli",
		Version:     Version,
		NetworkType: NetworkType,
		BuildTags:   BuildTags,
		GoVersion:   fmt.Sprintf("%s %s/%s", runtime.Version(), runtime.GOOS, runtime.GOARCH),
	}
}

func (vi VersionInfo) String() string {
	return fmt.Sprintf(`%s: %s
build tags: %s
%s`,
		vi.Name, vi.Version, vi.BuildTags, vi.GoVersion,
	)
}

// Copyright 2021 switch-st[switch.st@gmail.com]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Original source: github.com/switch-st/gosu/version/version.go

package version

import (
	"strconv"

	ver "github.com/hashicorp/go-version"
)

var (
	// versionName set by build flags, like:
	// go build -ldflags "-X github.com/switch-st/gosu/version.versionName=$(GOSU_VERSION_NAME)"
	// GOSU_VERSION_NAME=$(shell git describe --tags --always --dirty)
	versionName string

	// versionCode set by build flags, like:
	// go build -ldflags "-X github.com/switch-st/gosu/version.versionCode=$(GOSU_VERSION_CODE)"
	// GOSU_VERSION_CODE=$(shell git rev-list --all --count)
	versionCode string

	// commitID set by build flags, like:
	// go build -ldflags "-X github.com/switch-st/gosu/version.commitID=$(GOSU_COMMIT_ID)"
	// GOSU_COMMIT_ID=$(shell git describe --always --dirty --abbrev=0)
	commitID string

	// buildDate set by build flags, like:
	// go build -ldflags "-X github.com/switch-st/gosu/version.buildDate=$(GOSU_BUILD_DATE)"
	// GOSU_BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')
	buildDate string
)

// GetVersionName return SemVer specs version name
func GetVersionName() string {
	if versionName == "" {
		return ""
	}
	return "v" + ver.Must(ver.NewSemver(versionName)).String()
}

// GetVersionCode return version code
func GetVersionCode() int {
	if versionCode == "" {
		return 0
	}
	code, err := strconv.Atoi(versionCode)
	if err != nil {
		panic(err)
	}
	return code
}

// GetCommitID return git commit id
func GetCommitID() string {
	return commitID
}

// GetBuildDate return build date
func GetBuildDate() string {
	return buildDate
}

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
// Original source: github.com/switch-st/gosu/version/version_test.go

package version

import (
	"testing"
	"time"
)

const (
	buildFormat = "2006-01-02T15:04:05-0700"
)

func TestVersion(t *testing.T) {
	versionName = "v1.2.3"
	versionNameExpect := "v1.2.3"
	versionCode = "123"
	versionCodeExpect := 123
	buildDate = time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local).Format(buildFormat)
	buildDateExpect := time.Date(2006, 01, 02, 15, 04, 05, 0, time.Local).Format(buildFormat)

	if Name() != versionNameExpect {
		t.Fatalf("VersionName: expect: %s, actual: %s\n", versionNameExpect, Name())
	}
	if Code() != versionCodeExpect {
		t.Fatalf("VersionCode: expect: %d, actual: %d\n", versionCodeExpect, Code())
	}
	if BuildDate() != buildDateExpect {
		t.Fatalf("BuildDate: expect: %s, actual: %s\n", buildDateExpect, BuildDate())
	}
}

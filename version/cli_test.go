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
// Original source: github.com/switch-st/gosu/version/cli_test.go

package version

import (
	"fmt"
	"testing"

	"github.com/urfave/cli/v2"
)

func TestWithCliVersion(t *testing.T) {
	versionName = "v1.2.3"
	versionNameExpect := "version = v1.2.3"
	versionNameResult := ""
	cmd := []string{"main", "-v"}

	printer := func(c *cli.Context) {
		versionNameResult = fmt.Sprintf("version = %s", c.App.Version)
	}
	cli.VersionPrinter = printer

	app := cli.NewApp()
	WithCliVersion(app)
	_ = app.Run(cmd)

	if versionNameResult != versionNameExpect {
		t.Fatalf("%s: expect: %s, actual: %s\n", t.Name(), versionNameExpect, versionNameResult)
	}
}

func TestWithCliVersionNil(t *testing.T) {
	if WithCliVersion(nil) != nil {
		t.Fatalf("%s: expect: nil, actual: not nil\n", t.Name())
	}
}

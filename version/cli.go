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
// Original source: github.com/switch-st/gosu/version/cli.go

package version

import (
	"fmt"
	"strings"

	"github.com/urfave/cli/v2"
)

var (
	DefaultPrintFunc = cliVersionPrinter
)

func cliVersionPrinter(c *cli.Context) {
	n := len(c.App.Name)
	head := strings.Repeat(" ", n)
	_, _ = fmt.Fprintf(c.App.Writer, "%v version %v-r%v\n", c.App.Name, GetVersionName(), GetVersionCode())
	_, _ = fmt.Fprintf(c.App.Writer, "%v commit with %v\n", head, GetCommitID())
	_, _ = fmt.Fprintf(c.App.Writer, "%v build on %v\n", head, GetBuildDate())
}

// WithCliVersion set version name to *cli.App instance
func WithCliVersion(c *cli.App) *cli.App {
	if c == nil {
		return nil
	}
	c.HideVersion = false
	c.Version = GetVersionName()
	cli.VersionPrinter = DefaultPrintFunc
	return c
}

// WithCliVersion set version name to *cli.App instance
func (v *Version) WithCliVersion(c *cli.App) *cli.App {
	if v == nil || c == nil {
		return c
	}
	c.HideVersion = false
	c.Version = v.GetVersionName()
	return c
}

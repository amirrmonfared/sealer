// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

type RunFlags struct {
	Masters string
	Nodes   string

	User        string
	Password    string
	Port        uint16
	Pk          string
	PkPassword  string
	CustomEnv   []string
	Mode        string
	ClusterFile string

	// override default Cmds of sealer image
	Cmds []string
	// override default APPNames of sealer image
	// Only one can be selected for LaunchCmds and AppNames
	AppNames []string
}

type ApplyFlags struct {
	Masters string
	Nodes   string

	User       string
	Password   string
	Port       uint16
	Pk         string
	PkPassword string

	ClusterFile string
	Mode        string
	CustomEnv   []string
	ForceDelete bool
}

type ScaleUpFlags struct {
	Masters string
	Nodes   string

	User       string
	Password   string
	Port       uint16
	Pk         string
	PkPassword string
	CustomEnv  []string
}

type DeleteFlags struct {
	Masters     string
	Nodes       string
	CustomEnv   []string
	ClusterFile string
	DeleteAll   bool
	ForceDelete bool
}

type MergeFlags struct {
	Masters string
	Nodes   string

	User       string
	Password   string
	Port       uint16
	Pk         string
	PkPassword string
	CustomEnv  []string

	// override default Cmds of sealer image
	Cmds []string
	// override default APPNames of sealer image
	AppNames []string
}

type UpgradeFlags struct {
	ClusterFile string
}

type RollbackFlags struct {
	AppNames []string // override default APPNames of sealer image
}

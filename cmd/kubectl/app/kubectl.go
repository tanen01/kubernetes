/*
Copyright 2015 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package app

import (
	"os"

	"github.com/spf13/pflag"
	"k8s.io/kubernetes/pkg/kubectl/cmd"
	cmdutil "k8s.io/kubernetes/pkg/kubectl/cmd/util"
)

func Run() error {
	cmd := cmd.NewKubectlCommand(cmdutil.NewFactory(nil), os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}

func AddFlags(fs *pflag.FlagSet, kubectlSubCmd string) {
	cmd := cmd.NewKubectlCommand(cmdutil.NewFactory(nil), os.Stdin, os.Stdout, os.Stderr)
	fs.AddFlagSet(cmd.LocalFlags())
	if kubectlSubCmd == "" {
		return
	}
	for _, child := range cmd.Commands() {
		if child.Name() == kubectlSubCmd {
			fs.AddFlagSet(child.LocalFlags())
			return
		}
	}
}

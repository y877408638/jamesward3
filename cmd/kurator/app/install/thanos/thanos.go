/*
Copyright Kurator Authors.

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

package thanos

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"kurator.dev/kurator/pkg/generic"
	plugin "kurator.dev/kurator/pkg/plugin/thanos"
)

var pluginArgs = plugin.InstallArgs{}

func NewCmd(opts *generic.Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "thanos",
		Short: "Install thanos component",
		RunE: func(c *cobra.Command, args []string) error {
			plugin, err := plugin.NewPlugin(opts, &pluginArgs)
			if err != nil {
				logrus.Errorf("thanos plugin init error: %v", err)
				return err
			}

			if err := plugin.Execute(args, nil); err != nil {
				logrus.Errorf("thanos execute error: %v", err)
				return err
			}

			logrus.Infof("thanos install completed")
			return nil
		},
	}

	f := cmd.PersistentFlags()
	f.StringVar(&pluginArgs.HostKubeconfig, "host-kubeconfig", "/etc/karmada/karmada-host.config",
		"Karmada host kubeconfig, default to /etc/karmada/karmada-apiserver.config")
	f.StringVar(&pluginArgs.HostContext, "host-context", "karmada", "name of karmada context")

	f.StringVar(&pluginArgs.ObjectStoreConfig, "object-store-config", "", "Object store configuration used by thanos, read more details here: https://prometheus-operator.dev/docs/operator/thanos/#configuring-thanos-object-storage")

	return cmd
}

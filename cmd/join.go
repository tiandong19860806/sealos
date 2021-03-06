// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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

package cmd

import (
	"github.com/fanux/sealos/install"

	"github.com/spf13/cobra"
)

// joinCmd represents the join command
var joinCmd = &cobra.Command{
	Use:   "join",
	Short: "Simplest way to join your kubernets HA cluster",
	Long:  `sealos join --node 192.168.0.5`,
	Run: func(cmd *cobra.Command, args []string) {
		beforeNodes:=install.ParseIPs(install.NodeIPs)
		beforeMasters:=install.ParseIPs(install.MasterIPs)

		c := &install.SealConfig{}
		c.Load("")
		install.BuildJoin(beforeMasters,beforeNodes)

		install.NodeIPs = append(c.Nodes,beforeNodes...)
		install.MasterIPs = append(c.Masters,beforeMasters...)
		c.Dump("")
	},
}

func init() {
	rootCmd.AddCommand(joinCmd)
	joinCmd.Flags().StringSliceVar(&install.MasterIPs, "master", []string{}, "kubernetes multi-master ex. 192.168.0.5-192.168.0.5")
	joinCmd.Flags().StringSliceVar(&install.NodeIPs, "node", []string{}, "kubernetes multi-nodes ex. 192.168.0.5-192.168.0.5")
}

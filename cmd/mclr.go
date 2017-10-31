// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
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
	"fmt"
	"os"
	"github.com/spf13/cobra"
<<<<<<< HEAD
	"Agenda-GO/entity/meeting"
	"Agenda-GO/user"
=======
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
)

// mclrCmd represents the mclr command
var mclrCmd = &cobra.Command{
	Use:   "mclr",
	Short: "to clear all meetings that you sponsor",
	Long: `to clear all meetings that you sponsor(Login first)
	For example:

./app mclr`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mclr called")
<<<<<<< HEAD
		if !user.IsLogin() {
			fmt.Println("please login first!")
			os.Exit(1)
		}
		meeting.ClearAllMeeting()
=======
		if !IsLogin() {
			fmt.Println("please login first!")
			os.Exit(1)
		}
		if err := ClearAllMeeting(); err != nil {
			
			os.Exit(2)
		}
>>>>>>> 015ad6e43e4e87e3f14180abcf8579d3401169b1
	},
}

func init() {
	RootCmd.AddCommand(mclrCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mclrCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mclrCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

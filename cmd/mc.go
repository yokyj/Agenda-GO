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
)

// mcCmd represents the mc command
var mcCmd = &cobra.Command{
	Use:   "mc",
	Short: "to create a new meeting",
	Long: `to create a new meeting with title, participator,starttime and endtime.
	 For example:

./app mc -ttest -pPeter -pMarry -s"2017-10-28 09:30" -e"2017-10-28 10:30"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mc called")
		title, _ := cmd.Flags().GetString("title")
		participators, _ := cmd.Flags().GetStringArray("parti")
		stime, _ := cmd.Flags().GetString("stime")
		etime, _ := cmd.Flags().GetString("etime")
		if  !IsLogin() {
			fmt.Println("Login first!")
			os.Exit(1)
		}
		if title == "" {
			fmt.Println("title can not be blank")
			os.Exit(2)
		}
		if stime == "" {
			fmt.Println("stime can not be blank")
			os.Exit(3)
		}
		if etime == "" {
			fmt.Println("etime can not be blank")
			os.Exit(4)
		}
		if err := CreateMeeting(title, participator, stime, etime); err != nil {
			
			os.Exit(5)
		}
		//MeetingCreate(title, participators, stime, etime)
		fmt.Println(title + " created")
		
	},
}

func init() {
	RootCmd.AddCommand(mcCmd)
	var strarr []string
	

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mcCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mcCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	mcCmd.Flags().StringP("title", "t", "", "title of the meeting which should be unique")
	mcCmd.Flags().StringArrayP("parti", "p", strarr, "participators of the meeting ")
	mcCmd.Flags().StringP("stime", "s", "", "time when the meeting will begin")
	mcCmd.Flags().StringP("etime", "e", "", "time when the meeting will end")
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/yigecaiji/AgendaGO/entity"

	"os"

	"log"
)

// findUserCmd represents the findUser command
var findUserCmd = &cobra.Command{
	Use:   "findUser",
	Short: "find all users",
	Long:  `return all messages of all users`,
	Run: func(cmd *cobra.Command, args []string) {
		logFile, _ := os.OpenFile("logger.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
		logger := log.New(logFile, "logger: ", log.Ldate|log.Ltime)
		username, _ := cmd.Flags().GetString("userName")
		logger.Print("AgendaGO findUser ")
		var user entity.User
		user.Init(username, "", "", "")
		userList, ok := user.FindAllUsers()
		if ok {
			logger.Print("-- succeed")
			for _, user := range userList {
				user.PrintUser()
			}
		} else {
			logger.Print("-- failed: failed to find users. Maybe you don't log in yet")
			fmt.Println("failed to find users. Maybe you don't log in yet")
		}
		logFile.Close()
	},
}

func init() {
	rootCmd.AddCommand(findUserCmd)
	findUserCmd.Flags().StringP("userName", "u", "Anonymous", "username of user")
package cmd

import (
    "os"
	"fmt"
    "agenda/entity"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long: "agenda login -u [username] -p [password]",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("login called")
        usern,_ := cmd.Flags().GetString("username")
        pass,_ := cmd.Flags().GetString("password")
        state := entity.Login(usern,pass)
        if state == 0{
            fmt.Println("The user doesn't exit")
        }else if(state == 1){
            fmt.Fprintf(os.Stdout,"%s login succeed\n",usern)
           // entity.saveFile()
        }else if(state == 2){
            fmt.Println("the username can't match with the password")
        }
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
    entity.ReadFile()
    loginCmd.Flags().StringP("username","u","","")
    loginCmd.Flags().StringP("password","p","","")
}
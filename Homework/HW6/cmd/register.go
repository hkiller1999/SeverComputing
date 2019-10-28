package cmd

import (
	"fmt"
    "agenda/entity"
	"github.com/spf13/cobra"
)

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "resister",
	Long: "agenda registr -u [username] -p [password] -e [email] -h [phone]",
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("register called")
        usern,_ := cmd.Flags().GetString("username")
        pass,_ := cmd.Flags().GetString("password")
        email,_ := cmd.Flags().GetString("email")
        phone,_ := cmd.Flags().GetString("phone")
        state := entity.Register(usern,pass,email,phone)
        if state == 0{
            fmt.Println("A same username have been used")
        }else if(state == 1){
            fmt.Println("register succeed")
           // entity.saveFile()
        }else if(state == 2){
            fmt.Println("agenda registr -u [username] -p [password] -e [email] -h [phone]")
        }
	},
}

func init() {
	rootCmd.AddCommand(registerCmd)
    entity.ReadFile()
    registerCmd.Flags().StringP("username","u","","")
    registerCmd.Flags().StringP("password","p","","")
    registerCmd.Flags().StringP("email","e","","")
    registerCmd.Flags().StringP("phone","t","","")
}
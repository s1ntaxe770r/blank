package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/s1ntaxe770r/blank/blank/utils"
	"github.com/spf13/cobra"
)

const URL_ERROR = `
Could not find server to upload file to. Consider setting exporting your blink server address to an environment variable

EG: export BLINK_SERVER = https://yourblinkserveraddress.file
`

func init() {
	rootCmd.AddCommand(uploadcmd)
}

// check if blink instance requires auth
func GetCreds() (bool, utils.Credentials) {
	uname := os.Getenv("BLINK_ADMIN")
	pass := os.Getenv("BLINK_PASS")
	if uname == "" && pass == "" {
		return false, utils.Credentials{}
	}
	credentials := &utils.Credentials{Username: uname, Password: pass}
	return true, *credentials
}

var uploadcmd = &cobra.Command{
	Use:   "upload [ File ]",
	Short: "upload a file ",
	Long:  "Upload a file to your blink instance ",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		s := spinner.New(spinner.CharSets[36], 100*time.Millisecond)
		blink_server := os.Getenv("BLINK_SERVER")
		if blink_server == "" {
			fmt.Println(color.RedString(URL_ERROR))
			os.Exit(1)
		}
		file := args[0]
		// if err != nil {
		// 	fmt.Printf("could not open file %v. Reason %s", file.Name(), err.Error())
		// 	os.Exit(1)
		// }
		_, credentials := GetCreds()
		fmt.Println(color.YellowString("attempting to upload file"))
		s.Start()
		uploaderr := utils.UploadFile(file, blink_server, credentials)
		if uploaderr != nil {
			fmt.Println(uploaderr.Error())
			s.Stop()
			os.Exit(1)
		}
		s.Stop()
		fmt.Println(color.HiGreenString("upload successfull. File url has been copied to your clipboard"))
		os.Exit(0)
	},
}

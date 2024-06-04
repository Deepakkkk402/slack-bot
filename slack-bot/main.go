package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	_ "github.com/slack-go/slack"
)

func main() {

	os.Setenv("SLACK_BOT", "xoxb-7041482616663-7185815040148-fwAznN1Wiu4C2aK47XJNpsm3")
	os.Setenv("CHANNEL_ID", "C071K3U5EUV")
	api := slack.New(os.Getenv("SLACK_BOT"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Deepak_Nair_.docx"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: channelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("Name is %s , URL:%s\n", file.Name, file.URL)
	}

}

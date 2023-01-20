package main

import (
	"context"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viper.ReadinConfig()
	apiKey := viper.GetString("API_KEY")
	if apiKey == "" {
		panic("API KEY Missing")
	}

	ctx := context.Background()
	client := gpt3.NewClient(apiKey)
	rootCmd := &cobra.Command{
		Use: "chatgpt",
		Short: "Chat with ChatGPT in console.",
		Run: func(cmd *cobra.Command , args []string){
			scanner := bufio.NewScanner(os.Stdin)
			quit := false

			for !quit{
				fmt.Print("Say Something('quit' to end):")
			}
		}
	}

}

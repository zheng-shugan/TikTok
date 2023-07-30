package main

import "github.com/sunflower10086/TikTok/http/cmd"

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		panic(err)
	}
}

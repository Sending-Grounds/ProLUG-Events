package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)
func main() {
	/* 
     Evaluate the .env file if one is located in the current directory for the module
     You also have the option of adding any environment variables to path 
  */
	godotenv.Load()

	discordAuthToken := os.Getenv("DISCORD_AUTH_TOKEN")

	if discordAuthToken == "" {
		fmt.Println("ERROR: Could not located the auth token for the discord bot. Please ensure DISCORD_AUTH_TOKEN is set in the .env or your current shell environment")
		os.Exit(1)
	}

	/* Initialize the discord session */
	dg, err := discordgo.New("Bot " + discordAuthToken)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	/* For now we are only looking at messages. If the bot intent changes. Let KingBunz know so he can update on the developer portal. */
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	dg.AddHandler(messageCreate)


	/* Open websocket connection */
	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening connection: ", err)
		return
	}

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	/* Ignore messages created by the bot */
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}

	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
}

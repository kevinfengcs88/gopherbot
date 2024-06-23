package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"gopherbot/gophercommands"
	"gopherbot/gopherutils"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()
	botToken := fmt.Sprintf("Bot %s", os.Getenv("TOKEN"))
	dg, err := discordgo.New(botToken)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(handleMessage)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	os.Stdout.Sync()
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself
	if m.Author.ID == s.State.User.ID && !strings.HasPrefix(m.Content, "!") {
		return
	}

	cmd := gopherutils.RemovePrefix(m.Content)

	switch cmd {
	case "ping":
		gophercommands.Ping(s, m)
	case "help":
		gophercommands.Help(s, m)
	case "serverstatus", "statusserver", "status":
		gophercommands.ServerStatus(s, m)
	case "serverstart", "startserver", "start":
		gophercommands.ServerStart(s, m)
	case "serverstop", "stopserver", "stop", "serverkill", "killserver", "kill":
		gophercommands.ServerStop(s, m)
	}
}

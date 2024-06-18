package gophercommands

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/bwmarrin/discordgo"
)

func ServerStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Checking the status of the server...")
	cmd := exec.Command("/home/kevinfengcs88/go/src/gopherbot/shell/server.sh", "status")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh status: %v", err)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error executing server.sh status: %v", err))
	}
	s.ChannelMessageSend(m.ChannelID, string(stdout))
}

func ServerStart(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Starting up the server...")
	cmd := exec.Command("/home/kevinfengcs88/go/src/gopherbot/shell/server.sh", "start")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh start: %v", err)
	}

	fmt.Println(string(stdout))
}

func ServerStop(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Stopping the server...")
	cmd := exec.Command("/home/kevinfengcs88/go/src/gopherbot/shell/server.sh", "stop")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh stop: %v", err)
	}

	fmt.Println(string(stdout))
}

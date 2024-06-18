package gophercommands

import (
	"fmt"
	"gopherbot/gopherutils"
	"log"
	"os/exec"
	"regexp"

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
	serverStatus := gopherutils.CleanedBytesToString(stdout)
	serverDownPattern := regexp.MustCompile(`\w+:\sDOWN`)
	serverDownMatch := serverDownPattern.MatchString(serverStatus)
	if serverDownMatch {
		serverStatus = gopherutils.Redify(serverStatus)
	} else {
		serverStatus = gopherutils.Greenify(serverStatus)
	}
	s.ChannelMessageSend(m.ChannelID, serverStatus)
}

func ServerStart(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Starting up the server...")
	cmd := exec.Command("/home/kevinfengcs88/go/src/gopherbot/shell/server.sh", "start")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh start: %v", err)
	}
	s.ChannelMessageSend(m.ChannelID, string(stdout))
}

func ServerStop(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Stopping the server...")
	cmd := exec.Command("/home/kevinfengcs88/go/src/gopherbot/shell/server.sh", "stop")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh stop: %v", err)
	}
	s.ChannelMessageSend(m.ChannelID, string(stdout))
}

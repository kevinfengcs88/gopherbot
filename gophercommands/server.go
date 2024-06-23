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
	cmd := exec.Command("/home/kevinfengcs88/go/src/gopherbot/shell/server.sh", "status")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh status: %v", err)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error executing server.sh status: %v", err))
	}
	output := gopherutils.CleanedBytesToString(stdout)
	serverDownPattern := regexp.MustCompile(`(?i)down`)
	serverDownMatch := serverDownPattern.MatchString(output)
	if serverDownMatch {
		output = gopherutils.Redify(output)
	} else {
		output = gopherutils.Greenify(output)
	}
	s.ChannelMessageSend(m.ChannelID, output)
}

func ServerStart(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Attempting to start the server...")
	cmd := exec.Command("/home/kevinfengcs88/go/src/gopherbot/shell/server.sh", "start")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh start: %v", err)
	}
	output := gopherutils.CleanedBytesToString(stdout)
	serverUpPattern := regexp.MustCompile(`(?i)up`)
	serverUpMatch := serverUpPattern.MatchString(output)
	if serverUpMatch {
		output = gopherutils.Greenify(output)
	}
	s.ChannelMessageSend(m.ChannelID, output)
}

func ServerStop(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Attempting to stop the server...")
	cmd := exec.Command("/home/kevinfengcs88/go/src/gopherbot/shell/server.sh", "stop")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh stop: %v", err)
	}
	output := gopherutils.CleanedBytesToString(stdout)
	serverDownPattern := regexp.MustCompile(`(?i)down`)
	serverDownMatch := serverDownPattern.MatchString(output)
	if serverDownMatch {
		output = gopherutils.Redify(output)
	}
	s.ChannelMessageSend(m.ChannelID, output)
}

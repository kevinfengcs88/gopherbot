package gophercommands

import (
	"fmt"
	"gopherbot/gopherutils"
	"log"
	"os/exec"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var serverScript string = "/home/kevinfengcs88/go/src/gopherbot/shell/server.sh"

func ServerStatus(s *discordgo.Session, m *discordgo.MessageCreate) {
	cmd := exec.Command(serverScript, "status")

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
	cmd := exec.Command(serverScript, "start")

	_, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh start: %v", err)
	}

	// this shee broken ong frfr no skibidi cap, cuh ohio
	ServerStatus(s, m)
}

func ServerStop(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Attempting to stop the server...")
	cmd := exec.Command(serverScript, "stop")

	stdout, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error executing server.sh stop: %v", err)
	}
	output := gopherutils.CleanedBytesToString(stdout)
	serverDownPattern := regexp.MustCompile(`(?i)down`)
	serverDownMatch := serverDownPattern.MatchString(output)
	if serverDownMatch {
		output = gopherutils.Redify(output)
	} else if strings.HasPrefix(output, "SUCCESS") {
		output = gopherutils.Greenify(output)
	}
	s.ChannelMessageSend(m.ChannelID, output)
}

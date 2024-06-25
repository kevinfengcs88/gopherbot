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

func ProbingServerStatus(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	cmd := exec.Command(serverScript, "status")

	stdout, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatalf("Error executing server.sh status: %v", err)
		s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("Error executing server.sh status: %v", err))
	}

	output := gopherutils.CleanedBytesToString(stdout)
	serverDownPattern := regexp.MustCompile(`(?i)down`)
	serverDownMatch := serverDownPattern.MatchString(output)

	var status bool

	if serverDownMatch {
		// no output needed if server is down
		status = false
		return status
	} else {
		output = gopherutils.Greenify(output)
		status = true
	}
	s.ChannelMessageSend(m.ChannelID, output)
	return status
}

func ServerStart(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Attempting to start the server...")
	cmd := exec.Command(serverScript, "start")

	go func() {
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Error executing server.sh start: %v", err)
		}
	}()

	for status := false; !status; {
		status = ProbingServerStatus(s, m)
	}
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

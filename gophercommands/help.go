package gophercommands

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	commandsFile, err := os.Open("data/commands.json")
	if err != nil {
		log.Fatalf("Failed to open the file: %s", err)
	}
	defer commandsFile.Close()

	byteValue, err := io.ReadAll(commandsFile)
	if err != nil {
		log.Fatalf("Failed to read the file: %s", err)
	}

	var commands map[string]string
	if err := json.Unmarshal(byteValue, &commands); err != nil {
		log.Fatalf("Failed to unmarshal JSON: %s", err)
	}

	var helpOutput string
	var builder strings.Builder

	for key, value := range commands {
		builder.WriteString(fmt.Sprintf("%s: %s\n", key, value))
	}

	helpOutput = builder.String()

	s.ChannelMessageSend(m.ChannelID, helpOutput)
}

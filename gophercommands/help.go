package gophercommands

import (
    "github.com/bwmarrin/discordgo"
    "encoding/json"
    "os"
    "log"
    "io"
    "fmt"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
    type CommandsMap map[string]string

    // try using interface method instead???
    // to simulate unknown type and depth of this data
    commandsFile, err := os.Open("data/commands.json")
    if err != nil {
        log.Fatalf("Failed to open the file: %s", err)
    }
    defer commandsFile.Close()

    byteValue, err := io.ReadAll(commandsFile)
    if err != nil {
        log.Fatalf("Failed to read the file: %s", err)
    }

    var commands CommandsMap
    if err := json.Unmarshal(byteValue, &commands); err != nil {
        log.Fatalf("Failed to unmarshal JSON: %s", err)
    }

    for key, value := range commands {
        fmt.Printf("%s: %s\n", key, value)
    }

    s.ChannelMessageSend(m.ChannelID, "Help contents printed to terminal, sir")
}

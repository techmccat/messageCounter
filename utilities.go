package main

import (
	"bytes"
	"github.com/bwmarrin/discordgo"
	"github.com/bwmarrin/lit"
	"github.com/goccy/go-json"
	"sort"
)

// Sends embed as response to an interaction
func sendEmbedInteraction(s *discordgo.Session, embed *discordgo.MessageEmbed, i *discordgo.Interaction) {
	sliceEmbed := []*discordgo.MessageEmbed{embed}
	err := s.InteractionRespond(i, &discordgo.InteractionResponse{Type: discordgo.InteractionResponseChannelMessageWithSource, Data: &discordgo.InteractionResponseData{Embeds: sliceEmbed}})
	if err != nil {
		lit.Error("InteractionRespond failed: %s", err)
	}
}

func sendEmbedInteractionFollowup(s *discordgo.Session, embed *discordgo.MessageEmbed, i *discordgo.Interaction) {
	sliceEmbed := []*discordgo.MessageEmbed{embed}
	_, err := s.FollowupMessageCreate(s.State.User.ID, i, false, &discordgo.WebhookParams{Embeds: sliceEmbed})
	if err != nil {
		lit.Error("FollowupMessageCreate failed: %s", err)
	}
}

// Sorts a map into an array
func sorting(classifica map[string]int) []kv {
	var ss []kv
	for k, v := range classifica {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}

// isCommandEqual compares two command by marshalling them to JSON. Yes, I know. I don't want to write recursive things.
func isCommandEqual(c *discordgo.ApplicationCommand, v *discordgo.ApplicationCommand) bool {
	c.Version = ""
	c.ID = ""
	c.ApplicationID = ""
	c.Type = 0
	cBytes, _ := json.Marshal(&c)

	v.Version = ""
	v.ID = ""
	v.ApplicationID = ""
	v.Type = 0
	vBytes, _ := json.Marshal(&v)

	return bytes.Compare(cBytes, vBytes) == 0
}

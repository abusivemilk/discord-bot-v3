package bot

import (
	"github.com/VATUSA/discord-bot-v3/internal/config"
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func IntervalRefreshAll(s *discordgo.Session) {
	for {
		log.Printf("Fetching all guilds")
		ProcessAllGuilds(s)
		time.Sleep(1 * time.Hour)
	}
}

func ProcessAllGuilds(s *discordgo.Session) {
	for _, guild := range s.State.Guilds {
		cfg := config.GetServerConfig(guild.ID)
		if cfg != nil && cfg.Active {
			err := RequestGuildMembers(s, guild, cfg)
			if err != nil {
				log.Printf("Error while processing guild members for guild %s (%s): %s\n",
					guild.ID, cfg.Name, err.Error())
			}
		}
	}
}

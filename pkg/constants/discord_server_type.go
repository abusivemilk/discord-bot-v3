package constants

type DiscordServerType = uint

const (
	DivisionOfficial DiscordServerType = iota
	SubdivisionOfficial
	Affiliated
	Other
)

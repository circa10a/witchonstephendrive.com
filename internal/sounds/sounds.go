package sounds

// SupportedSounds is a slice of available mp3's in the sounds directory
var SupportedSounds = []string{
	"dracula",
	"ghost",
	"halloween-organ",
	"leave-now",
	"pumpkin-king",
	"scream",
	"this-is-halloween",
	"werewolf",
	"witch-laugh",
}

// PlaySoundPayload is the type supported by assistant-relay to cast custom media
type PlaySoundPayload struct {
	Device string `json:"device"`
	Source string `json:"source"`
	Type   string `json:"type"`
}

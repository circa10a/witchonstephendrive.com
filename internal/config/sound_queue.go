package config

import (
	"github.com/oleiade/lane"
)

// InitSoundsQueue creates the capped sound queue
func (w *WitchConfig) InitSoundQueue() {
	// Create new queue to process "sound" jobs one at a time with a max limit to eliminate spam
	w.SoundQueue = lane.NewCappedDeque(w.SoundQueueCapacity)
}

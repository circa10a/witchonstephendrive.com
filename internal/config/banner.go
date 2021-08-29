package config

import (
	"math/rand"

	"github.com/fatih/color"
)

//nolint:lll // long lines gonna be long
const banner = `
@@@  @@@   @@@@@@   @@@@@@@   @@@@@@@   @@@ @@@     @@@  @@@   @@@@@@   @@@       @@@        @@@@@@   @@@  @@@  @@@  @@@@@@@@  @@@@@@@@  @@@  @@@
@@@  @@@  @@@@@@@@  @@@@@@@@  @@@@@@@@  @@@ @@@     @@@  @@@  @@@@@@@@  @@@       @@@       @@@@@@@@  @@@  @@@  @@@  @@@@@@@@  @@@@@@@@  @@@@ @@@
@@!  @@@  @@!  @@@  @@!  @@@  @@!  @@@  @@! !@@     @@!  @@@  @@!  @@@  @@!       @@!       @@!  @@@  @@!  @@!  @@!  @@!       @@!       @@!@!@@@
!@!  @!@  !@!  @!@  !@!  @!@  !@!  @!@  !@! @!!     !@!  @!@  !@!  @!@  !@!       !@!       !@!  @!@  !@!  !@!  !@!  !@!       !@!       !@!!@!@!
@!@!@!@!  @!@!@!@!  @!@@!@!   @!@@!@!    !@!@!      @!@!@!@!  @!@!@!@!  @!!       @!!       @!@  !@!  @!!  !!@  @!@  @!!!:!    @!!!:!    @!@ !!@!
!!!@!!!!  !!!@!!!!  !!@!!!    !!@!!!      @!!!      !!!@!!!!  !!!@!!!!  !!!       !!!       !@!  !!!  !@!  !!!  !@!  !!!!!:    !!!!!:    !@!  !!!
!!:  !!!  !!:  !!!  !!:       !!:         !!:       !!:  !!!  !!:  !!!  !!:       !!:       !!:  !!!  !!:  !!:  !!:  !!:       !!:       !!:  !!!
:!:  !:!  :!:  !:!  :!:       :!:         :!:       :!:  !:!  :!:  !:!   :!:       :!:      :!:  !:!  :!:  :!:  :!:  :!:       :!:       :!:  !:!
::   :::  ::   :::   ::        ::          ::       ::   :::  ::   :::   :: ::::   :: ::::  ::::: ::   :::: :: :::    :: ::::   :: ::::   ::   ::
 :   : :   :   : :   :         :           :         :   : :   :   : :  : :: : :  : :: : :   : :  :     :: :  : :     : :: ::   : :: ::   ::    :
`

// PrintBanner prints a HAPPY HALLOWEEN banner on startup
func (w *WitchConfig) PrintBanner() {
	sliceOfColorFuncs := []func(string, ...interface{}){
		color.Green,
		color.Red,
		color.Yellow,
	}

	//nolint:gosec // don't need crypto here for random, cmon it's 3 colors
	randomIndex := rand.Intn(len(sliceOfColorFuncs))
	colorOutput := sliceOfColorFuncs[randomIndex]
	colorOutput(banner)
}

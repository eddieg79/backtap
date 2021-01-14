package input

import (
	"io"
)

/*
Press power button:

/dev/input/event0: 0001 0074 00000001
/dev/input/event0: 0000 0000 00000000
/dev/input/event0: 0001 0074 00000000
/dev/input/event0: 0000 0000 00000000

*/

const (
	POWER_BUTTON EventCode = 0x0074
)

func PressPowerButton(output io.Writer) (err error) {
	return runEvents(output, []InputEvent{
		{
			Type:  EV_KEY,
			Code:  POWER_BUTTON,
			Value: DOWN,
		},
		eventSynReport,
		{
			Type:  PAUSE,
			Value: 25,
		},
		{
			Type:  EV_KEY,
			Code:  POWER_BUTTON,
			Value: UP,
		},
		eventSynReport,
	})
}

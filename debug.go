package main

import (
	"fmt"
	"time"
)

func debug(message string) {

	debugImpl(message)

}

func desktopDebug(message string) {

	if debugTextView != nil {

		debugText = fmt.Sprintf("\n > %s%s", message, debugText)

		// keep only the last 10 messages
		if len(debugText) > 1000 {
			debugText = debugText[:1000]
		}
	}

}

func terminalDebug(message string) {
	fmt.Println(message)
}

func updateDebug() {
	for {
		time.Sleep(500 * time.Millisecond)
		if debugTextView != nil {
			app.QueueUpdateDraw(func() {
				debugTextView.SetText(debugText)
			})
		}
	}
}

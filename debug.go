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

func refreshView() {

	initialLang := lang

	for {
		time.Sleep(50 * time.Millisecond)

		if debugTextView != nil {
			app.QueueUpdateDraw(func() {

				if initialLang != lang {
					loadLangFromJsonFile()
					initialLang = lang
				}

				timeNow = time.Now()

				calculateWorkingDayDuration()

				debugTextView.SetText(debugText)
			})
		}
	}
}

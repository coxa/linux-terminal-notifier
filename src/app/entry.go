package main

import (
	"flag"
	notify "github.com/mqu/go-notify"
)
import (
	"fmt"
	"os"
	"time"
)

const (
	DELAY = 3000
)

func main() {
	//arguments
	var title string
	var message string
	var group string
	var appIcon string
	var tmp string

	flag.StringVar(&title, "title", "Test Title", "Specify title for message.")
	flag.StringVar(&message, "message", "Test Message", "Specify message to display.")
	flag.StringVar(&group, "group", "", "Specify group for message.")
	flag.StringVar(&appIcon, "appIcon", "", "Specify app icon for notification.")


	flag.StringVar(&tmp, "subtitle", "", "Not implemented")
	flag.StringVar(&tmp, "contentImage", "", "Not implemented")
	flag.StringVar(&tmp, "execute", "", "Not implemented")
	flag.StringVar(&tmp, "open", "", "Not implemented")

	flag.Parse()
	//--arguments
	notify.Init("Hello World!")
	hello := notify.NotificationNew(title,
		message,
		appIcon)

	if hello == nil {
		fmt.Fprintf(os.Stderr, "Unable to create a new notification\n")
		return
	}

	notify.NotificationSetCategory(hello, group)

	// hello.SetTimeout(3000)
	notify.NotificationSetTimeout(hello, DELAY)

	// hello.Show()
	if e := notify.NotificationShow(hello); e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e.Message())
		return
	}

	time.Sleep(DELAY * 1000000)
	// hello.Close()
	notify.NotificationClose(hello)

	notify.UnInit()
}
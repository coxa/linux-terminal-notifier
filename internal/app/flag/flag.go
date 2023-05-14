package flag

import "flag"

type Flag struct {
	AppName string
	AppIcon string
	Title   string
	Message string
	Group   string
	Tmp     string
}

func New() *Flag {
	f := new(Flag)
	flag.StringVar(&f.Title, "title", "Test Title", "Specify title for message.")
	flag.StringVar(&f.Message, "message", "Test Message", "Specify message to display.")
	flag.StringVar(&f.Group, "group", "", "Specify group for message.")
	flag.StringVar(&f.AppIcon, "appIcon", "", "Specify app icon for notification.")
	flag.StringVar(&f.Tmp, "subtitle", "", "Not implemented")
	flag.StringVar(&f.Tmp, "contentImage", "", "Not implemented")
	flag.StringVar(&f.Tmp, "execute", "", "Not implemented")
	flag.StringVar(&f.Tmp, "open", "", "Not implemented")

	flag.Parse()

	return f
}

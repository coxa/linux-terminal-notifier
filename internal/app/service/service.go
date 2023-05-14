package service

import (
	"fmt"
	"github.com/mqu/go-notify"
	"os"
	"time"
)

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s Service) Notify(appName string, title string, message string, appIcon string, group string, delay int32) error {
	notify.Init(appName)
	notification := notify.NotificationNew(title,
		message,
		appIcon)

	if notification == nil {
		fmt.Fprintf(os.Stderr, "Unable to create a new notification\n")
		return nil
	}

	notify.NotificationSetCategory(notification, group)

	// notification.SetTimeout(3000)
	notify.NotificationSetTimeout(notification, delay)

	// notification.Show()
	if e := notify.NotificationShow(notification); e != nil {
		fmt.Fprintf(os.Stderr, "%s\n", e.Message())
		return nil
	}

	time.Sleep(1000000)
	// notification.Close()
	notify.NotificationClose(notification)

	notify.UnInit()
	return nil
}

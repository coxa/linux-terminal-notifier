package app

import (
	"github.com/coxa/linux-terminal-notifier/internal/app/flag"
	"github.com/coxa/linux-terminal-notifier/internal/app/service"
)

type App struct {
	s *service.Service
	f *flag.Flag
}

func New() (*App, error) {
	a := &App{}
	a.s = service.New()
	a.f = flag.New()

	return a, nil
}

func (a App) Run() error {
	err := a.s.Notify(a.f.AppName, a.f.Title, a.f.Message, a.f.AppIcon, a.f.Group, 10000)
	return err
}

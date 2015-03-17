package notificator

import (
	"os/exec"
	"runtime"
)

type Options struct {
	DefaultIcon string
	AppName     string
}

type notifier interface {
	push(title string, text string, iconPath string) *exec.Cmd
}

type Notificator struct {
	notifier    notifier
	defaultIcon string
}

func (n Notificator) Push(title string, text string, iconPath string) error {

	icon := n.defaultIcon

	if iconPath != "" {
		icon = iconPath
	}

	err := n.notifier.push(title, text, icon).Run()
	if err != nil {
		return err
	}
	return nil
}

type osxNotificator struct {
	AppName string
}

func (o osxNotificator) push(title string, text string, iconPath string) *exec.Cmd {
	return exec.Command("growlnotify", "-n", o.AppName, "--image", iconPath, "-m", title, text)
}

type linuxNotificator struct{}

func (l linuxNotificator) push(title string, text string, iconPath string) *exec.Cmd {
	return exec.Command("notify-send", "-i", iconPath, title, text)
}

type windowsNotificator struct{}

func (w windowsNotificator) push(title string, text string, iconPath string) *exec.Cmd {
	return exec.Command("growlnotify", "/i:", iconPath, "/t:", title, text)
}

func New(o Options) *Notificator {

	var notifier notifier

	switch runtime.GOOS {

	case "darwin":
		notifier = osxNotificator{AppName: o.AppName}
	case "linux":
		notifier = linuxNotificator{}
	case "windows":
		notifier = windowsNotificator{}

	}

	return &Notificator{notifier: notifier, defaultIcon: o.DefaultIcon}
}

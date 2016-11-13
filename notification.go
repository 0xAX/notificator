package notificator

import (
	"os/exec"
	"runtime"
	"fmt"
)

type Options struct {
	DefaultIcon string
	AppName     string
}

const (
	UR_NORMAL   =	"normal"
	UR_CRITICAL	=	"critical"
)

type notifier interface {
	push(title string, text string, iconPath string) *exec.Cmd
	pushCritical(title string, text string, iconPath string) *exec.Cmd
}

type Notificator struct {
	notifier    notifier
	defaultIcon string
}

func (n Notificator) Push(title string, text string, iconPath string, urgency string) error {
	icon := n.defaultIcon

	if iconPath != "" {
		icon = iconPath
	}

	if urgency == UR_CRITICAL {
		return n.notifier.pushCritical(title, text, icon).Run()
	}

	return n.notifier.push(title, text, icon).Run()

}

type osxNotificator struct {
	AppName string
}

func (o osxNotificator) push(title string, text string, iconPath string) *exec.Cmd {

	check_term_notif := exec.Command("which", "terminal-notifier");
	err := check_term_notif.Start()

	if err != nil {
		return exec.Command("terminal-notifier", "-title", o.AppName, "-message", title)
	} else {
		notification := fmt.Sprintf("display notification %s with title %s", o.AppName, title);
		return exec.Command("osascript", "-e", notification)
	}

	// return exec.Command("growlnotify", "-n", o.AppName, "--image", iconPath, "-m", title)
}

// Causes the notification to stick around until clicked.
func (o osxNotificator) pushCritical(title string, text string, iconPath string) *exec.Cmd {
	return exec.Command("notify-send", "-i", iconPath, title, text, "--sticky", "-p", "2")
}

type linuxNotificator struct{}

func (l linuxNotificator) push(title string, text string, iconPath string) *exec.Cmd {
	return exec.Command("notify-send", "-i", iconPath, title, text)
}

// Causes the notification to stick around until clicked.
func (l linuxNotificator) pushCritical(title string, text string, iconPath string) *exec.Cmd {
	return exec.Command("notify-send", "-i", iconPath, title, text, "-u", "critical")
}

type windowsNotificator struct{}

func (w windowsNotificator) push(title string, text string, iconPath string) *exec.Cmd {
	return exec.Command("growlnotify", "/i:", iconPath, "/t:", title, text)
}

// Causes the notification to stick around until clicked.
func (w windowsNotificator) pushCritical(title string, text string, iconPath string) *exec.Cmd {
	return exec.Command("notify-send", "-i", iconPath, title, text, "/s", "true", "/p", "2")
}


func New(o Options) *Notificator {

	var Notifier notifier

	switch runtime.GOOS {

	case "darwin":
		Notifier = osxNotificator{AppName: o.AppName}
	case "linux":
		Notifier = linuxNotificator{}
	case "windows":
		Notifier = windowsNotificator{}

	}

	return &Notificator{notifier: Notifier, defaultIcon: o.DefaultIcon}
}

package notification

import "reflect"
import "os/exec"

const defaultIcon = "icon/golang.png"

type Notificator struct {
    GnomeNotificator *gnomeNotificator
    KdeNotificator   *kdeNotificator
    OSXNotificator   *osxNotificator
    WindowsNotificator *windowsNotificator
}

type gnomeNotificator struct {
    IconPath string
    Title    string
    Text     string
}

type kdeNotificator struct {
    IconPath string
    Title    string
    Text     string
}

type osxNotificator struct {
    IconPath string
    AppName  string
    Title    string
    Text     string
}

type windowsNotificator struct {
    IconPath string
    Title    string
    Text     string
}

func Push(notificator interface{}) {
    // get notificator type
    n := reflect.TypeOf(notificator).String()
    // check notificator
    if  n == "*notification.gnomeNotificator" {
        notify := &Notificator{notificator.(*gnomeNotificator), nil, nil, nil}
        
        var icon string = ""

        if notify.GnomeNotificator.IconPath == "" {
            icon = defaultIcon
        } else {
            icon = notify.GnomeNotificator.IconPath
        }

        exec.Command("notify-send", "-i", icon, notify.GnomeNotificator.Title, notify.GnomeNotificator.Text)
        
    } else if n == "*notification.kdeNotificator" {
        notify := &Notificator{nil, notificator.(*kdeNotificator), nil, nil}

        var icon string = ""
        
        if notify.KdeNotificator.IconPath == "" {
            icon = defaultIcon
        } else {
            icon = notify.KdeNotificator.IconPath
        }

        exec.Command("kdialog", "--icon", icon, "--title", notify.KdeNotificator.Title, "--passivepopup", 
                     notify.KdeNotificator.Text)
    } else if n == "*notification.osxNotificator" {
        notify := &Notificator{nil, nil, notificator.(*osxNotificator), nil}

        var icon string = ""

        if notify.OSXNotificator.IconPath == "" {
            icon = defaultIcon
        } else {
            icon = notify.OSXNotificator.IconPath
        }

        exec.Command("growlnotify", "-n", notify.OSXNotificator.AppName, "--image", icon, "-m", 
                     notify.OSXNotificator.Title, notify.OSXNotificator.Text)
    } else {
        notify := &Notificator{nil, nil, nil, notificator.(*windowsNotificator)}
    
        var icon string
        
        if notify.WindowsNotificator.IconPath == "" {
            icon = defaultIcon
        } else {
            icon = notify.WindowsNotificator.IconPath
        }
        
        exec.Command("growlnotify", "/i:", icon, "/t:", notify.WindowsNotificator.Title, notify.WindowsNotificator.Text)
    }
}

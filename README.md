notificator
===========================

Desktop notification with golang for:

  * Windows with `growlnotify`;
  * Mac OS X with `growlnotify`;
  * Linux with `notify-send` for gnome and `kdialog` for kde.

usage
------

```go
package main

import "github.com/0xAX/notificator"

func main() {
  notificator.Push(&gnomeNotificator{"/home/user/icon.png", "title", "text"})
}
```

All notificators:

```go
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
```

todo
-----

  * Add more options for different notificators.

contribution
------------

  * Fork;
  * Make changes;
  * Send pull request;
  * Thank you.

author
----------

[@0xAX](https://twitter.com/0xAX)

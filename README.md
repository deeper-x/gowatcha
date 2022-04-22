# gowatcha
[![Go](https://github.com/deeper-x/gowatcha/actions/workflows/go.yml/badge.svg)](https://github.com/deeper-x/gowatcha/actions/workflows/go.yml)

[WIP]

File log watcher. Search for needle/occurrence, sending email notification when found.


Usage:
```bash
$ export USER=<EMAIL> 
$ export PASSWD=<PASSWORD> 
$ go run main.go <LOG_PATH> <RECIPIENT> <NEEDLE>

# example:
# zreq3b@SF314-59  (master)$ go run main.go /home/zreq3b/go/src/github.com/zreq3b/gowatcha/assets/demo.txt sviluppo@myskin.it 123

Starting...
   __________ _       _____  ______________  _____ 
  / ____/ __ \ |     / /   |/_  __/ ____/ / / /   |
 / / __/ / / / | /| / / /| | / / / /   / /_/ / /| |
/ /_/ / /_/ /| |/ |/ / ___ |/ / / /___/ __  / ___ |
\____/\____/ |__/|__/_/  |_/_/  \____/_/ /_/_/  |_|

===================================================

```

Parameters:
- [ ] LOG_PATH: absolute path to log file
- [ ] RECIPIENT: is the email to notify
- [ ] NEEDLE: the string occurrence we have to search

#TODO:
1. handle parameters 
2. handle configuration settings


![schema](https://i.ibb.co/2NCj3RC/Screenshot-from-2021-07-22-20-31-06.png)

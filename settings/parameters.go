package settings

import "time"

// NOTIFTXT is the notification text
const NOTIFTXT = "NOTIFICA"

// NOTIFINTVL is the do-not-send email interval
const NOTIFINTVL = 900

// SMTPSERVER mail auth server
const SMTPSERVER = "smtp.gmail.com"

// SMTPPORT to build socket
const SMTPPORT = "587"

// CHECKINTERVAL is the routine interval
const CHECKINTERVAL = time.Second * 10

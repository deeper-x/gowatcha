package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zreq3b/gowatcha/emailer"
	"github.com/zreq3b/gowatcha/lwatch"
	"github.com/zreq3b/gowatcha/rdb"
	"github.com/zreq3b/gowatcha/settings"
)

func main() {
	// instance name - #todo: handle cli args
	inst := os.Args[1]
	// path file to be logged - #todo: handle cli args
	path := os.Args[2]
	// recipient, email message - #todo: handle cli args
	rcpt := os.Args[3]
	//  occurrence we are searching for - #todo: handle cli args
	needle := os.Args[4]

	// Starting message
	fmt.Println(settings.MOTD)

	func() {
		// check logs every 10 secs
		for range time.Tick(settings.CHECKINTERVAL) {
			// build file instance
			lw, err := lwatch.New(path)
			if err != nil {
				log.Panic(err)
			}

			// read current file size, in order to compare with stored offset
			size, err := lw.GetFileSize()
			if err != nil {
				log.Panic(err)
			}

			// redis connection
			rdb := rdb.New()

			// reading last known offset
			offs, err := rdb.GetOffset(inst)
			if err != nil {
				log.Panic(err)
			}

			// if current size is greater than last known offset, something has been appended
			// and we have to grab it
			if size > offs {
				// step 1: get last part of file
				r, err := lw.GetTail(offs)
				if err != nil {
					log.Panic(err)
				}

				// step 2: read appended string in tail
				out, err := lw.GetAppendedString(size, r)
				if err != nil {
					log.Panic(err)
				}

				// step 3: define new offset, with current file size
				_, err = rdb.SetOffSet(inst, size)
				if err != nil {
					log.Panic(err)
				}

				// step 4: search for string occurrence
				if lw.NeedleExists(needle, out) {
					// step 4a: send email if found
					ls, err := rdb.GetLastSent(rcpt)
					if err != nil {
						log.Panic(err)
					}

					// step 4b: send notification if not already sent
					ok, err := emailer.Notify(rcpt, settings.NOTIFTXT, ls)
					if err != nil {
						log.Println(err)
					}

					// step 4c: if notification sent, update last sending TS
					if ok {
						rdb.WriteLastSent(rcpt)
					}
				}
			}
		}
	}()
}

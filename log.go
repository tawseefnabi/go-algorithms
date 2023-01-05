// Program in GO language to demonstrates how to use base log package.
package main

import (
	"log"
)

// The standard log entry contains below things:
// - a prefix (log.SetPrefix("LOG: "))
// - a datetime stamp (log.Ldate)
// - full path to the source code file writing to the log (log.Llongfile)
// - the line of code performing the write and finally the message.
func init() {
	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Println("init started")
}
func main() {
	// Println writes to the standard logger
	log.Println("main started")

	// Fatalln is Println() followed by a call to os.Exit(1)
	log.Fatalln("fatal message")

	// Panicln is Println() followed by a call to panic()
	log.Panicln("panic message")
}

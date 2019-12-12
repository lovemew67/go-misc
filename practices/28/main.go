package main

import (
	"log/syslog"
	"os"

	orilog "log"
	log "github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	airbrake "gopkg.in/gemnasium/logrus-airbrake-hook.v2"

	"github.com/sirupsen/logrus"
)

func init() {
	orilog.SetFlags(orilog.LstdFlags | orilog.LUTC | orilog.Lshortfile | orilog.Lmicroseconds)
	orilog.Println("Hello World: 28")

	log.SetOutput(os.Stdout)
	log.Println("Hello World: 28, one")

	log.SetFormatter(&log.TextFormatter{})
	log.Println("Hello World: 28, two")

	log.SetFormatter(&log.JSONFormatter{})
	log.Println("Hello World: 28, three")

	log.SetLevel(log.DebugLevel)

	// Use the Airbrake hook to report errors that have Error severity or above to
	// an exception tracker. You can create custom hooks, see the Hooks section.
	log.AddHook(airbrake.NewHook(123, "xyz", "production"))
	hook, err := logrus_syslog.NewSyslogHook("udp", "localhost:514", syslog.LOG_INFO, "")
	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}

	handler := func() {
		// not for panic
		orilog.Println("gracefully shutdown something...")
	}
	logrus.RegisterExitHandler(handler)
}

func main() {
	log.WithFields(log.Fields{
		"msg": "Hello World: 28",
	}).Trace("WTF")

	newInstance := logrus.New()
	requestLogger := newInstance.WithFields(logrus.Fields{
		"aaa": "bbb",
		"ccc": "ddd",
	})
	newInstance.WithFields(logrus.Fields{
		"ccc": "eee",
	}).Info("fff")
	newInstance.Info("gggg")
	requestLogger.Info("1111")

	log.Trace("Something very low level.")
	log.Debug("Useful debugging information.")
	log.Info("Something noteworthy happened!")
	log.Warn("You should probably take a look at this.")
	log.Error("Something failed but I'm not quitting.")
	// log.Fatal("Bye.") // Calls os.Exit(1) after logging
	// log.Panic("I'm bailing.") // Calls panic() after logging
}

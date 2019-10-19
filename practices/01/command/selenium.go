package command

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/tebeka/selenium"
)

const (
	// seleniumPathMac     = "/Users/zzh/Golang/gopath/src/github.com/lovemew67/go-misc/binary/selenium-server-standalone-2.45.0.jar"
	// chromeDriverPathMac = "/Users/zzh/Golang/gopath/src/github.com/lovemew67/go-misc/binary/chromedriver-75.0.3770.140"
	seleniumPathWin     = "C:/Golang/gopath/src/github.com/lovemew67/go-misc/01/binary/selenium-server-standalone-2.45.0.jar"
	chromeDriverPathWin = "C:/Golang/gopath/src/github.com/lovemew67/go-misc/01/binary/chromedriver-75.0.3770.140.exe"
	port                = 8080
	withServer          = true
)

func NewSeleniumCommand() *cobra.Command {
	var seleniumCmd = &cobra.Command{
		Use:   "selenium",
		Short: "start selenium operation",
		Long:  `start selenium operation`,
		Run: func(cmd *cobra.Command, args []string) {
			var opts []selenium.ServiceOption
			if withServer {
				opts = []selenium.ServiceOption{
					// Enable fake XWindow session.
					// selenium.StartFrameBuffer(),
					selenium.ChromeDriver(chromeDriverPathWin),
					selenium.Output(os.Stderr), // Output debug information to STDERR
				}
			} else {
				opts = []selenium.ServiceOption{
					// Enable fake XWindow session.
					// selenium.StartFrameBuffer(),
					selenium.Output(os.Stderr), // Output debug information to STDERR
				}
			}

			// selenium.SetDebug(true)
			var service *selenium.Service
			var err error
			if withServer {
				service, err = selenium.NewChromeDriverService(chromeDriverPathWin, port, opts...)
				if err != nil {
					panic(err)
				}
			} else {
				service, err = selenium.NewSeleniumService(seleniumPathWin, port, opts...)
				if err != nil {
					panic(err)
				}
			}
			defer func() { _ = service.Stop() }()

			caps := selenium.Capabilities{"browserName": "chrome"}
			wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://127.0.0.1:%d/wd/hub", port))
			if err != nil {
				panic(err)
			}
			defer func() { _ = wd.Quit() }()

			startOperation(wd)
		},
	}
	return seleniumCmd
}

// Credit to:
//	https://michaelchen.tech/selenium/manipulate-selenium-with-golang/
// 	https://www.golangnote.com/topic/230.html

func startOperation(wd selenium.WebDriver) {

	// open web page
	if err := wd.Get("https://tw.yahoo.com"); err != nil {
		panic(err)
	}

	// Get a reference to the text box containing code.
	elem, err := wd.FindElement(selenium.ByCSSSelector, "#UHSearchBox")
	if err != nil {
		panic(err)
	}

	// Remove the boilerplate code already in the text box.
	if err := elem.Clear(); err != nil {
		panic(err)
	}

	// Enter some new code in text box.
	if err = elem.SendKeys(`golang selenium scroll`); err != nil {
		panic(err)
	}

	// Click the run button.
	btn, err := wd.FindElement(selenium.ByCSSSelector, "#UHSearchWeb")
	if err != nil {
		panic(err)
	}
	if err := btn.Click(); err != nil {
		panic(err)
	}

	// Wait for the program to finish running and get the output.
	outputDiv, err := wd.FindElement(selenium.ByCSSSelector, ".pl-13")
	if err != nil {
		panic(err)
	}

	// Check got value
	var output string
	for {
		output, err = outputDiv.Text()
		if err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 5)
		fmt.Println("got: " + output)
		break
	}

	// Click
	if err := outputDiv.Click(); err != nil {
		panic(err)
	}

	// Scroll
	lastHeight, _ := wd.ExecuteScript("return document.body.scrollHeight", nil)
	for {
		// Scroll down to bottom
		_, _ = wd.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)

		// Wait to load page
		time.Sleep(time.Second * 5)

		// Calculate new scroll height and compare with last scroll height
		newHeight, _ := wd.ExecuteScript("return document.body.scrollHeight", nil)
		if newHeight == lastHeight {
			break
		}
		lastHeight = newHeight
	}

	// Current URL
	url, _ := wd.CurrentURL()
	title, _ := wd.Title()
	fmt.Println(url)
	fmt.Println(title)

}

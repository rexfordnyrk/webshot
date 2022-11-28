// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package capture

import (
	"bufio"
	"context"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func GetBatchScreenShot(filepath string, c *Config) {
	file, err := os.Open(filepath)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if len(strings.TrimSpace(scanner.Text())) == 0{
			continue
		}
		GetSingleScreenShot(scanner.Text(),c)
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}


func GetSingleScreenShot(site string, c *Config) {
	//byte slice to hold captured images in bytes
	var buf []byte
	log.Printf("................making request for screenshot using %s", site)

	//setting options for headless chrome to execute with
	var options []chromedp.ExecAllocatorOption
	options = append(options, chromedp.WindowSize(c.Width, c.Height))
	options = append(options, chromedp.DefaultExecAllocatorOptions[:]...)

	//setup context with options
	actx, acancel := chromedp.NewExecAllocator(context.Background(), options...)

	defer acancel()

	// create context
	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	tasks:= chromedp.Tasks{
		//loads page of the URL
		chromedp.Navigate(site),

		//waits for 5 secs
		chromedp.Sleep(5*time.Second),

		//Captures Screenshot with size, fullscreen or not
		getCaptureSizeElement(c.Size, &buf),
	}

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, tasks); err != nil {
		log.Fatal(err)
	}

	//naming file using provided URL without "/"s and current unix datetime
	filename := getFileName(site,c.Format)

	//write byte slice data of standard screenshot to file
	if err := ioutil.WriteFile(filename, buf, 0644); err != nil {
		log.Fatal(err)
	}

	//log completion and file name to
	log.Printf("..............saved screenshot to file %s", filename)
}

func getCaptureSizeElement(c string, buf *[]byte) chromedp.Action {
	if strings.EqualFold(c, "fullscreen") {
		return chromedp.FullScreenshot(buf,100)
	}
	return chromedp.CaptureScreenshot(buf)
}

func getFileName(url string, ext string) string {
	url = strings.Replace(url,"://","-",-1)
	url = strings.Replace(url,"/","-",-1)

	return 	fmt.Sprintf("%s_%d.%s",url, time.Now().UTC().Unix(), ext)
}
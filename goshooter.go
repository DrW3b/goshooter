package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	fileFlag    = flag.String("f", "", "file contains urls or domains")
	timeoutFlag = flag.Duration("t", 30*time.Second, "timeout")
	threadsFlag = flag.Int("th", 5, "number of threads")
)

const (
	colorRed   = "\033[31m"
	colorGreen = "\033[32m"
	colorReset = "\033[0m"
)

func main() {
	flag.Parse()

	printPattern(colorRed)

	log.SetOutput(os.Stdout)

	data, err := ioutil.ReadFile(*fileFlag)
	if err != nil {
		log.Fatal(err)
	}

	urls := strings.Split(string(data), "\n")

	date := time.Now().Format("2006-01-02_15-04-05")
	dir := fmt.Sprintf("screenshots_%s", date)

	err = os.MkdirAll(dir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	urlCh := make(chan string)

	for i := 0; i < *threadsFlag; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for url := range urlCh {
				ctx, _ := chromedp.NewContext(context.Background())
				takeScreenshot(ctx, dir, url)
			}
		}()
	}

	for _, url := range urls {
		urlCh <- url
	}

	close(urlCh)
	wg.Wait()
}

func printPattern(color string) {
	// The pattern to be printed
	pattern := fmt.Sprintf("%s\n ██████╗  ██████╗ ███████╗██╗  ██╗ ██████╗  ██████╗ ████████╗███████╗██████╗\n", color)
	pattern += fmt.Sprintf("██╔════╝ ██╔═══██╗██╔════╝██║  ██║██╔═══██╗██╔═══██╗╚══██╔══╝██╔════╝██╔══██╗\n")
	pattern += fmt.Sprintf("██║  ███╗██║   ██║███████╗███████║██║   ██║██║   ██║   ██║   █████╗  ██████╔╝\n")
	pattern += fmt.Sprintf("██║   ██║██║   ██║╚════██║██╔══██║██║   ██║██║   ██║   ██║   ██╔══╝  ██╔══██╗\n")
	pattern += fmt.Sprintf("╚██████╔╝╚██████╔╝███████║██║  ██║╚██████╔╝╚██████╔╝   ██║   ███████╗██║  ██║\n")
	pattern += fmt.Sprintf(" ╚═════╝  ╚═════╝ ╚══════╝╚═╝  ╚═╝ ╚═════╝  ╚═════╝    ╚═╝   ╚══════╝╚═╝  ╚═╝%s\n", colorReset)
	pattern += fmt.Sprintf("%s%-17s%s\n", colorRed, "                          goshooter v1.0", colorReset)
	pattern += fmt.Sprintf("%s%-13s%s\n", colorRed, "                          Coded by DrW3B", colorReset)
	pattern += fmt.Sprintf("%s                          Telegram: @DrW33B_Xo%s\n", colorRed, colorReset)

	fmt.Println(pattern)
}




func takeScreenshot(ctx context.Context, dir, url string) {
	if url == "" {
		return
	}

	fullPath := filepath.Join(dir, url, "screenshot.png")

	err := os.MkdirAll(filepath.Dir(fullPath), 0755)
	if err != nil {
		log.Fatalf("Failed to create directory: %v", err)
	}

	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		log.Fatalf("%sInvalid URL: %s. Please provide a domain without protocol.%s\n", colorRed, url, colorReset)
	}

	var buf []byte
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate(`http://` + url),
		chromedp.Sleep(*timeoutFlag),
		chromedp.CaptureScreenshot(&buf),
	})
	if err != nil {
		log.Fatalf("%sFailed to capture screenshot: %v%s\n", colorRed, err, colorReset)
	}

	if err := ioutil.WriteFile(fullPath, buf, 0644); err != nil {
		log.Fatalf("%sFailed to write screenshot to file: %v%s\n", colorRed, err, colorReset)
	}

	log.Printf("%sScreenshot of %s saved%s\n", colorGreen, url, colorReset)
}

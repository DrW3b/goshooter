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

func main() {
	flag.Parse()

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

func takeScreenshot(ctx context.Context, dir, url string) {
	
	if url == "" {
		return
	}

	fmt.Printf("Taking screenshot of %s\n", url)

	
	fullPath := filepath.Join(dir, url, "screenshot.png")

	
	err := os.MkdirAll(filepath.Dir(fullPath), 0755)
	if err != nil {
		log.Fatal(err)
	}

	
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		log.Fatalf("Invalid URL: %s. Please provide a domain without protocol.", url)
	}

	
	var buf []byte
	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.EmulateViewport(1920, 1080),
		chromedp.Navigate(`http://` + url),
		chromedp.Sleep(*timeoutFlag),
		chromedp.CaptureScreenshot(&buf),
	})
	if err != nil {
		log.Fatal(err)
	}

	if err := ioutil.WriteFile(fullPath, buf, 0644); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Screenshot of %s saved\n", url)
}


package main

import (
	"bufio"
	"bytes"
	//"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	//"reflect"
	"regexp"
	//"strings"
	"strconv"
	"sync"
	//"time"
)

var p = fmt.Print
var pL = fmt.Println //Alias for print Line
var user = "admin"
var pass = ""
var wg sync.WaitGroup

/*
--video boundary--
Content-length: 57875
Date: 11-17-2016 04:03:05 PM IO_00000000_PT_005_000
Content-type: image/jpeg
*/

func goGetemAll(url, user, pass, basename string, n int) {
	defer wg.Done()
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	contentType := resp.Header.Get("Content-Type")
	boundaryRe := regexp.MustCompile(`^multipart/x-mixed-replace;boundary=(.*)$`)
	boundary := boundaryRe.FindStringSubmatch(contentType)[1]
	buffer := bufio.NewReader(resp.Body)

	for j := 0; j < n; j++ {
		p(basename + " run ")
		pL(j)

		//delimiter := fmt.Sprintf("--%s\r\n", boundary)
		delimiter := "--" + boundary + "\r\n"
		data := make([]byte, 0)
		var imgSize int64
		for i := 0; i < 4; i++ {
			d, _ := buffer.ReadBytes('\n')
			if j == 0 {		//Black magic
				if i ==1 {	//Black magic
					//Get the raw decimal numbers from Content-length and convert to int64
					imgSize, _ = strconv.ParseInt(string(d[16:len(d)-2]), 10, 32)
					/*
					p("imgSize: ")
					pL(imgSize)
					*/
				}
			}else if i==0 {
				imgSize, _ = strconv.ParseInt(string(d[16:len(d)-2]), 10, 32)
				
			}
		}
		p(basename + " imgSize: ")
		pL(imgSize)
		
		for {
			line, _ := buffer.ReadBytes('\n')
			found := bytes.HasSuffix(line, []byte(delimiter))
			if (string(line)) == "\r\n" {
				continue
			}
			if found == true {
				data = append(data, line[:(len(line)-len(delimiter))]...)
				//strTime := t.Format("20060102150405.0000")
				filename := basename + "_" + strconv.Itoa(j) + ".jpg"
				ioutil.WriteFile(filename, data[:len(data)-2], 0644)
				break
			} else {
				data = append(data, line...)
			}
		}
	}
	resp.Body.Close()
}

func init() {
	//url = os.Getenv("url")
	/*if len(url) < 15 {
		os.Exit(2)
	}*/
	pass = os.Getenv("pass")
	if len(pass) < 1 {
		os.Exit(2)
	}
}

func main() {
	pL("Start!")
	var urls = []string{
		"http://187.183.214.57:8101/video.cgi",
		"http://187.183.214.57:8102/video.cgi",
		"http://187.183.214.57:8103/video.cgi",
		"http://187.183.214.57:8104/video.cgi",
		"http://187.183.214.57:8105/video.cgi",
		"http://187.183.214.57:8107/video.cgi",
	}
	var baseNames = []string{
		"Camera01",
		"Camera02",
		"Camera03",
		"Camera04",
		"Camera05",
		"Camera07",
	}
	for i, url := range urls{
		wg.Add(1)
		go goGetemAll(url, user, pass, baseNames[i], 5)
	}
	wg.Wait()
}
	/*for {
		if done {
			os.Exit(0)
		}
	}*/

	//ticker := time.NewTicker(time.Millisecond * 1000)
	//go func() {
		//for t := range ticker.C {
			//pL(t)
			

			/*fmt.Print("contentType: ")
			pL(contentType)
			fmt.Print("contentType Len: ")
			pL(len(contentType))
			fmt.Print("contentType Type: ")
			pL(reflect.TypeOf(contentType))
			fmt.Print("boundaryRe Type: ")
			pL(reflect.TypeOf(boundaryRe))
			fmt.Print("boundary: ")
			pL(boundary)
			fmt.Print("boundary Len: ")
			pL(len(boundary))
			fmt.Print("boundary Type: ")
			pL(reflect.TypeOf(boundary))
			fmt.Print("buffer Type: ")
			pL(reflect.TypeOf(buffer))*/

			
			//
			//j++
	//}()
	//time.Sleep(time.Millisecond * 62000)
    //ticker.Stop()
    //fmt.Println("Ticker stopped")

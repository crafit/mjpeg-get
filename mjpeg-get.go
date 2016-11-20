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
	//"time"
)

var p = fmt.Print
var pL = fmt.Println //Alias for print Line
//var boundDelimiter = "--video boundary--"
//var boundDelimiter = "video"
var url = "http://192.168.0.101:8101/video.cgi"
//var url = "http://danasmikro.dlinkddns.com:8105/video.cgi"
//var url = "http://casadecarnesdoge.dlinkddns.com:8107/video.cgi"
var user = "admin"

/*
--video boundary--
Content-length: 57875
Date: 11-17-2016 04:03:05 PM IO_00000000_PT_005_000
Content-type: image/jpeg
*/

func main() {
	pL("Start!")
	pass := os.Args[1]
	//j := 0
	//url := os.Args[1]
	/*var url = "http://192.168.0.101:8101/video.cgi"
	var user = "admin"
	var pass = "DMKInfo2012"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
		//req.Header.Add("Authorization", "Basic YWRtaW46RE1LSW5mbzIwMTI=")
	req.SetBasicAuth(user, pass)
	resp, err := client.Do(req)
	//resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}
	contentType := resp.Header.Get("Content-Type")
	pL(contentType)
	pL(resp.Status)
	pL(reflect.TypeOf(resp.Body))
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	pL(buf)*/
	//	pL(resp.ContentLength)

	//boundaryRe := regexp.MustCompile(`^multipart/x-mixed-replace;boundary=(.*)$`)
	//boundary := boundaryRe.FindStringSubmatch(contentType)[1]
	//boundary := resp.Header.Get("Content-Type").FindStringSubmatch("--")
	/*buffer := bufio.NewReader(resp.Body)

	delimiter := fmt.Sprintf("--%s\r\n", boundary)
	data := make([]byte, 0)
	for i := 0; i < 4; i++ {
		buffer.ReadBytes('\n')
	}

	for {
		line, _ := buffer.ReadBytes('\n')
		found := bytes.HasSuffix(line, []byte(delimiter))
		if found == true {
			data = append(data, line[:(len(line)-len(delimiter))]...)
			ioutil.WriteFile("frame.jpg", data, 0644)
			break
		} else {
			data = append(data, line...)
		}
	}*/

	/*resp, err := http.Get(os.Args[1])
	if err != nil {
		panic(err)
	}*/

	//ticker := time.NewTicker(time.Millisecond * 1000)
	//go func() {
		//for t := range ticker.C {
			//pL(t)
			client := &http.Client{}
			req, err := http.NewRequest("GET", url, nil)
			//req.Header.Add("Authorization", "Basic YWRtaW46RE1LSW5mbzIwMTI=")
			req.SetBasicAuth(user, pass)
			resp, err := client.Do(req)
			if err != nil {
				panic(err)
			}
			contentType := resp.Header.Get("Content-Type")
			boundaryRe := regexp.MustCompile(`^multipart/x-mixed-replace;boundary=(.*)$`)
			boundary := boundaryRe.FindStringSubmatch(contentType)[1]
			buffer := bufio.NewReader(resp.Body)

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

			for j := 0; j < 30; j++ {
			fmt.Print("Run ")
			pL(j)

			//delimiter := fmt.Sprintf("--%s\r\n", boundary)
			delimiter := "--" + boundary + "\r\n"
			data := make([]byte, 0)
			for i := 0; i < 4; i++ {
				d, _ := buffer.ReadBytes('\n')
				//pL(string(d))
				if j == 0 {
					if i ==1 {
						//Get the raw decimal numbers from Content-length and convert to int64
						imgSize, _ := strconv.ParseInt(string(d[16:len(d)-2]), 10, 32)
						p("imgSize: ")
						pL(imgSize)
					}
				//"Content-length: "
				//contentLenghtRe := regexp.MustCompile(`^Content-length: (.*)$`)
				//contentLenght := contentLenghtRe.FindStringSubmatch(string(d))
				//pL(contentLenght)
				//if d[:14] == "Content-length" {
					
				}else if i==0 {
					imgSize, _ := strconv.ParseInt(string(d[16:len(d)-2]), 10, 32)
					p("imgSize: ")
					pL(imgSize)
				}
			}
			for {
				line, _ := buffer.ReadBytes('\n')
				found := bytes.HasSuffix(line, []byte(delimiter))
				if (string(line)) == "\r\n" {
					continue
				}
				if found == true {
					data = append(data, line[:(len(line)-len(delimiter))]...)
					//pL(string(j))
					//strTime := t.Format("20060102150405.0000")
					//filename := "Ge07_" + strconv.Itoa(j) + ".jpg"
//					filename := "DanasMikro05_" + strconv.Itoa(j) + ".jpg"
					//filename := "danasmikro05_" + strTime + ".jpg"
					filename := "DMK01_" + strconv.Itoa(j) + ".jpg"
					//pL(j)
					//pL(reflect.TypeOf(j))
					//pL(string(j))
					//pL(reflect.TypeOf(string(j)))
					//pL(filename)
					ioutil.WriteFile(filename, data[:len(data)-2], 0644)
					break
				} else {
					data = append(data, line...)
					/*if linha1 {
						data = data[2:]
						linha1 = false
					}*/
				}
			}
			//resp.Body.Close()
			//j++
		}
	//}()
	//time.Sleep(time.Millisecond * 62000)
    //ticker.Stop()
    //fmt.Println("Ticker stopped")
}

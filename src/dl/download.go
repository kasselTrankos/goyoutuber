package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
	"net/url"
	"strings"
	"log"
)

type stream map[string]string

func toJSON(m interface{}) string {
  	js, err := json.Marshal(m)
  	if err != nil {
  		log.Fatal(err)
  	}
  	return strings.Replace(string(js), ",", ", ", -1)
  }
func main(){
	vId := flag.String("yl", "Xq_a8f24UJI", "introduce el video")
	flag.Parse()
	baseUri := "http://youtube.com/get_video_info?video_id="

	uri := bytes.NewBufferString("")
	uri.WriteString(baseUri)
	uri.WriteString(*vId)
	fmt.Println(uri.String())
	// fmt.Println("Este es el id de video elegido? :",*vId)

	resp, err := http.Get(uri.String())
	if err !=nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	videoInfo := string(body)
	// fmt.Println(videoInfo)
	answer, err := url.ParseQuery(videoInfo)

	status, ok := answer["status"]

	if !ok {
		fmt.Println("no response status found in the server's answer")
	}
	if status[0] == "fail" {
		reason, ok := answer["reason"]
		if ok {
			fmt.Println("'fail' response status found in the server's answer, reason: '%s'", reason[0])
		} else {
			fmt.Println("'fail' response status found in the server's answer, no reason given")
		}
	}
	if status[0] != "ok" {
		fmt.Println("non-success response status found in the server's answer (status: '%s')", status)
	}

	streamMap, ok := answer["url_encoded_fmt_stream_map"]

	// fmt.Println(toJSON(answer))
	if !ok {
		fmt.Println("no stream map found in the server's answer")

		// return err
	}else{
		fmt.Println(" vamos a por el stream que no se ni lo qure es aun")
		var streams []stream
		streamsList := strings.Split(streamMap[0], ",")
		for streamPos, streamRaw := range streamsList {
			streamQry, err := url.ParseQuery(streamRaw)
			if err != nil {
				log.Printf("An error occured while decoding one of the video's stream's information: stream %d: %s\n", streamPos, err)
				continue
			}
			var sig string
			if _, exist := streamQry["sig"]; exist {
				sig = streamQry["sig"][0]
			}

			streams = append(streams, stream{
				"quality": streamQry["quality"][0],
				"type":    streamQry["type"][0],
				"url":     streamQry["url"][0],
				"sig":     sig,
				"title":   answer["title"][0],
				"author":  answer["author"][0],
			})
		}

		lasturl := streams[0]["url"] + "&signature=" + streams[0]["sig"]
		///Estamos very cerca!!
		response, err := http.Get(lasturl)
		fmt.Println(lasturl)
		if err != nil {
			fmt.Println("Http.Get\nerror: %s\nurlBrokebn: %s\n", err, lasturl)
		}
		defer response.Body.Close()
		fmt.Println("Cerramos el body refer! aún no cantamos victoria!!!!")
		// var contentLength = float64(response.ContentLength)
		if response.StatusCode != 200 {
			fmt.Println("reading answer: non 200[code=%v] status code received: '%v'", response.StatusCode, err)
		}

	}






}


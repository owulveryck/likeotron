// Copyright 2016 Olivier Wulveryck
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Topic string    `json:"topic"`
	Date  time.Time `json:"date"`
	Total int64     `json:"total"`
	Score float64   `json:"score"`
}

type Message struct {
	Topic  string    `json:"topic"`
	Sender string    `json:"sender"`
	Msg    string    `json:"message"`
	Like   bool      `json:"like"`
	Date   time.Time `json:"-"`
}

type Msg struct {
	Name  string `json:"name"`
	State string `json:"state"`
}

type Communication struct {
	Msg  Msg
	Chan chan Msg
}

var topics map[string][]int64

func init() {
	topics = make(map[string][]int64, 0)
}

var upgrader = websocket.Upgrader{} // use default options

var communication = make(chan Communication)

func phone(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return

	}
	defer c.Close()

	// Read incoming message and pass it to the hub
	var channel = make(chan Msg)
	// launch a goroutine and wait
	go func(channel chan Msg) {
		for {
			response := <-channel
			log.Println("about to send ", response)

			err = websocket.WriteJSON(c, response)
			if err != nil {
				log.Println("Unable to send message", err)
			}
		}
	}(channel)
	for {

		var message Msg
		err := websocket.ReadJSON(c, &message)
		if err != nil {
			log.Println("Unable to read message", err)
		} else {
			log.Printf("=> %v is talking", message.Name)
			communication <- Communication{Msg: message, Chan: channel}
			log.Printf("=> Advertized ")

		}
	}
}

func progress(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return

	}
	defer c.Close()

	for {

		var message Message
		err := websocket.ReadJSON(c, &message)
		if err != nil {
			log.Println("Unable to read message", err)
		} else {
			log.Printf("=> Message: \n==> Topic:%v\n==>Sender:%v\n==>Date:%v ", message.Topic, message.Sender, message.Date)
		}

		var increment int64
		if message.Like {
			increment = 1
		} else {
			increment = -1
		}
		if _, ok := topics[message.Topic]; !ok {
			topics[message.Topic] = []int64{
				1,
				1,
			}
		}
		topics[message.Topic] = []int64{
			topics[message.Topic][0] + 1,
			topics[message.Topic][1] + increment,
		}
		var response Result
		response.Topic = message.Topic
		response.Date = time.Now()
		response.Total = topics[message.Topic][0]
		response.Score = float64(topics[message.Topic][1] * 100 / topics[message.Topic][0])
		log.Println("about to send ", response)

		err = websocket.WriteJSON(c, response)
		if err != nil {
			log.Println("Unable to send message", err)
		}
	}
}

func GetJson(w http.ResponseWriter, r *http.Request) {
}

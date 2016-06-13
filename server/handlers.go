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
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"time"
)

type Message struct {
	Topic  string    `json:"topic"`
	Sender string    `json:"sender"`
	Msg    string    `json:"message"`
	Like   bool      `json:"like"`
	Date   time.Time `json:"-"`
}

func echoHandler(ws *websocket.Conn) {

	for {

		var message Message
		err := websocket.JSON.Receive(ws, &message)
		if err != nil {
			log.Println("Unable to read message", err)
		}

		message.Sender = "Server"
		err = websocket.JSON.Send(ws, message)
		if err != nil {
			log.Println("Unable to send message", err)
		}
	}
}

func GetJson(w http.ResponseWriter, r *http.Request) {
}

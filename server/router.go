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
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	router.
		Methods("GET").
		Path("/json/{jsonContent}").
		Name("Dynamic").
		HandlerFunc(GetJson)

	router.
		Methods("GET").
		Path("/phone").
		Name("WebSocket").
		HandlerFunc(phone)
	router.
		Methods("GET").
		Path("/progress").
		Name("WebSocket").
		HandlerFunc(progress)

	router.
		Methods("GET").
		PathPrefix("/").
		Name("Static").
		Handler(http.FileServer(http.Dir("./htdocs")))
	go func() {
		type tempo struct {
			Channel chan Msg
			State   string
		}
		var attendee = make(map[string]tempo)
		for {
			message := <-communication
			log.Println("Message received ", message)
			attendee[message.Msg.Name] = tempo{Channel: message.Chan, State: message.Msg.State}
			for att, temp := range attendee {
				go func(att string, temp tempo) {
					channel := temp.Channel
					state := temp.State
					log.Println(state)
					if state == "start" {
						log.Printf("Sending to %v on channel %v", att, channel)
						channel <- Msg{"A", "running"}
					}
					if state == "stop" {
						log.Printf("Sending to %v on channel %v", att, channel)
						channel <- Msg{"A", "stopped"}
					}
				}(att, temp)
			}
		}
	}()

	return router

}

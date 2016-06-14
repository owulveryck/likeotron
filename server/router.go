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
		var attendee = make(map[string]chan Msg)
		for {
			message := <-communication
			log.Println("Message received ", message)
			if _, ok := attendee[message.Msg.Name]; !ok {
				attendee[message.Msg.Name] = message.Chan
			}
			for att, channel := range attendee {
				go func(att string, channel chan Msg) {
					log.Printf("Sending to %v on channel %v", att, channel)
					channel <- Msg{"A", "autonomous"}
				}(att, channel)
			}
		}
	}()

	return router

}

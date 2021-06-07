package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.RequestURI
	fmt.Printf(" [#] Decoding: %s\n", path)
	if path[1] == '/' {
		// Remove leading '/', if exists
		path = path[1:]
	}
	if path[1:3] == "v3" {
		// Handle URL defense v3 encoding
		// __<URL>__;!!<metadata>$
		dsturl := ""
		metadata := ""
		state := 0
		valid := false
		index := 3
		for index < len(path) {
			ch := path[index]
			switch state {
			case 0:
				if ch != '_' {
					break
				}
				state++
			case 1:
				if ch != '_' {
					break
				}
				state++
			case 2:
				if ch == '_' {
					// Check if encoded URL is over
					state++
				} else {
					if len(dsturl) >= 2 && dsturl[len(dsturl)-2:] == ":/" && ch != '/' {
						// Insert missing slash
						dsturl = dsturl + "/"
					}
					dsturl = dsturl + string(ch)
				}
			case 3:
				if ch != '_' {
					// Handle the case of having '_' in the URL
					dsturl = dsturl + "_" + string(ch)
					state = 2
				} else {
					state++
				}
			case 4:
				if ch != ';' {
					// Handle the case of having "__" in the URL
					dsturl = dsturl + "__" + string(ch)
					state = 2
				} else {
					state++
				}
			case 5:
				if ch != '!' {
					// Handle the case of having "__;" in the URL
					dsturl = dsturl + "__;" + string(ch)
					state = 2
				} else {
					state++
				}
			case 6:
				if ch != '!' {
					// Handle the case of having "__;!" in the URL
					dsturl = dsturl + "__;!" + string(ch)
					state = 2
				} else {
					state++
				}
			case 7:
				// Found a valid instance of "__<URL>__;!!"
				state++
				fallthrough
			case 8:
				if ch == '$' {
					// Found a valid instance of "__<URL>__;!!<metadata>$"
					valid = true
					break
				}
				metadata = metadata + string(ch)
			}
			index++
		}
		if valid {
			fmt.Printf(" [*] Redirecting to: %s\n", dsturl)
			// Send redirect with code 301 "permanently moved"
			http.Redirect(w, r, dsturl, 301)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServeTLS("localhost:443", "server.crt", "server.key", nil))
	//log.Fatal(http.ListenAndServe("localhost:80", nil))
}

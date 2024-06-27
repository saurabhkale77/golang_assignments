package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type StatusChecker interface {
	Check(name string) (status bool, err error)
}

// for use of body parsing
type httpChecker struct {
	Websites []string `json:"websites"`
}

// to maintain website and status map
var websiteVsStatusMap = make(map[string]string)

func main() {
	fmt.Printf("Starting server....\n")

	http.HandleFunc("/websites", websitesListAndStatusHandler)

	http.ListenAndServe("127.0.0.1:3000", nil)
}

var websitesList = httpChecker{}

func websitesListAndStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//  curl -v -X POST -d '{"websites" : ["http://www.google.com", "http://www.facebook.com", "http://hoiubefvbwoihbdfv.com", "http://127.0.0.1:8080/welcome"]}' http://127.0.0.1:3000/websites - right

		if err := json.NewDecoder(r.Body).Decode(&websitesList); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		fmt.Printf("\n\nwebsitesList.Websites - %v\n", websitesList.Websites)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Websites received"))

		go checkTheStatusOfWebsites()

		go checkTheStatusForEveryOneMinute()
	} else if r.Method == "GET" {
		r.ParseForm()
		requestQueryParams := r.Form

		nameOfWebsite := requestQueryParams["name"]
		// len() function returns zero for a nil slice.
		// so no need to check nameOfWebsite != nil
		if len(nameOfWebsite) != 0 {
			if _, isExists := websiteVsStatusMap[nameOfWebsite[0]]; !isExists {
				fmt.Fprintf(w, "Website entered does not exist in the list!")
				return
			}

			jsonData, err := json.Marshal(map[string]string{
				nameOfWebsite[0]: websiteVsStatusMap[nameOfWebsite[0]],
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")

			w.WriteHeader(http.StatusOK)
			w.Write(jsonData)

			// fmt.Fprint(w, websiteVsStatusMap[nameOfWebsite[0]])
			return
		}

		jsonData, err := json.Marshal(websiteVsStatusMap)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)

		// fmt.Fprintf(w, fmt.Sprintf("%v", websiteVsStatusMap))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (h *httpChecker) Check(url string) (status bool, err error) {
	// check a single website
	defer func() {
		if r := recover(); r != nil {
			status = false
			err = fmt.Errorf(fmt.Sprintf("%v", r))
		}
	}()

	resp, err := http.Get(url)
	if err != nil {
		return false, err
	}

	if resp.StatusCode != http.StatusOK {
		return false, nil
	}

	return true, nil
}

func checkTheStatusOfWebsites() {
	for _, website := range websitesList.Websites {
		websiteVsStatusMap[website] = "DOWN"
	}

	for _, website := range websitesList.Websites {
		go func() {
			statusOK, err := websitesList.Check(website)

			if !statusOK {
				if err != nil {
					fmt.Printf("Error occurred - %v\n", err)
				}
				return
			}

			fmt.Printf("Status ok for %s\n", website)

			websiteVsStatusMap[website] = "UP"
		}()
	}
}

var ticker *time.Ticker

func checkTheStatusForEveryOneMinute() {
	if ticker != nil {
		ticker.Stop()
	}

	ticker = time.NewTicker(15 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			checkTheStatusOfWebsites()
		}
	}
}

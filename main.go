package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	articleUrl          string = "https://storage.googleapis.com/aller-structure-task/articles.json"
	contentMarketingUrl string = "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"
)

func main() {
	log.Println("Starting at :8080")

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ciklum-test", mixArticles).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}

/* Main task flow to mix articles with Ad
   prints http response */
func mixArticles(w http.ResponseWriter, r *http.Request) {
	// grab items
	articleItems := getItems(articleUrl)
	marketingItems := getItems(contentMarketingUrl)

	lenArticles := len(articleItems)
	lenMarketing := len(marketingItems)

	// our mixed content slice
	var respContent []interface{}

	var start, end, step int
	// five articles step between ad
	step = 5

	// index for marketing items
	var mIdx int

	// compose content!
	for i := 0; i <= lenArticles; i += step {
		start = i
		end = i + step

		// prevent out of range/nil items
		if end > lenArticles {
			end = lenArticles
		}
		respContent = append(respContent, articleItems[start:end]...)

		// fill marketing items and prevent out of range panic
		if mIdx+1 > lenMarketing {
			// no items any more, insert default "Ad"
			ad := make(map[string]interface{})
			ad["type"] = "Ad"
			respContent = append(respContent, ad)
		} else {
			// we have items, fill with given marketing content
			respContent = append(respContent, marketingItems[mIdx])
			mIdx++
		}
	}

	w.Header().Set("Content-Type", "application/json")
	// print json response
	json.NewEncoder(w).Encode(respContent)
}

// Get items from given url and return slice of items with interface type
func getItems(url string) (items []interface{}) {
	resp, err := http.Get(url)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	items = result["response"].(map[string]interface{})["items"].([]interface{})
	return
}

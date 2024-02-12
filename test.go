package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	requestChannel := make(chan map[string]interface{})
	go worker(requestChannel)
	sendJSONRequest(requestChannel)
	startServer(requestChannel)
}

func startServer(requestChannel chan<- map[string]interface{}) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var req map[string]interface{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		requestChannel <- req

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Received the request successfully\n")
	})

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func worker(requestChannel <-chan map[string]interface{}) {
	url := "https://webhook.site/ab5b4c71-9481-4e7c-9565-7ac0975c944f"

	for req := range requestChannel {
		convertedData := make(map[string]interface{})

		// Extract common fields
		convertedData["event"] = req["ev"]
		convertedData["event_type"] = req["et"]
		convertedData["app_id"] = req["id"]
		convertedData["user_id"] = req["uid"]
		convertedData["message_id"] = req["mid"]
		convertedData["page_title"] = req["t"]
		convertedData["page_url"] = req["p"]
		convertedData["browser_language"] = req["l"]
		convertedData["screen_size"] = req["sc"]

		// Extract attributes
		attributes := make(map[string]map[string]interface{})
		// Extract traits
		traits := make(map[string]map[string]interface{})
		for key := range req {
			if strings.HasPrefix(key, "atrk") {
				attrNum := strings.TrimPrefix(key, "atrk")
				attrKey := req["atrk"+attrNum].(string)
				attrValue := req["atrv"+attrNum]
				attrType := req["atrt"+attrNum].(string)

				attrMap := map[string]interface{}{
					"value": attrValue,
					"type":  attrType,
				}

				attributes[attrKey] = attrMap
			} else if strings.HasPrefix(key, "uatrk") {
				traitNum := strings.TrimPrefix(key, "uatrk")
				traitKey := req["uatrk"+traitNum].(string)
				traitValue := req["uatrv"+traitNum]
				traitType := req["uatrt"+traitNum].(string)

				traitMap := map[string]interface{}{
					"value": traitValue,
					"type":  traitType,
				}

				traits[traitKey] = traitMap
			}
		}
		convertedData["attributes"] = attributes
		convertedData["traits"] = traits

		// Convert to JSON
		jsonBytes, err := json.Marshal(convertedData)
		if err != nil {
			fmt.Println("Error encoding JSON:", err)
			continue
		}

		// Send POST request
		resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonBytes))
		if err != nil {
			fmt.Println("Error sending request:", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Println("Response Status:", resp.Status)
	}
}

// for multiple json request
func sendJSONRequest(requestChannel chan<- map[string]interface{}) {
	jsonDatas := []map[string]interface{}{
		{
			"ev":     "contact_form_submitted",
			"et":     "form_submit",
			"id":     "cl_app_id_001",
			"uid":    "cl_app_id_001-uid-001",
			"mid":    "cl_app_id_001-uid-001",
			"t":      "Vegefoods - Free Bootstrap 4 Template by Colorlib",
			"p":      "http://shielded-eyrie-45679.herokuapp.com/contact-us",
			"l":      "en-US",
			"sc":     "1920 x 1080",
			"atrk1":  "form_varient",
			"atrv1":  "red_top",
			"atrt1":  "string",
			"atrk2":  "ref",
			"atrv2":  "XPOWJRICW993LKJD",
			"atrt2":  "string",
			"uatrk1": "name",
			"uatrv1": "iron man",
			"uatrt1": "string",
			"uatrk2": "email",
			"uatrv2": "ironman@avengers.com",
			"uatrt2": "string",
			"uatrk3": "age",
			"uatrv3": "32",
			"uatrt3": "integer",
		},
		{
			"ev":     "top_cta_clicked",
			"et":     "clicked",
			"id":     "cl_app_id_001",
			"uid":    "cl_app_id_001-uid-001",
			"mid":    "cl_app_id_001-uid-001",
			"t":      "Vegefoods - Free Bootstrap 4 Template by Colorlib",
			"p":      "http://shielded-eyrie-45679.herokuapp.com/contact-us",
			"l":      "en-US",
			"sc":     "1920 x 1080",
			"atrk1":  "button_text",
			"atrv1":  "Free trial",
			"atrt1":  "string",
			"atrk2":  "color_variation",
			"atrv2":  "ESK0023",
			"atrt2":  "string",
			"uatrk1": "user_score",
			"uatrv1": "1034",
			"uatrt1": "number",
			"uatrk2": "gender",
			"uatrv2": "m",
			"uatrt2": "string",
			"uatrk3": "tracking_code",
			"uatrv3": "POSERK093",
			"uatrt3": "string",
		},
		{
			"ev":     "top_cta_clicked",
			"et":     "clicked",
			"id":     "cl_app_id_001",
			"uid":    "cl_app_id_001-uid-001",
			"mid":    "cl_app_id_001-uid-001",
			"t":      "Vegefoods - Free Bootstrap 4 Template by Colorlib",
			"p":      "http://shielded-eyrie-45679.herokuapp.com/contact-us",
			"l":      "en-US",
			"sc":     "1920 x 1080",
			"atrk1":  "button_text",
			"atrv1":  "Free trial",
			"atrt1":  "string",
			"atrk2":  "color_variation",
			"atrv2":  "ESK0023",
			"atrt2":  "string",
			"atrk3":  "page_path",
			"atrv3":  "/blog/category_one/blog_name.html",
			"atrt3":  "string",
			"atrk4":  "source",
			"atrv4":  "facebook",
			"atrt4":  "string",
			"uatrk1": "user_score",
			"uatrv1": "1034",
			"uatrt1": "number",
			"uatrk2": "gender",
			"uatrv2": "m",
			"uatrt2": "string",
			"uatrk3": "tracking_code",
			"uatrv3": "POSERK093",
			"uatrt3": "string",
			"uatrk4": "phone",
			"uatrv4": "9034432423",
			"uatrt4": "number",
			"uatrk5": "coupon_clicked",
			"uatrv5": "true",
			"uatrt5": "boolean",
			"uatrk6": "opt_out",
			"uatrv6": "false",
			"uatrt6": "boolean",
		},
	}
	for _, jsonData := range jsonDatas {
		requestChannel <- jsonData
	}

}

// for single json request

// func sendJSONRequest(requestChannel chan<- map[string]interface{}) {
// 	jsonData := map[string]interface{}{
// 		"ev":     "contact_form_submitted",
// 		"et":     "form_submit",
// 		"id":     "cl_app_id_001",
// 		"uid":    "cl_app_id_001-uid-001",
// 		"mid":    "cl_app_id_001-uid-001",
// 		"t":      "Vegefoods - Free Bootstrap 4 Template by Colorlib",
// 		"p":      "http://shielded-eyrie-45679.herokuapp.com/contact-us",
// 		"l":      "en-US",
// 		"sc":     "1920 x 1080",
// 		"atrk1":  "form_varient",
// 		"atrv1":  "red_top",
// 		"atrt1":  "string",
// 		"atrk2":  "ref",
// 		"atrv2":  "XPOWJRICW993LKJD",
// 		"atrt2":  "string",
// 		"uatrk1": "name",
// 		"uatrv1": "iron man",
// 		"uatrt1": "string",
// 		"uatrk2": "email",
// 		"uatrv2": "ironman@avengers.com",
// 		"uatrt2": "string",
// 		"uatrk3": "age",
// 		"uatrv3": 32,
// 		"uatrt3": "integer",
// 	}

// 	requestChannel <- jsonData
// }

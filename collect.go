package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func Collect() {
	id, secret := LoadUser("User1")
	for {
		collectData(id, secret)
		time.Sleep(60 * time.Second)
	}

}

// The oauth token for NIBE data
type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
}

// The data is queried in three phases. The first phase is asking which systems
// are available for the user: /api/v1/systems
// Systems contains all the systems the user has.
type Systems struct {
	Objects []Object `json:"objects"`
}

// Object contains a single system that the user has.
// The interesting parts of the response:
// {
//     "objects": [
//         {
//             "systemId": 1234567
//         }
//     ]
// }
type Object struct {
	SystemId uint `json:"systemId"`
}

// Then, for all the systems, we get the categories:
// /api/v1/systems/<systemID>/serviceinfo/categories
// [
//     {
//         "categoryId": "STATUS",
//         "name": "status",
//         "parameters": null
//     }
// ]
type Category struct {
	CategoryID string `json:"categoryId"`
}

// Then, for each category, we get the parameters:
// /api/v1/systems/<systemID>/serviceinfo/categories/<categoryID>
// [
//     {
//         "parameterId": 40067,
//         "name": "40067",
//         "title": "avg. outdoor temp",
//         "designation": "BT1",
//         "unit": "°C",
//         "displayValue": "-6.1°C",
//         "rawValue": -61
//     },
// ]
type Parameter struct {
	Title    string `json:"title"`
	RawValue int    `json:"rawValue"`
}

func collectData(id, secret string) {
	data, err := ioutil.ReadFile("./token.json")
	if err != nil {
		log.Println(err)
		return
	}

	var token Token
	err = json.Unmarshal(data, &token)
	if err != nil {
		log.Println("error:", err)
		return
	}

	oauth := &Oauth{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ID:           id,
		Secret:       secret,
	}
	objs := getSystems(oauth)
	for _, obj := range objs {
		cats := getCategories(oauth, obj.SystemId)
		for _, cat := range cats {
			params := getParams(oauth, obj.SystemId, cat.CategoryID)
			fmt.Printf("Received %d params\n", len(params))
			for _, param := range params {
				saveParam(param)
			}
		}
	}
}

// getSystems returns the systems available for this users
func getSystems(oauth *Oauth) []Object {
	url := nibeURL + "/api/v1/systems"
	systems := Systems{}
	if err := getJSON(oauth, url, &systems); err != nil {
		log.Println(err)
		return nil
	}
	return systems.Objects
}

// getCategories returns the available categories
func getCategories(oauth *Oauth, systemID uint) []Category {
	url := fmt.Sprintf(
		"%s/api/v1/systems/%d/serviceinfo/categories",
		nibeURL, systemID)
	categories := []Category{}
	if err := getJSON(oauth, url, &categories); err != nil {
		log.Println(err)
		return nil
	}
	return categories
}

func getParams(oauth *Oauth, systemID uint, catID string) []Parameter {
	url := fmt.Sprintf(
		"%s/api/v1/systems/%d/serviceinfo/categories/%s",
		nibeURL, systemID, catID)
	params := []Parameter{}
	if err := getJSON(oauth, url, &params); err != nil {
		log.Println(err)
		return nil
	}
	return params
}

// getJSON makes a REST request and parses the JSON response to the provided target
func getJSON(oauth *Oauth, url string, target interface{}) error {
	var resp *http.Response
	var err error
	if resp, err = oauth.Get(url); err != nil {
		return err
	}
	if resp.StatusCode == http.StatusUnauthorized {
		body, err := oauth.Refresh(nibeURL + "/oauth/token")
		if err != nil {
			return err
		}
		SaveToken("User1", body)

		if resp, err = oauth.Get(url); err != nil {
			return err
		}
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		err = json.Unmarshal(body, target)
		if err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("%s returned %d", url, resp.StatusCode)
}

func saveParam(param Parameter) {
	trySaveParam(param, "evaporator", "temperature", .1)
	trySaveParam(param, "supply air", "temperature", .1)
	trySaveParam(param, "extract air", "temperature", .1)
	trySaveParam(param, "fan speed", "percent", 1)
	trySaveParam(param, "exhaust air", "temperature", .1)
	trySaveParam(param, "calculated flow temp.", "temperature", .1)
	trySaveParam(param, "heat medium flow", "temperature", .1)
	trySaveParam(param, "return temp.", "temperature", .1)
	trySaveParam(param, "imm. heater sensor", "temperature", .1)
	trySaveParam(param, "hot water charging", "temperature", .1)
	trySaveParam(param, "hot water top", "temperature", .1)
	trySaveParam(param, "outdoor temp.", "temperature", .1)
	trySaveParam(param, "compressor starts", "count", 1)
	trySaveParam(param, "compressor operating time", "hours", 1)
	trySaveParam(param, "compressor operating time hot water", "hours", 1)
	trySaveParam(param, "electrical addition power", "power", .01)
	trySaveParam(param, "time factor", "hours", .1)
}

func trySaveParam(param Parameter, name, fieldKey string, scale float64) {
	if param.Title == name {
		name := strings.ReplaceAll(param.Title, " ", "-")
		name = strings.ReplaceAll(name, ".", "")
		SaveValue("F470", "param", name, fieldKey, param.RawValue, scale)
	}

}

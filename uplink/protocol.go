package uplink

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mlehikoi/nibe/internal/constants"
	"github.com/mlehikoi/nibe/internal/utils"
)

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
	SystemId         uint      `json:"systemId"`
	Name             string    `json:"name"`
	ProductName      string    `json:"productName"`
	SecurityLevel    string    `json:"securityLevel"`
	SerialNumber     string    `json:"serialNumber"`
	LastActivityDate time.Time `json:"lastActivityDate"`
	ConnectionStatus string    `json:"connectionStatus"`
	HasAlarmed       bool      `json:"hasAlarmed"`
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
	Name       string `json:"name"`
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
	ParameterId  int    `json:"parameterId"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	Designation  string `json:"designation"`
	Unit         string `json:"unit"`
	DisplayValue string `json:"displayValue"`
	RawValue     int    `json:"rawValue"`
}

// getSystems returns the systems available for this users
func getSystems(oauth *utils.Oauth) []Object {
	url := constants.NibeURL + "/api/v1/systems"
	systems := Systems{}
	if err := getJSON(oauth, url, &systems); err != nil {
		log.Println(err)
		return nil
	}
	return systems.Objects
}

// getCategories returns the available categories
func getCategories(oauth *utils.Oauth, systemID uint) []Category {
	url := fmt.Sprintf(
		"%s/api/v1/systems/%d/serviceinfo/categories",
		constants.NibeURL, systemID)
	categories := []Category{}
	if err := getJSON(oauth, url, &categories); err != nil {
		log.Println(err)
		return nil
	}
	return categories
}

func getParams(oauth *utils.Oauth, systemID uint, catID string) []Parameter {
	url := fmt.Sprintf(
		"%s/api/v1/systems/%d/serviceinfo/categories/%s",
		constants.NibeURL, systemID, catID)
	params := []Parameter{}
	if err := getJSON(oauth, url, &params); err != nil {
		log.Println(err)
		return nil
	}
	return params
}

/// getJSON makes a REST request and parses the JSON response to the provided target
func getJSON(oauth *utils.Oauth, url string, target interface{}) error {
	resp, err := oauth.Get(url)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusUnauthorized {
		resp.Body.Close()

		fmt.Println("Refresh token ............... ")
		token, err := oauth.Refresh(constants.NibeURL + "/oauth/token")
		if err != nil {
			return err
		}
		if err := token.Dump("nibe-uplink-token.json"); err != nil {
			return err
		}
		// After refreshing the token, get the contents again
		resp, err = oauth.Get(url)
		if err != nil {
			return err
		}
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))
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

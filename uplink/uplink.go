package uplink

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/mlehikoi/nibe/internal/constants"
	"github.com/mlehikoi/nibe/internal/utils"
)

type Uplink struct {
	id      string
	secret  string
	Systems []System
}

func NewUplink(id string, secret string) *Uplink {
	return &Uplink{
		id:     id,
		secret: secret,
	}
}

func (u *Uplink) Update() {
	u.collectData(u.id, u.secret)
}

func (u *Uplink) updateSystems(system System) {
	// Using range to iterate over the slice
	for _, s := range u.Systems {
		if s.systemId == system.systemId {
			s = system
			return
		}
	}
	u.Systems = append(u.Systems, system)
}

func (u *Uplink) collectData(id, secret string) {
	data, err := ioutil.ReadFile(constants.TokenFilename)
	if err != nil {
		log.Println(err)
		return
	}

	var token utils.Token
	err = json.Unmarshal(data, &token)
	if err != nil {
		log.Println("error:", err)
		return
	}

	oauth := &utils.Oauth{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ID:           id,
		Secret:       secret,
	}
	// fmt.Println("access: ", oauth.AccessToken)
	// fmt.Println("refresh: ", oauth.RefreshToken)
	// fmt.Println("id: ", oauth.ID)
	// fmt.Println(oauth.Secret)
	objs := getSystems(oauth)
	var status Status
	for _, obj := range objs {
		sys := newSystem(obj)

		cats := getCategories(oauth, obj.SystemId)
		for _, cat := range cats {
			params := getParams(oauth, obj.SystemId, cat.CategoryID)
			switch cat.CategoryID {
			case "STATUS": // status
				sys.Status = newStatus(params)
			case "CPR_INFO_EP14": //  compressor module
				sys.Compressor = newCompressor(params)
			case "VENTILATION": //  ventilation
				sys.Ventilation = newVentilation(params)
			case "SYSTEM_1": //  climate system 1
				sys.Climate = newClimate(params)
			case "ADDITION": //  addition
				sys.Addition = newAddition(params)
			case "AUX_IN_OUT": //  soft in/outputs
				// AUX1:0
				// AUX2:0
				// AUX3:0
				// AUX4:0
				// AUX5:8
				// X7:17
			case "SYSTEM_INFO": //  info
				// country:15
				// product:-32768
				// serial number:-32768
				// version:-32768
			}
		}
		u.updateSystems(sys)
	}
	fmt.Println(status.OutdoorTemp)
}

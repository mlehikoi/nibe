// Contains functions for saving measurements to InfluxDB
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

type InfluxSettings struct {
	Host         string `json:"host"`
	Token        string `json:"token"`
	Organization string `json:"organization"`
	Bucket       string `json:"bucket"`
}

type Settings struct {
	Influx InfluxSettings `json:"influx"`
}

var settings Settings

func init() {
	data, err := ioutil.ReadFile("./settings.json")
	if err != nil {
		panic(err)
	}

	if err = json.Unmarshal(data, &settings); err != nil {
		panic(err)
	}
}

// SaveValue saves a single value in the database. Organization and bucket are
// specified by the configurations.
func SaveValue(measurement, tagKey, tagVal, fieldKey string, rawValue int, scale float64) {
	client := influxdb2.NewClient(settings.Influx.Host, settings.Influx.Token)
	defer client.Close()
	writeAPI := client.WriteAPIBlocking(settings.Influx.Organization,
		settings.Influx.Bucket)

	value := fmt.Sprintf("%d", rawValue)
	if scale != 1 {
		value = fmt.Sprintf("%f", float64(rawValue)*scale)
	}
	cmd := fmt.Sprintf("%s,%s=%s %s=%s",
		measurement,
		tagKey, tagVal,
		fieldKey, value)
	err := writeAPI.WriteRecord(context.Background(), cmd)
	if err != nil {
		log.Fatal(err)
	}
}

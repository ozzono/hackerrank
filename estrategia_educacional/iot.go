package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strings"
)

func main() {
	fmt.Println(avgRotorSpeed("running", 4))
	// fmt.Println(avgRotorSpeed("malfunctioning", 2))
	fmt.Println(avgRotorSpeed("stopped", 1))
}

func avgRotorSpeed(statusQuery string, parentId int32) int32 {
	fmt.Printf("status: %s pID: %d\n", statusQuery, parentId)
	type asset struct {
		ID    int    `json:"id"`
		Alias string `json:"alias"`
	}

	type parent struct {
		ID    int    `json:"id"`
		Alias string `json:"alias"`
	}

	type operatingParams struct {
		RotorSpeed    int     `json:"rotorSpeed"`
		Slack         float64 `json:"slack"`
		RootThreshold float64 `json:"rootThreshold"`
	}

	type device struct {
		ID              int             `json:"id"`
		Timestamp       int64           `json:"timestamp"`
		Status          string          `json:"status"`
		OperatingParams operatingParams `json:"operatingParams"`
		Asset           asset           `json:"asset"`
		Parent          parent          `json:"parent,omitempty"`
	}

	type deviceData struct {
		Page       string   `json:"page"`
		PerPage    int      `json:"per_page"`
		Total      int      `json:"total"`
		TotalPages int      `json:"total_pages"`
		Devices    []device `json:"data"`
	}

	request := func() ([]device, error) {
		request := func(statusQuery string, page int) (deviceData, error) {
			url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/iot_devices/search?status=%s&page=%d", strings.ToUpper(statusQuery), page)
			fmt.Printf("url %s\n", url)

			req, err := http.NewRequest("GET", url, nil)
			if err != nil {
				return deviceData{}, fmt.Errorf("http.NewRequest err: %v", err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				return deviceData{}, fmt.Errorf("http.DefaultClient.Do err: %v", err)
			}

			defer res.Body.Close()
			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				return deviceData{}, fmt.Errorf("ioutil.ReadAll err: %v", err)
			}
			output := deviceData{}
			err = json.Unmarshal(body, &output)
			if err != nil {
				return deviceData{}, fmt.Errorf("json.Unmarshal err :%v\nbody: %s", err, string(body))
			}
			return output, nil
		}

		r0, err := request(statusQuery, 0)
		if err != nil {
			return []device{}, err
		}

		devices := map[int]device{}
		for i := 1; i <= r0.TotalPages; i++ {
			r, err := request(statusQuery, i)
			if err != nil {
				return []device{}, err
			}
			for i := range r.Devices {
				devices[r.Devices[i].ID] = r.Devices[i]
			}
		}

		output := []device{}
		for i := range devices {
			output = append(output, devices[i])
		}

		return output, nil
	}

	devices, err := request()
	if err != nil {
		log.Fatal(err)
	}

	filter := func(devices []device, parentID int) []device {
		output := []device{}
		for i := range devices {
			if devices[i].Parent.ID == parentID {
				output = append(output, devices[i])
			}
		}
		return output
	}
	list := filter(devices, int(parentId))

	var sum float64
	for i := range list {
		sum += float64(list[i].OperatingParams.RotorSpeed)
	}
	if len(list) == 0 {
		return 0
	}
	return int32(math.Floor(sum / float64(len(list))))
}

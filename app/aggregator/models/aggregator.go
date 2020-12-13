package aggregatormodels

import (
	// "errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"errors"
	"encoding/json"
	// "encoding/gob"
    // "bytes"
	// "requestaggregator/config"
)

const aggregatorCollection = "aggregator"

type CountryWiseData struct {
	WebRequest uint `json:"webreq"`
	TimeSpent uint `json:"timespent"`
	DeviceData []map[string]map[string]uint `json:"deviceData"`
}

type RequestTree struct {
	Data map[string]CountryWiseData `json:"data"`
}

type Aggregator struct {
	Dimension []DimensionKeyValue `json:"dim"`
	Metrics []MetricsKeyValue `json:"metrics"`
}

type DimensionKeyValue struct {
	Key string `json:"key"`
	Value interface{} `json:"val"`
}

type MetricsKeyValue struct {
	Key string `json:"key"`
	Value interface{} `json:"val"`
}

func InsertData(aggregatorInsertData *Aggregator) (response RequestTree, err error) {

	filePath, _ := filepath.Abs("./app/data/data.json")
	fileContent, readErr := ioutil.ReadFile(filePath)

	var data RequestTree
	if readErr == nil {
		jsonErr := json.Unmarshal(fileContent, &data)
		if jsonErr != nil {
			fmt.Println("Error while reading json: ", jsonErr)
		}
	} else {
		fmt.Println("Error while reading file: ", readErr)
	}

	countryFilter := ""
	deviceFilter := ""
	for _, dt := range aggregatorInsertData.Dimension {
		if dt.Key == "country" {
			countryFilter = fmt.Sprint(dt.Value)
		}
		if dt.Key == "device" {
			deviceFilter = fmt.Sprint(dt.Value)
		}
	}

	var webreqI, timespentI float64
	for _, dt := range aggregatorInsertData.Metrics {
		if dt.Key == "webreq" {
			webreqI, _ = dt.Value.(float64)
		}
		if dt.Key == "timespent" {
			timespentI, _ = dt.Value.(float64)
		}
	}
	webreq := uint(webreqI)
	timespent := uint(timespentI)

	if _, ok := data.Data[countryFilter]; ok {
		countryData := data.Data[countryFilter]
		countryData.WebRequest = countryData.WebRequest + webreq
		countryData.TimeSpent = countryData.TimeSpent + timespent

		deviceExists := false
		var deviceIndex int
		for i, dt := range countryData.DeviceData {
			if _, ok:= dt[deviceFilter]; ok {
				deviceExists = true
				deviceIndex = i
			}
		}

		if deviceExists {
			countryData.DeviceData[deviceIndex][deviceFilter]["webreq"] = countryData.DeviceData[deviceIndex][deviceFilter]["webreq"] + webreq
			countryData.DeviceData[deviceIndex][deviceFilter]["timespent"] = countryData.DeviceData[deviceIndex][deviceFilter]["timespent"] + timespent
		} else {
			metricWiseData := make(map[string]map[string]uint)
			metricWiseData[deviceFilter] = make(map[string]uint)
			metricWiseData[deviceFilter]["webreq"] = webreq
			metricWiseData[deviceFilter]["timespent"] = timespent
			countryData.DeviceData = append(countryData.DeviceData, metricWiseData)
		}
		data.Data[countryFilter] = countryData
	} else {
		metricWiseDataList := make([]map[string]map[string]uint, 0)
		metricWiseData := make(map[string]map[string]uint)
		metricWiseData[deviceFilter] = make(map[string]uint)
		metricWiseData[deviceFilter]["webreq"] = webreq
		metricWiseData[deviceFilter]["timespent"] = timespent
		metricWiseDataList = append(metricWiseDataList, metricWiseData)

		countryWiseData := CountryWiseData{
			WebRequest: webreq,
			TimeSpent: timespent,
			DeviceData: metricWiseDataList,
		}
		if len(data.Data) > 0 {
			data.Data[countryFilter] = countryWiseData
		} else {
			cData := make(map[string]CountryWiseData)
			cData[countryFilter] = countryWiseData
			data.Data = cData
		}
	}

	bytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error while marshalling file data:", err)
		return RequestTree{}, err
	}
	_err := ioutil.WriteFile(filePath, bytes, 0644)
	if _err != nil {
		fmt.Println("Error while writing data to json")
		return RequestTree{}, _err
	}

	return data, nil
}

func GetData(aggregatorGetData *Aggregator) (response Aggregator, err error) {
	fmt.Println("aggregatorInsertData: ", aggregatorGetData)
	validParams := false
	countryFilter := ""
	deviceFilter := ""
	for _, dt := range aggregatorGetData.Dimension {
		if dt.Key == "country" {
			validParams = true
			countryFilter = fmt.Sprint(dt.Value)
		}
		if dt.Key == "device" {
			deviceFilter = fmt.Sprint(dt.Value)
		}
	}

	if validParams {
		filePath, _ := filepath.Abs("./app/data/data.json")
		fileContent, readErr := ioutil.ReadFile(filePath)
		var data RequestTree
		if readErr == nil {
			jsonErr := json.Unmarshal(fileContent, &data)
			if jsonErr != nil {
				fmt.Println("Error while reading json: ", jsonErr)
			}
		} else {
			fmt.Println("Error while reading file: ", readErr)
		}

		countryData := data.Data[countryFilter]

		if len(countryFilter)>0 && len(deviceFilter)>0 {
			webreq := uint(0)
			timespent := uint(0)
			fmt.Println("deviceFilter: ", deviceFilter)
			for _, dt := range countryData.DeviceData {
				fmt.Println("dt[deviceFilter]: ", dt)
				if _, ok := dt[deviceFilter]; ok {
					fmt.Println("dt[deviceFilter]: ", dt)
					webreq = dt[deviceFilter]["webreq"]
					timespent = dt[deviceFilter]["timespent"]
				}
			}

			fmt.Println("data: ", webreq, timespent)

			dim := make([]DimensionKeyValue, 0)
			dim = append(dim, DimensionKeyValue{
				Key: "country",
				Value: countryFilter,
			})
			dim = append(dim, DimensionKeyValue{
				Key: "device",
				Value: deviceFilter,
			})
			metrics := make([]MetricsKeyValue, 0)
			metrics = append(metrics, MetricsKeyValue{
				Key: "webreq",
				Value: webreq,
			})
			metrics = append(metrics, MetricsKeyValue{
				Key: "timespent",
				Value: timespent,
			})
			response = Aggregator{
				Dimension: dim,
				Metrics: metrics,
			}
		} else if len(countryFilter)>0 {
			dim := make([]DimensionKeyValue, 0)
			dim = append(dim, DimensionKeyValue{
				Key: "country",
				Value: countryFilter,
			})
			metrics := make([]MetricsKeyValue, 0)
			metrics = append(metrics, MetricsKeyValue{
				Key: "webreq",
				Value: countryData.WebRequest,
			})
			metrics = append(metrics, MetricsKeyValue{
				Key: "timespent",
				Value: countryData.TimeSpent,
			})
			response = Aggregator{
				Dimension: dim,
				Metrics: metrics,
			}
		}
		return response, nil
	}
	err = errors.New("Invalid Params")
	return
}

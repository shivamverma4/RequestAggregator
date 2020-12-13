package aggregator

import (
	"requestaggregator/app/utils"
	"fmt"
	"net/http"
	aggregatormodels "requestaggregator/app/aggregator/models"

	"github.com/labstack/echo"
)

func HandleGetAggregatorData(c echo.Context) (err error) {

	aggregatorData := new(aggregatormodels.Aggregator)
	if err = c.Bind(aggregatorData); err != nil {
		return
	}

	response, err := GetAggregationData(aggregatorData)
	if err != nil {
		fmt.Println("Error: ", err)
		resp := utils.CustomHTTPResponse{}
		resp.Data = 0
		resp.Message = "Incomplete params"
		return c.JSON(http.StatusBadRequest, resp)	
	}

	return c.JSON(http.StatusOK, response)
}

func HandleInsertAggregatorData(c echo.Context) (err error) {

	aggregatorData := new(aggregatormodels.Aggregator)
	if err = c.Bind(aggregatorData); err != nil {
		return
	}

	validParams := false
	if (aggregatorData.Dimension[0].Key == "country" && aggregatorData.Dimension[1].Key == "device") || (aggregatorData.Dimension[1].Key == "country" && aggregatorData.Dimension[0].Key == "device") {
		validParams = true
	}

	resp := utils.CustomHTTPResponse{}
	if validParams {
		response, err := InsertAggregationData(aggregatorData)
		if err != nil {
			resp.Data = 0
			resp.Message = "Data can not be inserted"
			return c.JSON(http.StatusBadRequest, resp)
		}
		fmt.Println("response: ", response)
		resp.Data = response
		resp.Message = "Data Inserted"
		return c.JSON(http.StatusOK, resp)
	}

	resp.Data = 0
	resp.Message = "Data can not be inserted"
	return c.JSON(http.StatusBadRequest, resp)
}

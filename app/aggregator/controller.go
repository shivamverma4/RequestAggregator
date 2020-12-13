package aggregator

import (
	aggregatormodels "requestaggregator/app/aggregator/models"
)

func GetAggregationData(aggregatorData *aggregatormodels.Aggregator) (aggregatormodels.Aggregator, error) {
	resp, err := aggregatormodels.GetData(aggregatorData)
	if err != nil {
		return aggregatormodels.Aggregator{}, err
	}
	return resp, nil
}

func InsertAggregationData(aggregatorData *aggregatormodels.Aggregator) (aggregatormodels.RequestTree, error) {
	resp, err := aggregatormodels.InsertData(aggregatorData)
	if err != nil {
		return aggregatormodels.RequestTree{}, err
	}
	return resp, nil
}

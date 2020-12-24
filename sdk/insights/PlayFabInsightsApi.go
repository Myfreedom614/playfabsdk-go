package insights

// This code was generated by a tool. Any changes may be overwritten

import (
	"encoding/json"

	playfab "github.com/Myfreedom614/playfabsdk-go/sdk"

	"github.com/mitchellh/mapstructure"
)

// GetDetails gets the current values for the Insights performance and data storage retention, list of pending operations, and the
// performance and data storage retention limits.
// https://api.playfab.com/Documentation/Insights/method/GetDetails
func GetDetails(settings *playfab.Settings, postData *InsightsEmptyRequestModel, entityToken string) (*InsightsGetDetailsResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Insights/GetDetails", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &InsightsGetDetailsResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// GetLimits retrieves the range of allowed values for performance and data storage retention values as well as the submeter details
// for each performance level.
// https://api.playfab.com/Documentation/Insights/method/GetLimits
func GetLimits(settings *playfab.Settings, postData *InsightsEmptyRequestModel, entityToken string) (*InsightsGetLimitsResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Insights/GetLimits", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &InsightsGetLimitsResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// GetOperationStatus gets the status of a SetPerformance or SetStorageRetention operation.
// https://api.playfab.com/Documentation/Insights/method/GetOperationStatus
func GetOperationStatus(settings *playfab.Settings, postData *InsightsGetOperationStatusRequestModel, entityToken string) (*InsightsGetOperationStatusResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Insights/GetOperationStatus", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &InsightsGetOperationStatusResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// GetPendingOperations gets a list of pending SetPerformance and/or SetStorageRetention operations for the title.
// https://api.playfab.com/Documentation/Insights/method/GetPendingOperations
func GetPendingOperations(settings *playfab.Settings, postData *InsightsGetPendingOperationsRequestModel, entityToken string) (*InsightsGetPendingOperationsResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Insights/GetPendingOperations", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &InsightsGetPendingOperationsResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// SetPerformance sets the Insights performance level value for the title.
// https://api.playfab.com/Documentation/Insights/method/SetPerformance
func SetPerformance(settings *playfab.Settings, postData *InsightsSetPerformanceRequestModel, entityToken string) (*InsightsOperationResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Insights/SetPerformance", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &InsightsOperationResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

// SetStorageRetention sets the Insights data storage retention days value for the title.
// https://api.playfab.com/Documentation/Insights/method/SetStorageRetention
func SetStorageRetention(settings *playfab.Settings, postData *InsightsSetStorageRetentionRequestModel, entityToken string) (*InsightsOperationResponseModel, error) {
	if entityToken == "" {
		return nil, playfab.NewCustomError("entityToken should not be an empty string", playfab.ErrorGeneric)
	}
	b, errMarshal := json.Marshal(postData)
	if errMarshal != nil {
		return nil, playfab.NewCustomError(errMarshal.Error(), playfab.ErrorMarshal)
	}

	sourceMap, err := playfab.Request(settings, b, "/Insights/SetStorageRetention", "X-EntityToken", entityToken)
	if err != nil {
		return nil, err
	}

	result := &InsightsOperationResponseModel{}

	config := mapstructure.DecoderConfig{
		DecodeHook: playfab.StringToDateTimeHook,
		Result:     result,
	}

	decoder, errDecoding := mapstructure.NewDecoder(&config)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	errDecoding = decoder.Decode(sourceMap)
	if errDecoding != nil {
		return nil, playfab.NewCustomError(errDecoding.Error(), playfab.ErrorDecoding)
	}

	return result, nil
}

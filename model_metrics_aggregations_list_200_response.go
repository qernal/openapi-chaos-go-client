/*
Chaos

Central Management API - publicly exposed set of APIs for managing deployments

API version: 1.0.0
Contact: support@qernal.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_chaos_client

import (
	"encoding/json"
	"gopkg.in/validator.v2"
	"fmt"
)

// MetricsAggregationsList200Response - struct for MetricsAggregationsList200Response
type MetricsAggregationsList200Response struct {
	MetricHttpAggregation *MetricHttpAggregation
	MetricResourceAggregation *MetricResourceAggregation
}

// MetricHttpAggregationAsMetricsAggregationsList200Response is a convenience function that returns MetricHttpAggregation wrapped in MetricsAggregationsList200Response
func MetricHttpAggregationAsMetricsAggregationsList200Response(v *MetricHttpAggregation) MetricsAggregationsList200Response {
	return MetricsAggregationsList200Response{
		MetricHttpAggregation: v,
	}
}

// MetricResourceAggregationAsMetricsAggregationsList200Response is a convenience function that returns MetricResourceAggregation wrapped in MetricsAggregationsList200Response
func MetricResourceAggregationAsMetricsAggregationsList200Response(v *MetricResourceAggregation) MetricsAggregationsList200Response {
	return MetricsAggregationsList200Response{
		MetricResourceAggregation: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *MetricsAggregationsList200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into MetricHttpAggregation
	err = newStrictDecoder(data).Decode(&dst.MetricHttpAggregation)
	if err == nil {
		jsonMetricHttpAggregation, _ := json.Marshal(dst.MetricHttpAggregation)
		if string(jsonMetricHttpAggregation) == "{}" { // empty struct
			dst.MetricHttpAggregation = nil
		} else {
			if err = validator.Validate(dst.MetricHttpAggregation); err != nil {
				dst.MetricHttpAggregation = nil
			} else {
				match++
			}
		}
	} else {
		dst.MetricHttpAggregation = nil
	}

	// try to unmarshal data into MetricResourceAggregation
	err = newStrictDecoder(data).Decode(&dst.MetricResourceAggregation)
	if err == nil {
		jsonMetricResourceAggregation, _ := json.Marshal(dst.MetricResourceAggregation)
		if string(jsonMetricResourceAggregation) == "{}" { // empty struct
			dst.MetricResourceAggregation = nil
		} else {
			if err = validator.Validate(dst.MetricResourceAggregation); err != nil {
				dst.MetricResourceAggregation = nil
			} else {
				match++
			}
		}
	} else {
		dst.MetricResourceAggregation = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.MetricHttpAggregation = nil
		dst.MetricResourceAggregation = nil

		return fmt.Errorf("data matches more than one schema in oneOf(MetricsAggregationsList200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(MetricsAggregationsList200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src MetricsAggregationsList200Response) MarshalJSON() ([]byte, error) {
	if src.MetricHttpAggregation != nil {
		return json.Marshal(&src.MetricHttpAggregation)
	}

	if src.MetricResourceAggregation != nil {
		return json.Marshal(&src.MetricResourceAggregation)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *MetricsAggregationsList200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.MetricHttpAggregation != nil {
		return obj.MetricHttpAggregation
	}

	if obj.MetricResourceAggregation != nil {
		return obj.MetricResourceAggregation
	}

	// all schemas are nil
	return nil
}

type NullableMetricsAggregationsList200Response struct {
	value *MetricsAggregationsList200Response
	isSet bool
}

func (v NullableMetricsAggregationsList200Response) Get() *MetricsAggregationsList200Response {
	return v.value
}

func (v *NullableMetricsAggregationsList200Response) Set(val *MetricsAggregationsList200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsAggregationsList200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsAggregationsList200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsAggregationsList200Response(val *MetricsAggregationsList200Response) *NullableMetricsAggregationsList200Response {
	return &NullableMetricsAggregationsList200Response{value: val, isSet: true}
}

func (v NullableMetricsAggregationsList200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsAggregationsList200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}



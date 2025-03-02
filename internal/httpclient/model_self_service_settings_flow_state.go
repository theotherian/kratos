/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: v0.8.1-alpha.1
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// SelfServiceSettingsFlowState show_form: No user data has been collected, or it is invalid, and thus the form should be shown. success: Indicates that the settings flow has been updated successfully with the provided data. Done will stay true when repeatedly checking. If set to true, done will revert back to false only when a flow with invalid (e.g. \"please use a valid phone number\") data was sent.
type SelfServiceSettingsFlowState string

// List of selfServiceSettingsFlowState
const (
	SELFSERVICESETTINGSFLOWSTATE_SHOW_FORM SelfServiceSettingsFlowState = "show_form"
	SELFSERVICESETTINGSFLOWSTATE_SUCCESS   SelfServiceSettingsFlowState = "success"
)

func (v *SelfServiceSettingsFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelfServiceSettingsFlowState(value)
	for _, existing := range []SelfServiceSettingsFlowState{"show_form", "success"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelfServiceSettingsFlowState", value)
}

// Ptr returns reference to selfServiceSettingsFlowState value
func (v SelfServiceSettingsFlowState) Ptr() *SelfServiceSettingsFlowState {
	return &v
}

type NullableSelfServiceSettingsFlowState struct {
	value *SelfServiceSettingsFlowState
	isSet bool
}

func (v NullableSelfServiceSettingsFlowState) Get() *SelfServiceSettingsFlowState {
	return v.value
}

func (v *NullableSelfServiceSettingsFlowState) Set(val *SelfServiceSettingsFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceSettingsFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceSettingsFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceSettingsFlowState(val *SelfServiceSettingsFlowState) *NullableSelfServiceSettingsFlowState {
	return &NullableSelfServiceSettingsFlowState{value: val, isSet: true}
}

func (v NullableSelfServiceSettingsFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceSettingsFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

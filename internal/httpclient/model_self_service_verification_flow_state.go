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

// SelfServiceVerificationFlowState The state represents the state of the verification flow.  choose_method: ask the user to choose a method (e.g. recover account via email) sent_email: the email has been sent to the user passed_challenge: the request was successful and the recovery challenge was passed.
type SelfServiceVerificationFlowState string

// List of selfServiceVerificationFlowState
const (
	SELFSERVICEVERIFICATIONFLOWSTATE_CHOOSE_METHOD    SelfServiceVerificationFlowState = "choose_method"
	SELFSERVICEVERIFICATIONFLOWSTATE_SENT_EMAIL       SelfServiceVerificationFlowState = "sent_email"
	SELFSERVICEVERIFICATIONFLOWSTATE_PASSED_CHALLENGE SelfServiceVerificationFlowState = "passed_challenge"
)

func (v *SelfServiceVerificationFlowState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SelfServiceVerificationFlowState(value)
	for _, existing := range []SelfServiceVerificationFlowState{"choose_method", "sent_email", "passed_challenge"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SelfServiceVerificationFlowState", value)
}

// Ptr returns reference to selfServiceVerificationFlowState value
func (v SelfServiceVerificationFlowState) Ptr() *SelfServiceVerificationFlowState {
	return &v
}

type NullableSelfServiceVerificationFlowState struct {
	value *SelfServiceVerificationFlowState
	isSet bool
}

func (v NullableSelfServiceVerificationFlowState) Get() *SelfServiceVerificationFlowState {
	return v.value
}

func (v *NullableSelfServiceVerificationFlowState) Set(val *SelfServiceVerificationFlowState) {
	v.value = val
	v.isSet = true
}

func (v NullableSelfServiceVerificationFlowState) IsSet() bool {
	return v.isSet
}

func (v *NullableSelfServiceVerificationFlowState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSelfServiceVerificationFlowState(val *SelfServiceVerificationFlowState) *NullableSelfServiceVerificationFlowState {
	return &NullableSelfServiceVerificationFlowState{value: val, isSet: true}
}

func (v NullableSelfServiceVerificationFlowState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSelfServiceVerificationFlowState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

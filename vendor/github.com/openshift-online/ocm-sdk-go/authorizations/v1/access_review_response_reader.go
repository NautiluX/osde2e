/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/authorizations/v1

import (
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// accessReviewResponseData is the data structure used internally to marshal and unmarshal
// objects of type 'access_review_response'.
type accessReviewResponseData struct {
	AccountUsername *string "json:\"account_username,omitempty\""
	Action          *string "json:\"action,omitempty\""
	Allowed         *bool   "json:\"allowed,omitempty\""
	ClusterID       *string "json:\"cluster_id,omitempty\""
	OrganizationID  *string "json:\"organization_id,omitempty\""
	ResourceType    *string "json:\"resource_type,omitempty\""
	SubscriptionID  *string "json:\"subscription_id,omitempty\""
}

// MarshalAccessReviewResponse writes a value of the 'access_review_response' to the given target,
// which can be a writer or a JSON encoder.
func MarshalAccessReviewResponse(object *AccessReviewResponse, target interface{}) error {
	encoder, err := helpers.NewEncoder(target)
	if err != nil {
		return err
	}
	data, err := object.wrap()
	if err != nil {
		return err
	}
	return encoder.Encode(data)
}

// wrap is the method used internally to convert a value of the 'access_review_response'
// value to a JSON document.
func (o *AccessReviewResponse) wrap() (data *accessReviewResponseData, err error) {
	if o == nil {
		return
	}
	data = new(accessReviewResponseData)
	data.AccountUsername = o.accountUsername
	data.Action = o.action
	data.Allowed = o.allowed
	data.ClusterID = o.clusterID
	data.OrganizationID = o.organizationID
	data.ResourceType = o.resourceType
	data.SubscriptionID = o.subscriptionID
	return
}

// UnmarshalAccessReviewResponse reads a value of the 'access_review_response' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalAccessReviewResponse(source interface{}) (object *AccessReviewResponse, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(accessReviewResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'access_review_response' type.
func (d *accessReviewResponseData) unwrap() (object *AccessReviewResponse, err error) {
	if d == nil {
		return
	}
	object = new(AccessReviewResponse)
	object.accountUsername = d.AccountUsername
	object.action = d.Action
	object.allowed = d.Allowed
	object.clusterID = d.ClusterID
	object.organizationID = d.OrganizationID
	object.resourceType = d.ResourceType
	object.subscriptionID = d.SubscriptionID
	return
}

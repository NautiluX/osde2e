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

// ResourceReview represents the values of the 'resource_review' type.
//
// Contains the result of performing a resource access review.
type ResourceReview struct {
	accountUsername *string
	action          *string
	clusterIDs      []string
	organizationIDs []string
	resourceType    *string
	subscriptionIDs []string
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ResourceReview) Empty() bool {
	return o == nil || (o.accountUsername == nil &&
		o.action == nil &&
		len(o.clusterIDs) == 0 &&
		len(o.organizationIDs) == 0 &&
		o.resourceType == nil &&
		len(o.subscriptionIDs) == 0 &&
		true)
}

// AccountUsername returns the value of the 'account_username' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Name of the account that is trying to perform the access.
func (o *ResourceReview) AccountUsername() string {
	if o != nil && o.accountUsername != nil {
		return *o.accountUsername
	}
	return ""
}

// GetAccountUsername returns the value of the 'account_username' attribute and
// a flag indicating if the attribute has a value.
//
// Name of the account that is trying to perform the access.
func (o *ResourceReview) GetAccountUsername() (value string, ok bool) {
	ok = o != nil && o.accountUsername != nil
	if ok {
		value = *o.accountUsername
	}
	return
}

// Action returns the value of the 'action' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Action that will the user is trying to perform.
func (o *ResourceReview) Action() string {
	if o != nil && o.action != nil {
		return *o.action
	}
	return ""
}

// GetAction returns the value of the 'action' attribute and
// a flag indicating if the attribute has a value.
//
// Action that will the user is trying to perform.
func (o *ResourceReview) GetAction() (value string, ok bool) {
	ok = o != nil && o.action != nil
	if ok {
		value = *o.action
	}
	return
}

// ClusterIDs returns the value of the 'cluster_IDs' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Identifiers of the clusters that the user has permission to perform the action upon.
func (o *ResourceReview) ClusterIDs() []string {
	if o == nil {
		return nil
	}
	return o.clusterIDs
}

// GetClusterIDs returns the value of the 'cluster_IDs' attribute and
// a flag indicating if the attribute has a value.
//
// Identifiers of the clusters that the user has permission to perform the action upon.
func (o *ResourceReview) GetClusterIDs() (value []string, ok bool) {
	ok = o != nil && o.clusterIDs != nil
	if ok {
		value = o.clusterIDs
	}
	return
}

// OrganizationIDs returns the value of the 'organization_IDs' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Identifiers of the organizations that the user has permissions to perform the action
// upon.
func (o *ResourceReview) OrganizationIDs() []string {
	if o == nil {
		return nil
	}
	return o.organizationIDs
}

// GetOrganizationIDs returns the value of the 'organization_IDs' attribute and
// a flag indicating if the attribute has a value.
//
// Identifiers of the organizations that the user has permissions to perform the action
// upon.
func (o *ResourceReview) GetOrganizationIDs() (value []string, ok bool) {
	ok = o != nil && o.organizationIDs != nil
	if ok {
		value = o.organizationIDs
	}
	return
}

// ResourceType returns the value of the 'resource_type' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Type of resource.
func (o *ResourceReview) ResourceType() string {
	if o != nil && o.resourceType != nil {
		return *o.resourceType
	}
	return ""
}

// GetResourceType returns the value of the 'resource_type' attribute and
// a flag indicating if the attribute has a value.
//
// Type of resource.
func (o *ResourceReview) GetResourceType() (value string, ok bool) {
	ok = o != nil && o.resourceType != nil
	if ok {
		value = *o.resourceType
	}
	return
}

// SubscriptionIDs returns the value of the 'subscription_IDs' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Identifiers of the subscriptions that the user has permission to perform the action upon.
func (o *ResourceReview) SubscriptionIDs() []string {
	if o == nil {
		return nil
	}
	return o.subscriptionIDs
}

// GetSubscriptionIDs returns the value of the 'subscription_IDs' attribute and
// a flag indicating if the attribute has a value.
//
// Identifiers of the subscriptions that the user has permission to perform the action upon.
func (o *ResourceReview) GetSubscriptionIDs() (value []string, ok bool) {
	ok = o != nil && o.subscriptionIDs != nil
	if ok {
		value = o.subscriptionIDs
	}
	return
}

// ResourceReviewList is a list of values of the 'resource_review' type.
type ResourceReviewList struct {
	items []*ResourceReview
}

// Len returns the length of the list.
func (l *ResourceReviewList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ResourceReviewList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ResourceReviewList) Get(i int) *ResourceReview {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ResourceReviewList) Slice() []*ResourceReview {
	var slice []*ResourceReview
	if l == nil {
		slice = make([]*ResourceReview, 0)
	} else {
		slice = make([]*ResourceReview, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ResourceReviewList) Each(f func(item *ResourceReview) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ResourceReviewList) Range(f func(index int, item *ResourceReview) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}

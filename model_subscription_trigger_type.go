/*
 * Geo API
 *
 * The Geo API provides access to the different features of the Telenor Geo server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package geoclient
// SubscriptionTriggerType the model 'SubscriptionTriggerType'
type SubscriptionTriggerType string

// List of SubscriptionTriggerType
const (
	INSIDE SubscriptionTriggerType = "inside"
	OUTSIDE SubscriptionTriggerType = "outside"
	ENTERED SubscriptionTriggerType = "entered"
	EXITED SubscriptionTriggerType = "exited"
)

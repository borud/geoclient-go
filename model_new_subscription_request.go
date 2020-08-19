/*
 * Geo API
 *
 * The Geo API provides access to the different features of the Telenor Geo server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package geoclient
// NewSubscriptionRequest struct for NewSubscriptionRequest
type NewSubscriptionRequest struct {
	// The ID of the Team for the Subscription
	TeamId int64 `json:"teamId"`
	// The name of the Subscription
	Name string `json:"name,omitempty"`
	// The description of the Subscription
	Description string `json:"description,omitempty"`
	// Boolean showing whether the subscription is active
	Active bool `json:"active,omitempty"`
	Output SubscriptionOutput `json:"output"`
	TriggerCriteria SubscriptionResponseTriggerCriteria `json:"triggerCriteria"`
	// The ID of the ShapeCollection containing shapes to subscribe to
	ShapeCollectionId int64 `json:"shapeCollectionId"`
	Trackable SubscriptionResponseTrackable `json:"trackable"`
}

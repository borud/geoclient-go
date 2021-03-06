/*
 * Geo API
 *
 * The Geo API provides access to the different features of the Telenor Geo server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package geoclient
// SubscriptionResponseTrackable Details of which entities the subscription tracks
type SubscriptionResponseTrackable struct {
	// The ID of the trackable entity
	Id float32 `json:"id"`
	Type SubscriptionTrackerType `json:"type"`
}

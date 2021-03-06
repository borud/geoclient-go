/*
 * Geo API
 *
 * The Geo API provides access to the different features of the Telenor Geo server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package geoclient
// SubscriptionTrackerType the model 'SubscriptionTrackerType'
type SubscriptionTrackerType string

// List of SubscriptionTrackerType
const (
	TRACKER SubscriptionTrackerType = "tracker"
	COLLECTION SubscriptionTrackerType = "collection"
)

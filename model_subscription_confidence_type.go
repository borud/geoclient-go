/*
 * Geo API
 *
 * The Geo API provides access to the different features of the Telenor Geo server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package geoclient
// SubscriptionConfidenceType the model 'SubscriptionConfidenceType'
type SubscriptionConfidenceType string

// List of SubscriptionConfidenceType
const (
	LOW SubscriptionConfidenceType = "low"
	MEDIUM SubscriptionConfidenceType = "medium"
	HIGH SubscriptionConfidenceType = "high"
)

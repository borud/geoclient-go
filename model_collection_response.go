/*
 * Geo API
 *
 * The Geo API provides access to the different features of the Telenor Geo server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package geoclient
// CollectionResponse CollectionResponse is the response you get when interacting with the Collection API
type CollectionResponse struct {
	// The ID of the Collection
	Id int64 `json:"id"`
	// Name of the Collection
	Name string `json:"name"`
	// Description of the Collection
	Description string `json:"description"`
}

/*
 * Geo API
 *
 * The Geo API provides access to the different features of the Telenor Geo server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package geoclient
// TeamResponse TeamResponse is the response you get when interacting with the Team API
type TeamResponse struct {
	// The ID of the Team
	Id int64 `json:"id"`
	// Name of the Team
	Name string `json:"name"`
	// Description of the Team
	Description string `json:"description"`
}

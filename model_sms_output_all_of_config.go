/*
 * Geo API
 *
 * The Geo API provides access to the different features of the Telenor Geo server, including capabilities such as pub/sub, geofencing and storage of tracking data. 
 *
 * API version: v0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package geoclient
// SmsOutputAllOfConfig SMS configuration
type SmsOutputAllOfConfig struct {
	// Twilio API Key to be used when sending SMS
	TwilioApiKey string `json:"twilioApiKey"`
	// Phone number which will receive the notifications
	ReceivePhoneNumber string `json:"receivePhoneNumber"`
}
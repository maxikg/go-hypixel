/*
Package hypixel provides a client for using the Hypixel API.

Construct a new Hypixel client to access the various aspects of their public api. For example, to show an api keys
statistics:

	client := hypixel.NewClient("your-api-key", nil)
	response, err := client.KeyInfo()
 */
package hypixel

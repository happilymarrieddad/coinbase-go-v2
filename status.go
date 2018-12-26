package coinbasev2

// Status - status obj
var Status map[int]string

func init() {
	Status = make(map[int]string)

	Status[200] = "OK"
	Status[201] = "Object Saved"
	Status[204] = "No content"
	Status[400] = "Bad Request"
	Status[401] = "Couldn't authenticate your request"
	Status[402] = "2FA Token Required"
	Status[403] = "User hasn’t authorized necessary scope"
	Status[404] = "No such object"
	Status[429] = "Your connect is being rate limited"
	Status[500] = "Internal Server Error"
	Status[503] = "Your connection is being throttled or the service is down for maintenance"
}

// 200 OK Successful request
// 201 Created New object saved
// 204 No content Object deleted
// 400 Bad Request Returns JSON with the error message
// 401 Unauthorized Couldn’t authenticate your request
// 402 2FA Token required Re-try request with user’s 2FA token as CB-2FA-Token header
// 403 Invalid scope User hasn’t authorized necessary scope
// 404 Not Found No such object
// 429 Too Many Requests Your connection is being rate limited
// 500 Internal Server Error Something went wrong
// 503 Service Unavailable Your connection is being throttled or the service is down for maintenance

package coinbasev2

import "net/http"

// ShowResponse - show response
type ShowResponse struct {
	Data User `json:"data"`
}

// Show - GET https://api.coinbase.com/v2/users/:user_id
func Show(id string, auth string) (user User, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.coinbase.com/v2/users/"+id, nil)
	if err != nil {
		return
	}

	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)

	return
}

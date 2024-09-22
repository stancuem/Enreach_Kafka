package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func enricData(user User) User {
	resp, err := http.Post("http://dobat' url/exaple", "application/json", bytes.NewBuffer([]byte(fmt.Sprintf(`{"first_name":"%s","last_name":"%s"}`, user.FirstName, user.LastName))))
	if err != nil {
		return user
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return user
	}

	var enrichedUser User
	if err := json.Unmarshal(body, &enrichedUser); err != nil {
		return user
	}

	enrichedUser.FirstName = user.FirstName
	enrichedUser.LastName = user.LastName

	return enrichedUser
}

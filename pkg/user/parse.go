package user

import (
	"encoding/json"
	"os"
)

// Parse takes a path to a JSON file containing information on a
// Facebook user, and loads the data into a Profile struct. It returns a
// pointer to the Profile struct.
func Parse(filePath string) (*Profile, error) {
	var profile struct {
		Profile Profile `json:"profile_v2"`
	}

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(data, &profile); err != nil {
		return nil, err
	}

	return &profile.Profile, nil
}

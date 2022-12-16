package user

import (
	"encoding/json"
	"os"

	"github.com/samsapti/CleanBook/internal/utils"
)

func fixEncoding(p *Profile) {
	utils.FixEncoding(&p.Name.Full)
	utils.FixEncoding(&p.Name.First)
	utils.FixEncoding(&p.Name.Middle)
	utils.FixEncoding(&p.Name.Last)
	utils.FixEncoding(&p.Username)
}

// Parse takes a path to a JSON file containing information on a
// Facebook user, and loads the data into a Profile struct. It returns a
// pointer to the Profile struct.
func Parse(filePath string) (*Profile, error) {
	var p struct {
		Profile Profile `json:"profile_v2"`
	}

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Read JSON data
	if err = json.NewDecoder(file).Decode(&p); err != nil {
		return nil, err
	}

	// Fix encoding in strings
	fixEncoding(&p.Profile)

	return &p.Profile, nil
}

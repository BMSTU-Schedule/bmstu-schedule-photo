package parse

import "io/ioutil"

//easyjson:json
type Group struct {
	URL       string `json:"url"`
	GroupName string `json:"group"`
}

//easyjson:json
type Groups []*Group

// ParseJsonFile gets groups from JSON file and return Groups struct
func ParseJsonFile(pathToJSON string) (*Groups, error) {
	file, err := ioutil.ReadFile(pathToJSON)
	if err != nil {
		return nil, err
	}

	var groups Groups
	err = groups.UnmarshalJSON(file)
	if err != nil {
		return nil, err
	}

	return &groups, nil
}

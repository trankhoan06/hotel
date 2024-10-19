package upload

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Image struct {
	Url       string `json:"url"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Cloud     string `json:"cloud"`
	Extension string `json:"extension"`
}

func (Image) TableName() string { return "image" }
func (i *Image) Fullfill(s string) {
	i.Url = fmt.Sprintf("%s/%s", s, i.Url)
}
func (i *Image) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("Can't unmarshal JSONB value")
	}
	var img Image
	if err := json.Unmarshal(bytes, &img); err != nil {
		return err
	}

	*i = img
	return nil
}
func (i *Image) Value() (driver.Value, error) {
	if i == nil {
		return nil, nil
	}
	return json.Marshal(i)
}

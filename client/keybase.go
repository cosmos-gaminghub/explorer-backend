package client

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos-gaminghub/explorer-backend/logger"
	"github.com/cosmos-gaminghub/explorer-backend/utils"
)

const (
	KeybaseUrl = "https://keybase.io/_/api/1.0/user/lookup.json?key_suffix=%s&fields=pictures"
)

type KeyBase struct {
	Them []KeyBaseThem `json:"them"`
}

type KeyBaseThem struct {
	Id       string             `json:"id"`
	Pictures KeyBaseThemPriture `json:"pictures"`
}

type KeyBaseThemPriture struct {
	Primary struct {
		Url string `json:"url"`
	} `json:"primary"`
}

// GetBlock queries for a block by height. An error is returned if the query fails.
func GetKeyBase(identity string) (KeyBase, error) {
	url := fmt.Sprintf(KeybaseUrl, identity)
	resBytes, err := utils.Get(url)
	if err != nil {
		logger.Error("Get keybase error", logger.String("err", err.Error()))
	}

	var result KeyBase
	if err := json.Unmarshal(resBytes, &result); err != nil {
		logger.Error("Unmarshal keybase error", logger.String("err", err.Error()))
	}

	return result, nil
}

func GetImageUrl(identity string) (url string) {
	keybase, err := GetKeyBase(identity)

	if err != nil {
		return url
	}

	if len(keybase.Them) > 0 {
		return keybase.Them[0].Pictures.Primary.Url
	}

	return url
}

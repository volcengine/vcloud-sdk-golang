package imagex

import (
	"fmt"
	"net/url"
)

func (c *ImageX) GetImageThemeColor(uri string) (*GetImageThemeColorResult, error) {
	u := url.Values{}
	u.Set("Uri", uri)
	data, _, err := c.Client.Query("GetImageThemeColor", u)
	if err != nil {
		return nil, fmt.Errorf("fail to request api GetImageThemeColor, %v", err)
	}
	result := new(GetImageThemeColorResult)
	if err := UnmarshalResultInto(data, result); err != nil {
		return nil, err
	} else {
		return result, nil
	}
}

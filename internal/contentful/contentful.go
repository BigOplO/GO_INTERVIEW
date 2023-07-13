package contentful

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/BigOplO/GO_INTERVIEW/internal/db"
)

const contentfulAPI = "https://cdn.contentful.com/spaces/%s/entries/%s?access_token=%s"

var breadIDs = []string{
	"6QRk7gQYmOyJ1eMG9H4jbB",
	"41RUO5w4oIpNuwaqHuSwEc",
	"4Li6w5uVbJNVXYVxWjWVoZ",
}

type Response struct {
	Sys struct {
		ID        string `json:"id"`
		CreatedAt string `json:"createdAt"`
	} `json:"sys"`
	Fields struct {
		Name string `json:"name"`
	} `json:"fields"`
}

func FetchAndSave() error {
	for _, breadID := range breadIDs {
		url := fmt.Sprintf(contentfulAPI, "2vskphwbz4oc", breadID, "fV3hEZOxdoupON58JlInyV4UYFm2zTO3pkQRQCSS5KY")
		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		// Parse body to get bread info
		var r Response
		if err := json.Unmarshal(body, &r); err != nil {
			return err
		}

		bread := &db.Bread{
			ID:        r.Sys.ID,
			Name:      r.Fields.Name,
			CreatedAt: r.Sys.CreatedAt,
		}
		fmt.Println(bread)
		db.SaveBread(bread)
	}
	return nil
}

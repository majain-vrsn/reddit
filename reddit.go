// Package reddit implements a basic client for the Reddit API.
package reddit

import (
	"fmt"
	"net/http"
	"encoding/json"
	"errors"
)

type response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

// Item describes a Reddit item.
type Item struct {
	Title		string
	URL		string
	Comments	int 	`json:"num_comments"`
}

// Get fetches the most recent items posted to the specified subedit
func Get(reddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit);
	resp, err := http.Get(url);
        if(err != nil) {
		return nil, err;
        }

	defer resp.Body.Close();

        if(resp.StatusCode != http.StatusOK) {
                return nil, errors.New(resp.Status);
        }

        r := new(response)
        err = json.NewDecoder(resp.Body).Decode(r);
        if(err != nil) {
                return nil, err;
        }
	
	items := make([]Item, len(r.Data.Children));	
	for i, child := range r.Data.Children {
		items[i] = child.Data;
	}
	return items, nil;
}

func (item Item) String() string {
	com := "";
	
	switch item.Comments {
	case 0:
		//
	case 1:
		com = "(1 comment):";
	default:
		com = fmt.Sprintf("(%d comments):", item.Comments);
	}

	return fmt.Sprintf("%s %s ==> %s", com, item.Title, item.URL);
}


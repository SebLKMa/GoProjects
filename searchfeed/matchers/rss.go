package matchers

import (
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"../search"
)

type (
	// item contains the fields associated with the item tag in the rss document
	item struct {
		XMLName     xml.Name `xml:"item"`
		PubDate     string   `xml:"pubDate"`
		Title       string   `xml:"title"`
		Description string   `xml:"description"`		
		Link        string   `xml:"link"`
		GUID        string   `xml:"guid"`
		GeoRssPoint string   `xml:"georss:point"`
	}

	// image containsa the fields associated with the image tag in the rss document
	image struct {
		XMLName     xml.Name `xml:"image"`
		URL			string	 `xml:"url"`
		Title       string   `xml:"title"`
		Link        string   `xml:"link"`
	}

	// channel containsa the fields associated with the channel tag in the rss document
	channel struct {
		XMLName        xml.Name `xml:"channel"`
		Title          string   `xml:"title"`
		Description    string   `xml:"description"`
		Link           string   `xml:"link"`
		PubDate        string   `xml:"pubDate"`
		LastBuildDate  string   `xml:"lastBuildDate"`
		TTL            string   `xml:"ttl"`
		Language       string   `xml:"language"`
		ManagingEditor string   `xml:"managingEditor"`
		WebMaster      string   `xml:"webMaster"`
		Image          string   `xml:"image"`
		Item           []item   `xml:"item"`
	}

	// rssDocument containsa the fields associated with the rss document
	rssDocument struct {
		XMLName xml.Name `xml:"rss"`
		Channel channel	 `xml:"channel"`
	}
)

// Implements the Matcher interface, empty sttruct, no states, just implements Search() below
type rssMatcher struct{}

// Registers this matcher with the program
func init() {
	var matcher rssMatcher
	search.Register("rss", matcher)
}

// Invokes HTTP GET request for rss feed and decodes. This function is unexported.
func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil, errors.New("No rss feed URI provided")
	}

	// get the rss feed document from the web
	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil, err
	}

	// ensure response is closed on return
	defer resp.Body.Close()

	// check status 200 for valid response
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Response Error Code: %d\n", resp.StatusCode)
	}

	// decode rss feed and populate our struct, return err to caller (if any)
	var document rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&document)
	return &document, err
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result

	log.Printf("Search Feed Type[%s] Site[%s] For Uri[%s]\n", feed.Type, feed.Name, feed.URI)

	// retrieve the data to search
	document, err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _, channelItem := range document.Channel.Item {
		// check if title matches the provided search term
		matched, err := regexp.MatchString(searchTerm, channelItem.Title)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result {
				Field: "Title",
				Content: channelItem.Title,
			})
		}

		// check if description matches the provided search term
		matched, err = regexp.MatchString(searchTerm, channelItem.Description)
		if err != nil {
			return nil, err
		}

		if matched {
			results = append(results, &search.Result {
				Field: "Description",
				Content: channelItem.GeoRssPoint,
			})
		}
	}

	return results, nil
}


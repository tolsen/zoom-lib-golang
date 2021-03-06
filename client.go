package zoom

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	apiURI     = "api.zoom.us"
	apiVersion = "/v2"
)

var (
	// Debug causes debugging message to be printed, using the log package,
	// when set to true
	Debug = false

	// APIKey is a package-wide API key, used when no client is instantiated
	APIKey string

	// APISecret is a package-wide API secret, used when no client is instantiated
	APISecret string

	defaultClient *Client
)

// Client is responsible for making API requests
type Client struct {
	Key       string
	Secret    string
	Transport http.RoundTripper
	Timeout   time.Duration // set to value > 0 to enable a request timeout
	endpoint  string
}

// NewClient returns a new API client
func NewClient(apiKey string, apiSecret string) *Client {
	var uri = url.URL{
		Scheme: "https",
		Host:   apiURI,
		Path:   apiVersion,
	}

	return &Client{
		Key:      apiKey,
		Secret:   apiSecret,
		endpoint: uri.String(),
	}
}

type RequestV2Opts struct {
	Client         *Client
	Method         HTTPMethod
	URLParameters  interface{}
	Path           string
	DataParameters interface{}
	Ret            interface{}
	// HeadResponse represents responses that don't have a body
	HeadResponse bool
}

func initializeDefault(c *Client) *Client {
	if c == nil {
		if defaultClient == nil {
			defaultClient = NewClient(APIKey, APISecret)
		}

		return defaultClient
	}

	return c
}

func (c *Client) executeRequest(opts RequestV2Opts) (*http.Response, error) {
	client := c.httpClient()
	req, err := c.addRequestAuth(c.httpRequest(opts))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	return client.Do(req)
}

func (c *Client) httpRequest(opts RequestV2Opts) (*http.Request, error) {
	var buf bytes.Buffer

	// encode body parameters if any
	if err := json.NewEncoder(&buf).Encode(&opts.DataParameters); err != nil {
		return nil, err
	}

	// set URL parameters
	values, err := query.Values(opts.URLParameters)
	if err != nil {
		return nil, err
	}

	// set request URL
	requestURL := c.endpoint + opts.Path
	if len(values) > 0 {
		requestURL += "?" + values.Encode()
	}

	if Debug {
		log.Printf("Request URL: %s", requestURL)
		log.Printf("URL Parameters: %s", values.Encode())
		log.Printf("Body Parameters: %s", buf.String())
	}

	// create HTTP request
	return http.NewRequest(string(opts.Method), requestURL, &buf)
}

func (c *Client) httpClient() *http.Client {
	client := &http.Client{Transport: c.Transport}
	if c.Timeout > 0 {
		client.Timeout = c.Timeout
	}

	return client
}

func (c *Client) RequestV2(opts RequestV2Opts) error {
	// make sure the defaultClient is not nil if we are using it
	c = initializeDefault(c)

	var backoffDurations []time.Duration = []time.Duration{
		time.Minute,
		5 * time.Minute,
		10 * time.Minute,
		40 * time.Minute,
		90 * time.Minute,
	}

	// execute HTTP request
	var err error
	for _, backoffDuration := range backoffDurations {
		var resp *http.Response
		resp, err = c.executeRequest(opts)
		if err != nil {
			if strings.Contains(err.Error(), "GOAWAY") {
				fmt.Printf("Got GOAWAY. Sleeping %v . err=%v\n", backoffDuration, err)
				time.Sleep(backoffDuration)
				continue
			}

			return err
		}

		// If there is no body in response
		if opts.HeadResponse {
			return c.RequestV2HeadOnly(resp)
		}

		if err = c.RequestV2WithBody(opts, resp); err != nil {
			if strings.Contains(err.Error(), "GOAWAY") {
				fmt.Printf("Got GOAWAY. Sleeping %v . err=%v\n", backoffDuration, err)
				time.Sleep(backoffDuration)
				continue
			}
			return err
		}

		return nil
	}

	return err
}

func (c *Client) RequestV2WithBody(opts RequestV2Opts, resp *http.Response) error {
	// read HTTP response
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if Debug {
		log.Printf("Response Body: %s", string(body))
	}

	// check for Zoom errors in the response
	if err := checkError(body); err != nil {
		return err
	}

	// unmarshal the response body into the return object
	return json.Unmarshal(body, &opts.Ret)
}

func (c *Client) RequestV2HeadOnly(resp *http.Response) error {
	if resp.StatusCode != 204 {
		return errors.New(resp.Status)
	}

	// there were no errors, just return
	return nil
}

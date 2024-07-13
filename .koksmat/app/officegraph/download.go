package officegraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/viper"
)

// GraphResponse represents a typical response from Microsoft Graph API with paging
type GraphResponse struct {
	Value    []interface{} `json:"value"`
	NextLink string        `json:"@odata.nextLink"`
}

// FetchGraphObjects fetches data from the Microsoft Graph API, handling pagination and throttling
func FetchGraphObjects(maxPages int, url string, accessToken string) (string, error) {
	var allData []interface{}
	client := &http.Client{}

	for i := 0; i < maxPages; i++ {
		log.Printf("Fetching page %d\n", i+1)
		log.Printf("URL: %s\n", url)
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return "", fmt.Errorf("failed to create request: %v", err)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
		req.Header.Set("Accept", "application/json")
		req.Header.Set("Accept", "*/*")
		resp, err := client.Do(req)
		if err != nil {
			return "", fmt.Errorf("request failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode == 429 {
			retryAfter := resp.Header.Get("Retry-After")
			if retryAfter != "" {
				retrySeconds, err := strconv.Atoi(retryAfter)
				if err == nil {
					time.Sleep(time.Duration(retrySeconds) * time.Second)
					i-- // retry the same page
					continue
				}
			}
			time.Sleep(5 * time.Second)
			i-- // retry the same page
			continue
		}

		if resp.StatusCode != http.StatusOK {
			return "", fmt.Errorf("request failed with status: %s", resp.Status)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to read response body: %v", err)
		}

		var result GraphResponse
		err = json.Unmarshal(body, &result)
		if err != nil {
			return "", fmt.Errorf("failed to parse response: %v", err)
		}

		allData = append(allData, result.Value...)

		if result.NextLink == "" {
			break
		}
		url = result.NextLink
	}

	allDataBytes, err := json.Marshal(allData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal all data: %v", err)
	}

	return string(allDataBytes), nil
}

func FetchGraphObject(url string, accessToken string) (string, error) {

	client := &http.Client{}

	log.Printf("URL: %s\n", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept", "*/*")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("request failed with status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %v", err)
	}

	return string(body), nil
}
func Download(url string, accessToken string, maxPages int) (*string, error) {

	data, err := FetchGraphObjects(maxPages, url, accessToken)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return nil, err
	}
	return &data, nil

}

func CacheObject(tag string, actor string, url string, token string) error {

	json, err := FetchGraphObject(url, token)
	if err != nil {
		log.Println("Error fetching data: ", err)
		return err
	}

	connectionString := viper.GetString("POSTGRES_DB")
	conn, err := ConnectDB(connectionString)
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return err
	}
	defer conn.Close()
	err = InsertIntoGraphCacheTable(context.Background(), conn, tag, tag, actor, token, url, json)

	if err != nil {
		log.Println("Error inserting data: ", err)
		return err
	}
	return nil
}

func CacheObjects(tag string, actor string, url string, token string, maxPages int) error {

	json, err := FetchGraphObjects(maxPages, url, token)
	if err != nil {
		log.Println("Error fetching data: ", err)
		return err
	}

	connectionString := viper.GetString("POSTGRES_DB")
	conn, err := ConnectDB(connectionString)
	if err != nil {
		log.Println("Error connecting to database: ", err)
		return err
	}
	defer conn.Close()
	err = InsertIntoGraphCacheTable(context.Background(), conn, tag, tag, actor, token, url, json)

	if err != nil {
		log.Println("Error inserting data: ", err)
		return err
	}
	return nil
}

package main

import (
	"context"
	"encoding/json"
	"log"
	"math/rand"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

// User represents the structure of user data in Firestore
type User struct {
	Badges    []map[string]interface{} `json:"badges"`
	Countries map[string]Country       `json:"countries"`
}

type Region struct {
	RegionCode string `json:"regionCode"`
	// Add other relevant fields from the JSON structure
}

type Country struct {
	CountryCode string            `json:"countryCode"`
	Regions     map[string]Region `json:"regions"`
}

// Country represents the structure of country data in Firestore
// type Country struct {
// 	CountryCode string `json:"countryCode"`
// }

// EuroExplorerBadgeID is the unique identifier for the Euro Explorer badge
const EuroExplorerBadgeID = "euro_explorer"
const AsianExplorerBadgeID = "asian_explorer"
const GlobetrotterBadgeID = "globetrotter"
const PolarExplorerBadgeID = "polar_explorer"
const OzzieExplorerBadgeID = "ozzie_explorer"

// Function to generate a random color code
func randomColor() string {
	colors := []string{"#FF5733", "#33FF57", "#3357FF", "#FF33F5", "#FFC300"}
	return colors[rand.Intn(len(colors))]
}

// Function to get a list of country codes based on a criteria
func getCountryCodes(userCountries map[string]Country, criteria map[string]bool) []string {
	var codes []string
	for code := range userCountries {
		if criteria[code] {
			codes = append(codes, code)
		}
	}
	return codes
}

// structToMap converts a struct to a map using json marshaling
func structToMap(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}

	var mapData map[string]interface{}
	if err := json.Unmarshal(data, &mapData); err != nil {
		return nil, err
	}

	return mapData, nil
}

func connectFirestore(ctx context.Context) (*firestore.Client, error) {
	sa := option.WithCredentialsFile("/Users/adglad/Downloads/triparific100-fa4917b5dd13.json") // Replace with your JSON file path
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {
	ctx := context.Background()
	client, err := connectFirestore(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Firestore: %v", err)
	}
	defer client.Close()

	// Get a handle for your collection
	iter := client.Collection("users").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break
			}
			log.Fatalf("Failed to iterate: %v", err)
		}

		var user User
		err = doc.DataTo(&user)
		if err != nil {
			log.Printf("Failed to read document: %v", err)
			continue
		}

		// processUser(&user)
		processAsianExplorerBadge(&user) // Process Asian Explorer badge
		processGlobetrotterBadge(&user)  // Process Globetrotter badge
		processPolarExplorerBadge(&user) // Process Polar Explorer badge
		processEuroExplorerBadge(&user)  // Process Polar Explorer badge
		processOzzieExplorerBadge(&user) // Process Polar Explorer badge

		// Convert user struct to a map for updating
		userMap, err := structToMap(user)
		if err != nil {
			log.Printf("Failed to convert user to map: %v", err)
			continue
		}

		_, err = doc.Ref.Set(ctx, userMap, firestore.MergeAll)
		if err != nil {
			log.Printf("Failed to update user: %v", err)
		}
	}
}

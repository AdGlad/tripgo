package main

import (
	"context"
	"log"
	"strings"
	"time"
	"unicode"

	"cloud.google.com/go/firestore"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// City represents a city with a name and country.
type City struct {
	Name    string `firestore:"city"`
	Country string `firestore:"country"`
}

// CitiesDocument represents the structure of the Firestore document.
type CitiesDocument struct {
	Top50 []City `firestore:"top50"`
}

// normalizeString transforms a string to a normalized form without diacritics.
func normalizeString(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: Mark, Nonspacing
	}), norm.NFC)
	result, _, _ := transform.String(t, s)
	return strings.ToLower(result)
}
func processCityExplorerBadge(user *User, firestoreClient *firestore.Client) {
	log.Printf("Processing CityExplorerBadge")
	// Fetch the top 50 cities list from Firestore
	ctx := context.Background()
	citiesRef := firestoreClient.Collection("toppoi").Doc("cities")
	citiesDoc, err := citiesRef.Get(ctx)
	if err != nil {
		log.Fatalf("Error fetching top cities: %v", err)
		return
	}

	var citiesDocStruct CitiesDocument
	if err := citiesDoc.DataTo(&citiesDocStruct); err != nil {
		log.Fatalf("Error parsing top cities data: %v", err)
		return
	}

	// Create a map for easier comparison of top cities
	topCities := make(map[string]bool)
	for _, city := range citiesDocStruct.Top50 {
		key := normalizeString(city.Name) + "," + strings.ToLower(city.Country)
		topCities[key] = true
	}

	// Step 1: Generate a discrete list of visited cities
	visitedCities := make(map[string]bool)
	for _, country := range user.Countries {
		countryCode := strings.ToLower(country.CountryCode)
		for _, region := range country.Regions {
			for _, placeHistory := range region.PlaceHistory {
				cityCountryCombo := normalizeString(placeHistory.City) + "," + countryCode
				visitedCities[cityCountryCombo] = true
			}
		}
	}

	// Step 2: Compare the discrete list to top cities
	for cityCountryCombo := range visitedCities {
		if topCities[cityCountryCombo] {
			// Splitting cityCountryCombo to log city and country separately
			parts := strings.Split(cityCountryCombo, ",")
			city, country := parts[0], parts[1]
			badgeId := normalizeString(city) + "-" + country

			log.Printf("Match found! City: %s, Country: %s", city, country)

			// Check if the user already has the badge
			alreadyHasBadge := false
			for _, badge := range user.Badges {
				if id, ok := badge["badgeId"].(string); ok && id == badgeId {
					log.Printf("Badge already awarded: %s", badgeId)
					alreadyHasBadge = true
					break
				}
			}

			if !alreadyHasBadge {
				newBadge := map[string]interface{}{
					"achievedOn":  time.Now().Format(time.RFC3339),
					"badgeId":     badgeId,
					"criteria":    "Visit a Top City",
					"description": "Awarded for visiting " + city,
					"name":        "City Explorer: " + strings.Title(city),
					"type":        "City",
					"color":       randomColor(),
					"icon":        "city_symbol",
				}
				user.Badges = append(user.Badges, newBadge)
				log.Printf("New badge awarded: %s", badgeId)
			}

			// Implement logic to award a badge, for example
		} else {
			parts := strings.Split(cityCountryCombo, ",")
			city, country := parts[0], parts[1]
			log.Printf("No match found for City: %s, Country: %s", city, country)
		}
	}
}

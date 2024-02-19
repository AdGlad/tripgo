// utils.go
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// normalizeString transforms a string to a normalized form without diacritics.
func normalizeString(s string) string {
	t := transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool {
		return unicode.Is(unicode.Mn, r) // Mn: Mark, Nonspacing
	}), norm.NFC)
	result, _, _ := transform.String(t, s)
	return strings.ToLower(result)
}

// structToMap converts a Go struct to a map[string]interface{} using JSON marshaling and unmarshaling.
// This is useful for converting structs to a format that can be used with Firestore.
func structToMap(obj interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Printf("Failed to marshal object to JSON: %v", err)
		return nil, err
	}

	var mapData map[string]interface{}
	if err := json.Unmarshal(data, &mapData); err != nil {
		log.Printf("Failed to unmarshal JSON to map: %v", err)
		return nil, err
	}

	return mapData, nil
}

// Function to generate a random color code
// func randomColor() string {
// colors := []string{"#FF5733", "#33FF57", "#3357FF", "#FF33F5", "#FFC300"}
// return colors[rand.Intn(len(colors))]
// }
//
// // Function to get a list of country codes based on a criteria
// func getCountryCodes(userCountries map[string]Country, criteria map[string]bool) []string {
// var codes []string
// for code := range userCountries {
// if criteria[code] {
// codes = append(codes, code)
// }
// }
// return codes
// }
// randomColor returns a random color from a predefined list.
func randomColor() string {
	colors := []string{"#FF5733", "#33FF57", "#3357FF", "#FF33F5", "#FFC300"}
	rand.Seed(time.Now().UnixNano()) // Ensure different outcomes
	return colors[rand.Intn(len(colors))]
}

// getCountryCodes returns a slice of country codes based on a given criteria.
func getCountryCodes(userCountries map[string]Country, criteria map[string]bool) []string {
	var codes []string
	for code, country := range userCountries {
		if _, ok := criteria[country.CountryCode]; ok {
			codes = append(codes, code)
		}
	}
	return codes
}

func printUserAsJSON(user User) {
	jsonData, err := json.MarshalIndent(user, "", "  ") // Pretty print with 2 spaces indentation
	if err != nil {
		log.Printf("Failed to marshal user to JSON: %v", err)
		return
	}
	log.Println(string(jsonData))
}

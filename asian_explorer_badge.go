package main

import (
	"time"
)

func processAsianExplorerBadge(user *User) {
	// Define a list of Asian country codes (ISO Alpha-2)
	asianCountries := map[string]bool{
		"cn": true, "jp": true, "in": true, // ... add other Asian country codes
		// ...
		"th": true, "sg": true,
	}

	// Count how many Asian countries the user has visited
	var count int
	for countryCode := range user.Countries {
		if asianCountries[countryCode] {
			count++
		}
	}

	// Check if the user already has the Asian Explorer badge
	alreadyHasBadge := false
	for _, badge := range user.Badges {
		if badgeId, ok := badge["badgeId"].(string); ok && badgeId == AsianExplorerBadgeID {
			alreadyHasBadge = true
			break
		}
	}

	// Add the Asian Explorer badge if the user qualifies and doesn't already have it
	if count >= 5 && !alreadyHasBadge {
		newBadge := map[string]interface{}{
			"achievedOn":   time.Now().Format(time.RFC3339),
			"badgeId":      AsianExplorerBadgeID,
			"criteria":     "Visit 5 Asian Countries",
			"description":  "Awarded for visiting 5 Asian countries",
			"name":         "Asian Explorer",
			"type":         "Region",
			"color":        randomColor(),
			"icon":         "asia_symbol",
			"countryCodes": getCountryCodes(user.Countries, asianCountries),
		}
		user.Badges = append(user.Badges, newBadge)
	}
}

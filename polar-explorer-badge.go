package main
import (
	"time"
)

func processPolarExplorerBadge(user *User) {
	// Define a list of country codes associated with Arctic or Antarctic regions
	polarCountries := map[string]bool{
		"no": true, // Norway
		"gl": true, // Greenland
		"ru": true, // Russia
		"ar": true, // Argentina
		"cl": true, // Chile
	}

	// Check if the user has visited any of the polar countries
	hasVisitedPolar := false
	for countryCode := range user.Countries {
		if polarCountries[countryCode] {
			hasVisitedPolar = true
			break
		}
	}

	// Check if the user already has the Polar Explorer badge
	alreadyHasBadge := false
	for _, badge := range user.Badges {
		if badgeId, ok := badge["badgeId"].(string); ok && badgeId == PolarExplorerBadgeID {
			alreadyHasBadge = true
			break
		}
	}

	// Add the Polar Explorer badge if the user qualifies and doesn't already have it
	if hasVisitedPolar && !alreadyHasBadge {
		newBadge := map[string]interface{}{
			"achievedOn":  time.Now().Format(time.RFC3339),
			"badgeId":     PolarExplorerBadgeID,
			"criteria":    "Visited Arctic or Antarctic regions",
			"description": "Awarded for visiting the Arctic or Antarctic regions",
			"name":        "Polar Explorer",
		}
		user.Badges = append(user.Badges, newBadge)
	}
}

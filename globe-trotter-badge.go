package main

import (
	"time"
)

func processGlobetrotterBadge(user *User) {
	// Count the total number of different countries visited
	countryCount := len(user.Countries)

	// Check if the user already has the Globetrotter badge
	alreadyHasBadge := false
	for _, badge := range user.Badges {
		if badgeId, ok := badge["badgeId"].(string); ok && badgeId == GlobetrotterBadgeID {
			alreadyHasBadge = true
			break
		}
	}

	// Add the Globetrotter badge if the user qualifies and doesn't already have it
	if countryCount > 30 && !alreadyHasBadge {
		// Get the first 30 countries
		countryCodes := getFirst30Countries(user.Countries)

		newBadge := map[string]interface{}{
			"achievedOn":   time.Now().Format(time.RFC3339),
			"badgeId":      GlobetrotterBadgeID,
			"criteria":     "Visit more than 30 different countries",
			"description":  "Awarded for visiting more than 30 different countries",
			"name":         "Globetrotter",
			"color":        randomColor(),
			"icon":         "globe_symbol",
			"countryCodes": countryCodes,
		}
		user.Badges = append(user.Badges, newBadge)
	}
}

// getFirst30Countries returns the first 30 countries from the user's visited countries
func getFirst30Countries(userCountries map[string]Country) []string {
	var codes []string
	for code := range userCountries {
		codes = append(codes, code)
		if len(codes) == 30 {
			break
		}
	}
	return codes
}

// func processGlobetrotterBadge(user *User) {
//     // Count the total number of different countries visited
//     countryCount := len(user.Countries)

//     // Check if the user already has the Globetrotter badge
//     alreadyHasBadge := false
//     for _, badge := range user.Badges {
//         if badgeId, ok := badge["badgeId"].(string); ok && badgeId == GlobetrotterBadgeID {
//             alreadyHasBadge = true
//             break
//         }
//     }

//     // Add the Globetrotter badge if the user qualifies and doesn't already have it
//     if countryCount > 30 && !alreadyHasBadge {
//         newBadge := map[string]interface{}{
//             "achievedOn":   time.Now().Format(time.RFC3339),
//             "badgeId":      GlobetrotterBadgeID,
//             "criteria":     "Visit more than 30 different countries",
//             "description":  "Awarded for visiting more than 30 different countries",
//             "name":         "Globetrotter",
//             "color":        randomColor(),
//             "icon":         "globe_symbol",
//             "countryCodes": getCountryCodes(user.Countries, user.Countries), // All countries
//         }
//         user.Badges = append(user.Badges, newBadge)
//     }
// }

// processGlobetrotterBadge checks if the user qualifies for the Globetrotter badge and adds it if they do
// func processGlobetrotterBadge(user *User) {
// 	// Count the total number of different countries visited
// 	countryCount := len(user.Countries)

// 	// Check if the user already has the Globetrotter badge
// 	alreadyHasBadge := false
// 	for _, badge := range user.Badges {
// 		if badgeId, ok := badge["badgeId"].(string); ok && badgeId == GlobetrotterBadgeID {
// 			alreadyHasBadge = true
// 			break
// 		}
// 	}

// 	// Add the Globetrotter badge if the user qualifies and doesn't already have it
// 	if countryCount > 30 && !alreadyHasBadge {
// 		newBadge := map[string]interface{}{
// 			"achievedOn":  time.Now().Format(time.RFC3339),
// 			"badgeId":     GlobetrotterBadgeID,
// 			"criteria":    "Visit more than 30 different countries",
// 			"description": "Awarded for visiting more than 30 different countries",
// 			"name":        "Globetrotter",
// 		}
// 		user.Badges = append(user.Badges, newBadge)
// 	}
// }

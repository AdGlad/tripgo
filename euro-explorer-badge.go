package main

import (
	"time"
)

func processEuroExplorerBadge(user *User) {
	// Define a list of European country codes (ISO Alpha-2)

europeanCountries := map[string]bool{
    "al": true, "ad": true, "am": true, "at": true, "az": true, "by": true,
    "be": true, "ba": true, "bg": true, "hr": true, "cy": true, "cz": true,
    "dk": true, "ee": true, "fi": true, "fr": true, "ge": true, "de": true,
    "gr": true, "hu": true, "is": true, "ie": true, "it": true, "kz": true,
    "xk": true, "lv": true, "li": true, "lt": true, "lu": true, "mt": true,
    "md": true, "mc": true, "me": true, "nl": true, "mk": true, "no": true,
    "pl": true, "pt": true, "ro": true, "ru": true, "sm": true, "rs": true,
    "sk": true, "si": true, "es": true, "se": true, "ch": true, "tr": true,
    "ua": true, "gb": true,
}

	// Initialize count and alreadyHasBadge at the beginning of the function
	var count int
	alreadyHasBadge := false

	// Count how many European countries the user has visited
	for countryCode := range user.Countries {
		if europeanCountries[countryCode] {
			count++
		}
	}

	// Check if the user already has the Euro Explorer badge
	for _, badge := range user.Badges {
		if badgeId, ok := badge["badgeId"].(string); ok && badgeId == EuroExplorerBadgeID {
			alreadyHasBadge = true
			break
		}
	}

	// Add the Euro Explorer badge if the user qualifies and doesn't already have it
	if count >= 5 && !alreadyHasBadge {
		newBadge := map[string]interface{}{
			"achievedOn":   time.Now().Format(time.RFC3339),
			"badgeId":      EuroExplorerBadgeID,
			"criteria":     "Visit 5 European Countries",
			"description":  "Awarded for visiting 5 European countries",
			"name":         "Euro Explorer",
			"type":         "Region",
			"color":        randomColor(),
			"icon":         "euro_symbol",
			"countryCodes": getCountryCodes(user.Countries, europeanCountries),
		}
		user.Badges = append(user.Badges, newBadge)
	}
}

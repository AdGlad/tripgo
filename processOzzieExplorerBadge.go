package main

import (
	"log"
	"time"
)

func processOzzieExplorerBadge(user *User) {
	// Define a list of all Australian region codes
	allAustralianRegions := map[string]bool{
		"AU-NSW": true, // New South Wales
		"AU-QLD": true, // Queensland
		"AU-SA":  true, // South Australia
		"AU-TAS": true, // Tasmania
		"AU-VIC": true, // Victoria
		"AU-WA":  true, // Western Australia
		"AU-ACT": true, // Australian Capital Territory
		"AU-NT":  true, // Northern Territory
	}

	// Check if the user has visited Australia
	australianRegionsVisited, visitedAustralia := user.Countries["au"]
	if !visitedAustralia {
		log.Printf("User has not visited Australia")

		return // User has not visited Australia, so no need to proceed
	}

	// Check if the user has visited all regions in Australia
	hasVisitedAllRegions := true
	for regionCode := range allAustralianRegions {
		if _, visited := australianRegionsVisited.Regions[regionCode]; !visited {
			hasVisitedAllRegions = false
			break
		}
	}

	// Check if the user already has the Ozzie Explorer badge
	alreadyHasBadge := false
	for _, badge := range user.Badges {
		if badgeId, ok := badge["badgeId"].(string); ok && badgeId == OzzieExplorerBadgeID {
			log.Printf("User already has Ozzie Explorer badge")

			alreadyHasBadge = true
			break
		}
	}

	// Add the Ozzie Explorer badge if the user qualifies and doesn't already have it
	if hasVisitedAllRegions && !alreadyHasBadge {
		log.Printf("Awarding Ozzie Explorer badge to user")

		newBadge := map[string]interface{}{
			"achievedOn":   time.Now().Format(time.RFC3339),
			"badgeId":      OzzieExplorerBadgeID,
			"criteria":     "Visit all Australian regions",
			"description":  "Awarded for visiting all Australian regions",
			"name":         "Ozzie Explorer",
			"color":        randomColor(),
			"icon":         "australia_symbol",
			"countryCodes": []string{"AU"}, // Australian country code

		}
		user.Badges = append(user.Badges, newBadge)
	}
}

// func processOzzieExplorerBadge(user *User) {
// 	australianRegions := map[string]bool{
// 		"AU-NSW": true, // New South Wales
//         "AU-QLD": true, // Queensland
//         "AU-SA":  true, // South Australia
//         "AU-TAS": true, // Tasmania
//         "AU-VIC": true, // Victoria
//         "AU-WA":  true, // Western Australia
//         "AU-ACT": true, // Australian Capital Territory
//         "AU-NT":  true, // Northern Territory
//     	}

// 	// Assuming 'au' is the country code for Australia in your data
// 	australianData, found := user.Countries["au"]
// 	if !found {
// 		log.Printf("User has not visited Australia")
// 		return
// 	}

//     hasVisitedAllRegions := true
//     var visitedRegions []string
//     for regionCode := range australianRegions {
//         if _, visited := user.Countries[regionCode]; visited {
//             visitedRegions = append(visitedRegions, regionCode)
//         } else {
//             hasVisitedAllRegions = false
//         }
//     }
// 	alreadyHasBadge := false
// 	for _, badge := range user.Badges {
// 		if badge["badgeId"] == "ozzie_explorer" {
// 			alreadyHasBadge = true
// 			log.Printf("User already has Ozzie Explorer badge")
// 			break
// 		}
// 	}

// 	if hasVisitedAllRegions && !alreadyHasBadge {
// 		log.Printf("Awarding Ozzie Explorer badge to user")
// 		newBadge := map[string]interface{}{
// 			"achievedOn":  time.Now().Format(time.RFC3339),
// 			"badgeId":     "ozzie_explorer",
// 			"criteria":    "Visit all Australian regions",
// 			"description": "Awarded for visiting all Australian regions",
// 			"name":        "Ozzie Explorer",
// 			"color":       randomColor(),
// 			"icon":        "australia_symbol",
//             "countryCodes": visitedRegions,

// 		}
// 		user.Badges = append(user.Badges, newBadge)
// 	}
// }

// func processOzzieExplorerBadge(user *User) {
// 	// Define a list of Australian region codes
// 	australianRegions := map[string]bool{
// 		"AU-NSW": true, // New South Wales
// 		// "AU-QLD": true, // Queensland
// 		// "AU-SA":  true, // South Australia
// 		// "AU-TAS": true, // Tasmania
// 		// "AU-VIC": true, // Victoria
// 		// "AU-WA":  true, // Western Australia
// 		// "AU-ACT": true, // Australian Capital Territory
// 		// "AU-NT":  true, // Northern Territory
// 	}

// 	hasVisitedAllRegions := true
// 	for regionCode := range australianRegions {
// 		log.Printf("User regionCode visited: %s", regionCode)
// 		_, exists := user.Countries[regionCode]
// 		if !exists {
// 			hasVisitedAllRegions = false
// 			log.Printf("User has not visited: %s", regionCode)
// 			break
// 		} else {
// 			log.Printf("User regionCode visited: %s", regionCode)
// 		}
// 	}

// 	// Check if the user has visited all Australian regions
// 	// hasVisitedAllRegions := true
// 	// for regionCode := range australianRegions {
// 	// 	log.Printf("User regionCode visited: %s", regionCode) // Print missing region

// 	// 	if _, visited := user.Countries[regionCode]; !visited {
// 	// 		hasVisitedAllRegions = false
// 	// 		log.Printf("User has not visited: %s", regionCode) // Print missing region

// 	// 		break
// 	// 	}
// 	// }

// 	// Check if the user already has the Ozzie Explorer badge
// 	alreadyHasBadge := false
// 	for _, badge := range user.Badges {
// 		if badge["badgeId"] == "ozzie_explorer" {
// 			alreadyHasBadge = true
// 			log.Printf("User already has Ozzie Explorer badge") // Print if badge already exists

// 			break
// 		}
// 	}

// 	// Add the Ozzie Explorer badge if the user qualifies and doesn't already have it
// 	if hasVisitedAllRegions && !alreadyHasBadge {
// 		log.Printf("Awarding Ozzie Explorer badge to user") // Print before awarding the badge

// 		newBadge := map[string]interface{}{
// 			"achievedOn":  time.Now().Format(time.RFC3339),
// 			"badgeId":     "ozzie_explorer",
// 			"criteria":    "Visit all Australian regions",
// 			"description": "Awarded for visiting all Australian regions",
// 			"name":        "Ozzie Explorer",
// 			"color":       randomColor(),
// 			"icon":        "australia_symbol",
// 		}
// 		user.Badges = append(user.Badges, newBadge)
// 	}
// }

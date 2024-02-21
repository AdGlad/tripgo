package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// City and CitiesDocument represent the structure of the Firestore document for top cities.

// TopCitiesDocument represents the structure expected in the Firestore document.
// type CitiesDocument struct {
// 	Top50 []City `firestore:"top50"`
// }

// fetchTopCities retrieves the list of top cities from a Firestore document.
// fetchTopCities retrieves the list of top cities from Firestore within the specified structure.
func fetchTopCities(ctx context.Context, firestoreClient *firestore.Client) ([]City, error) {
	docRef := firestoreClient.Collection("toppoi").Doc("cities")
	docSnapshot, err := docRef.Get(ctx)
	if err != nil {
		log.Printf("Failed to fetch top cities document: %v", err)
		return nil, err
	}

	var citiesDoc CitiesDocument
	if err := docSnapshot.DataTo(&citiesDoc); err != nil {
		log.Printf("Failed to decode top cities document: %v", err)
		return nil, err
	}

	return citiesDoc.Top50, nil
}

func processCityExplorerBadge(ctx context.Context, firestoreClient *firestore.Client) {
	topCities, err := fetchTopCities(ctx, firestoreClient)
	if err != nil {
		log.Fatalf("Error fetching top cities: %v", err)
	}

	users, err := fetchAllUsers(ctx, firestoreClient) // Using fetchAllUsers now
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}

	for _, user := range users { // Iterating over User objects now
		log.Printf("user %s: ", user.UserID)

		placeHistories, err := fetchUserPlaceHistories(ctx, firestoreClient, user.UserID) // Use user.UserID
		if err != nil {
			log.Printf("Error fetching place histories for user %s: %v", user.UserID, err)
			continue
		}

		for _, placeHistory := range placeHistories {
			log.Printf("user: %s placeHistory city: %s ", user.UserID, placeHistory.City)

			cityCountryCombo := normalizeString(placeHistory.City) + "," + strings.ToLower(placeHistory.CountryCode)

			for _, topCity := range topCities {
				log.Printf("user: %s placeHistory city: %s topCity: %s Country: %s", user.UserID, placeHistory.City, topCity.Name, topCity.Country)

				topCityCombo := normalizeString(topCity.Name) + "," + strings.ToLower(topCity.Country)

				if cityCountryCombo == topCityCombo {
					// Check if badge already awarded
					badgeId := normalizeString(topCity.Name) + "-" + strings.ToLower(topCity.Country)

					alreadyHasBadge := false
					// for _, badge := range user.Badges {
					// 	if badgeName, ok := badge["name"].(string); ok && badgeName == "CityExplorer" {
					// 		log.Printf("City Explorer badge already awarded to user %s", user.UserID)
					// 		alreadyHasBadge = true
					// 		break
					// 	}
					// }

					for _, badge := range user.Badges {
						log.Printf("badge: %s badgeId: %s ", badge["badgeId"].(string), badgeId)

						if id, ok := badge["badgeId"].(string); ok && id == badgeId {
							log.Printf("Badge already awarded: %s", badgeId)
							alreadyHasBadge = true
							break
						}
					}

					if !alreadyHasBadge {
						// newBadge := map[string]interface{}{
						// 	"achievedOn":  time.Now(),
						// 	"badgeId":     badgeId,
						// 	"criteria":    "Visit a Top City",
						// 	"description": "Awarded for visiting " + normalizeString(topCity.Name),
						// 	"name":        "City Explorer: " + normalizeString(topCity.Name),
						// 	"type":        "City",
						// 	"color":       randomColor(),
						// 	"icon":        "city_symbol",
						// }
						newBadge := Badge{
							BadgeId:     badgeId, // Assume badgeId is a unique identifier for the badge
							Name:        "City Explorer: " + strings.Title(normalizeString(topCity.Name)),
							Type:        "City",
							Description: "Awarded for visiting " + normalizeString(topCity.Name),
							Criteria:    "Visit a Top City",
							Color:       randomColor(), // Ensure you have a function to generate a random color
							Icon:        "city_symbol",
							//	AchievedOn:  time.Now().Format(time.RFC3339),
							AchievedOn: time.Now(),
						}
						//user.Badges = append(user.Badges, newBadge)
						log.Printf("New badge awarded: %s", badgeId)

						awardBadge(ctx, firestoreClient, user.UserID, newBadge)

						//awardBadge(ctx, firestoreClient, user.UserID, Badge{
						//Name:        "CityExplorer",
						//AwardedOn:   time.Now().Format(time.RFC3339),
						//Description: "Awarded for visiting a top city",
						//})
						log.Printf("Awarded City Explorer badge to user %s for visiting %s", user.UserID, placeHistory.City)
					}
					break // Stop checking after awarding the badge
				}
			}
		}
	}
}

// fetchUserPlaceHistories queries Firestore for place history records associated with a given userID.
func fetchUserPlaceHistories(ctx context.Context, firestoreClient *firestore.Client, userID string) ([]PlaceHistory, error) {
	var placeHistories []PlaceHistory

	// Query the placeHistory collection for records where userId matches the provided userID.
	iter := firestoreClient.Collection("placehistory").Where("userId", "==", userID).Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break // Exit the loop when all documents have been processed.
			}
			log.Printf("Error iterating documents: %v", err)
			return nil, err // Return nil slice and an error.
		}

		var placeHistory PlaceHistory
		if err := doc.DataTo(&placeHistory); err != nil {
			log.Printf("Error decoding document to Placehistory: %v", err)
			continue // Optionally, handle the error differently.
		}

		// Append the fetched PlaceHistory to the slice.
		placeHistories = append(placeHistories, placeHistory)
	}

	return placeHistories, nil // Return the slice of PlaceHistory and nil as error.
}

func awardBadge(ctx context.Context, firestoreClient *firestore.Client, userID string, badge Badge) {
	userDocRef := firestoreClient.Collection("users").Doc(userID)

	// Use Firestore transactions to read-modify-write to ensure data integrity
	err := firestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(userDocRef) // Attempt to retrieve the current state of the user document
		if err != nil {
			return err
		}

		var badges []Badge
		if err := doc.DataTo(&badges); err != nil {
			// If the badges field doesn't exist or has a different type, initialize it as an empty slice
			badges = make([]Badge, 0)
		}

		// Check if the badge is already awarded
		for _, b := range badges {
			if b.Name == badge.Name {
				// Badge already awarded, no action needed
				return nil
			}
		}

		// Add the new badge to the slice of badges
		updatedBadges := append(badges, badge)

		// Update the user document with the new slice of badges
		return tx.Set(userDocRef, map[string]interface{}{
			"badges": updatedBadges,
		}, firestore.MergeAll)
	})

	if err != nil {
		log.Printf("Failed to award badge %s to user %s: %v", badge.Name, userID, err)
	} else {
		fmt.Printf("Successfully awarded badge %s to user %s\n", badge.Name, userID)
	}
}

func fetchAllUsers(ctx context.Context, firestoreClient *firestore.Client) ([]User, error) {
	var users []User

	// Query the 'users' collection
	iter := firestoreClient.Collection("users").Documents(ctx)
	defer iter.Stop()

	for {
		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break // Exit the loop when all documents have been processed
			}
			log.Printf("Failed to iterate through users collection: %v", err)
			return nil, err // Return an error if unable to iterate through documents
		}

		var user User
		// Assuming the document data can be directly mapped to the User struct
		if err := doc.DataTo(&user); err != nil {
			log.Printf("Failed to decode document to User struct: %v", err)
			continue // Optionally handle the error differently
		}

		// Set UserID from the document ID
		user.UserID = doc.Ref.ID

		// Add the constructed User object to the slice
		users = append(users, user)
	}

	return users, nil // Return the slice of User objects and nil as the error
}

// fetchAllUserIDs queries the Firestore 'users' collection and retrieves all user IDs.
// func fetchAllUserIDs(ctx context.Context, firestoreClient *firestore.Client) ([]string, error) {
// 	var userIDs []string

// 	// Query the 'users' collection
// 	iter := firestoreClient.Collection("users").Documents(ctx)
// 	defer iter.Stop()

// 	for {
// 		doc, err := iter.Next()
// 		if err != nil {
// 			if err == iterator.Done {
// 				break // Exit the loop when all documents have been processed
// 			}
// 			log.Printf("Failed to iterate through users collection: %v", err)
// 			return nil, err // Return an error if unable to iterate through documents
// 		}

// 		// Add the document ID (userID) to the slice
// 		userIDs = append(userIDs, doc.Ref.ID)
// 	}

// 	return userIDs, nil // Return the slice of userIDs and nil as the error
// }

func contains(slice []string, item string) bool {
	for _, sliceItem := range slice {
		if item == strings.ToLower(sliceItem) {
			return true
		}
	}
	return false
}

func main() {
	ctx := context.Background()
	client, err := connectFirestore(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Firestore: %v", err)
	}
	defer client.Close()

	processCityExplorerBadge(ctx, client)
}

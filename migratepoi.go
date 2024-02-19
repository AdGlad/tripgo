package main

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

func updatePlaceHistory(user *User, firestoreClient *firestore.Client) {
	ctx := context.Background()

	for _, country := range user.Countries {
		for _, region := range country.Regions {
			for _, placeHistory := range region.PlaceHistory {
				if placeHistory.ID == "" {
					log.Printf("Skipping placeHistory record due to empty ID")
					continue // Skip this record to avoid the error
				}

				placeHistoryDoc := firestoreClient.Collection("placehistory").Doc(placeHistory.ID)
				_, err := placeHistoryDoc.Set(ctx, placeHistory)
				if err != nil {
					log.Printf("Failed to write placeHistory record with ID %s: %v", placeHistory.ID, err)
				} else {
					log.Printf("Successfully wrote placeHistory record with ID %s to the placeHistory collection.", placeHistory.ID)
				}
			}
		}
	}
}

func main() {
	ctx := context.Background()
	firestoreClient, err := connectFirestore(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to Firestore: %v", err)
	}
	defer firestoreClient.Close()

	// Get a handle for your collection
	iter := firestoreClient.Collection("users").Documents(ctx)
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
		// Pass the firestore client along with the user reference
		updatePlaceHistory(&user, firestoreClient) // Call to update place history records for each user
	}
}

package main

import (
	"context"
	//	"encoding/json"
	"log"

	//        "time"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

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

	//	updateAllUserPlaceHistories(ctx, client)

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

		processAsianExplorerBadge(&user)        // Process Asian Explorer badge
		processGlobetrotterBadge(&user)         // Process Globetrotter badge
		processPolarExplorerBadge(&user)        // Process Polar Explorer badge
		processEuroExplorerBadge(&user)         // Process Polar Explorer badge
		processOzzieExplorerBadge(&user)        // Process Polar Explorer badge
		processCityExplorerBadge(&user, client) // Process Polar Explorer badge

		//Convert user struct to a map for updating
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

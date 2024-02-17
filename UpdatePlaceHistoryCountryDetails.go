package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

// UpdatePlaceHistoryCountryDetails updates place history records with their parent country's code and name,
// but only if the place history's CountryCode is empty.
// UpdatePlaceHistoryCountryDetails updates place history records with their parent country's code and name,
// but only if the place history's CountryCode is empty.

func updateAllUserPlaceHistories(ctx context.Context, client *firestore.Client) error {
	iter := client.Collection("users").Documents(ctx)
	defer iter.Stop()
	log.Printf("updateAllUserPlaceHistories")

	for {

		log.Printf("for")

		doc, err := iter.Next()
		if err != nil {
			if err == iterator.Done {
				break // Properly handle the end of the collection
			}
			return fmt.Errorf("failed to iterate through documents: %v", err)
		}

		userData := doc.Data() // Retrieve the document data as a map
		countries, ok := userData["countries"].(map[string]interface{})
		if !ok {
			log.Printf("Document %s does not contain countries map", doc.Ref.ID)
			continue
		}

		var updates []firestore.Update
		for countryID, countryData := range countries {
//			if doc.Ref.ID == "wPxuLxEvYpY5D0AnSPktZriPIQa2" {

				log.Printf("Countries countryID: %s ", countryID)
//			}
			country, ok := countryData.(map[string]interface{})
			if !ok {
				continue // Skip if the structure is not as expected
			}

			newCountryCode, codeOk := country["countryCode"].(string)
			newCountryName, nameOk := country["countryName"].(string)

			if !codeOk || !nameOk {
				log.Printf("Country %s missing countryCode or countryName", countryID)
				continue
			}

			regions, ok := country["regions"].(map[string]interface{})
			if !ok {
				continue
			}

			for regionID, regionData := range regions {
	//			if doc.Ref.ID == "wPxuLxEvYpY5D0AnSPktZriPIQa2" {

					log.Printf("regions regionID: %s ", regionID)
	//			}
				region, ok := regionData.(map[string]interface{})
				if !ok {
					continue
				}

				placeHistories, ok := region["placehistory"].(map[string]interface{})
				if !ok {
					continue
				}

				for placeHistoryID, placeHistoryData := range placeHistories {
				//	if doc.Ref.ID == "wPxuLxEvYpY5D0AnSPktZriPIQa2" {

						log.Printf("placeHistories placeHistoryID: %s ", placeHistoryID)
				//	}
					placeHistory, ok := placeHistoryData.(map[string]interface{})
					if !ok {
						continue
					}
					// Update only if countryCode is empty
					if placeHistory["countryCode"] == "" {
						pathPrefix := fmt.Sprintf("countries.%s.regions.%s.placehistory.%s", countryID, regionID, placeHistoryID)
						//log.Printf("Updating user %s: pathPrefix: %s, newCountryCode: %s, newCountryName: %s", doc.Ref.ID, pathPrefix, newCountryCode, newCountryName)
				//		if doc.Ref.ID == "RaeLS1r88CbkCEiA8CqUBUjKEzJ3" {
							log.Printf("Updating user %s: pathPrefix: %s, newCountryCode: %s, newCountryName: %s", doc.Ref.ID, pathPrefix, newCountryCode, newCountryName)

							updates = append(updates,
								firestore.Update{Path: pathPrefix + ".countryCode", Value: newCountryCode},
								firestore.Update{Path: pathPrefix + ".countryName", Value: newCountryName},
							)
				//		}
					}
				}
			}
		}

		// Apply the updates if there are any
		if len(updates) > 0 {
			_, err = doc.Ref.Update(ctx, updates)
			if err != nil {
				return fmt.Errorf("failed to update document %s: %v", doc.Ref.ID, err)
			}
			log.Printf("Successfully updated document %s", doc.Ref.ID)
		}
	}

	return nil // Return nil to indicate success if no errors occurred
}

// func updateAllUserPlaceHistories(ctx context.Context, client *firestore.Client) error {
// 	iter := client.Collection("users").Documents(ctx)
// 	defer iter.Stop()

// 	for {
// 		doc, err := iter.Next()
// 		if err != nil {
// 			if err == iterator.Done {
// 				break // Properly handle the end of the collection
// 			}
// 			return fmt.Errorf("failed to iterate through documents: %v", err)
// 		}

// 		userData := doc.Data() // Correct method to retrieve the document data as a map
// 		countries, ok := userData["countries"].(map[string]interface{})
// 		if !ok {
// 			log.Printf("Document %s does not contain countries map", doc.Ref.ID)
// 			continue
// 		}

// 		var updates []firestore.Update
// 		for countryID, countryData := range countries {
// 			country, ok := countryData.(map[string]interface{})
// 			if !ok {
// 				continue // Skip if the structure is not as expected
// 			}

// 			// Assuming country name and code are direct fields of the country
// 			newCountryCode, codeOk := country["countryCode"].(string)
// 			newCountryName, nameOk := country["countryName"].(string) // Adjust if the country name is stored under a different field

// 			// Proceed only if both newCountryCode and newCountryName are available
// 			if !codeOk || !nameOk {
// 				log.Printf("Country %s missing countryCode or countryName", countryID)
// 				continue
// 			}

// 			regions, ok := country["regions"].(map[string]interface{})
// 			if !ok {
// 				continue
// 			}

// 			for regionID, regionData := range regions {
// 				region, ok := regionData.(map[string]interface{})
// 				if !ok {
// 					continue
// 				}

// 				placeHistories, ok := region["placehistory"].(map[string]interface{})
// 				if !ok {
// 					continue
// 				}

// 				for placeHistoryID, _ := range placeHistories {
// 					// Constructing the path for updating each placeHistory record
// 					pathPrefix := fmt.Sprintf("countries.%s.regions.%s.placehistory.%s", countryID, regionID, placeHistoryID)
// 					log.Printf("countries.%s.regions.%s.placehistory.%s", countryID, regionID, placeHistoryID)
// 					log.Printf("pathPrefix: %s newCountryCode: %s, newCountryName: %s", pathPrefix, newCountryCode, newCountryName)

// 					// updates = append(updates,
// 					// 	firestore.Update{Path: pathPrefix + ".countryCode", Value: newCountryCode},
// 					// 	firestore.Update{Path: pathPrefix + ".countryName", Value: newCountryName},
// 					// )
// 				}
// 			}
// 		}

// 		// Apply the updates if there are any
// 		if len(updates) > 0 {
// 			_, err = doc.Ref.Update(ctx, updates)
// 			if err != nil {
// 				log.Printf("Failed to update document %s: %v", doc.Ref.ID, err)
// 			} else {
// 				log.Printf("Successfully updated document %s", doc.Ref.ID)
// 			}
// 		}
// 	}

// 	return nil // Return nil to indicate success if no errors occurred
// }

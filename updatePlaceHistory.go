package main

//    "time"

// updatePlaceHistory reads all place history records for a user and writes them to the placeHistory collection.
// func updatePlaceHistory(user *User, firestoreClient *firestore.Client) {
//     ctx := context.Background()

//     for _, country := range user.Countries {
//         for _, region := range country.Regions {
//             for _, placeHistory := range region.PlaceHistory {
//                 if placeHistory.ID == "" {
//                     log.Printf("Skipping placeHistory record due to empty ID")
//                     continue // Skip this record to avoid the error
//                 }
//                 // Each placeHistory record becomes its own document in the placeHistory collection.
//                 placeHistoryDoc := firestoreClient.Collection("placehistory").Doc(placeHistory.ID)
//                 _, err := placeHistoryDoc.Set(ctx, placeHistory)
//                 if err != nil {
//                     log.Printf("Failed to write placeHistory record with ID %s: %v", placeHistory.ID, err)
//                 } else {
//                     log.Printf("Successfully wrote placeHistory record with ID %s to the placeHistory collection.", placeHistory.ID)
//                 }
//             }
//         }
//     }
// }

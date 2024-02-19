// firestore.go
package main

import (
    "context"
    "log"

    firebase "firebase.google.com/go"
    "cloud.google.com/go/firestore"
    "google.golang.org/api/option"
)

func connectFirestore(ctx context.Context) (*firestore.Client, error) {
	sa := option.WithCredentialsFile("/Users/adglad/Downloads/triparific100-fa4917b5dd13.json") // Replace with your JSON file path
	app, err := firebase.NewApp(ctx, nil, sa)
	if err != nil {
           log.Fatalf("Failed to create Firestore app: %v", err)

		return nil, err
	}

	client, err := app.Firestore(ctx)
	if err != nil {
                log.Fatalf("Failed to create Firestore client: %v", err)
		return nil, err
	}
	return client, nil
}

//func connectFirestoreProj(ctx context.Context) (*firestore.Client, error) {
	//// Update the path to your Firebase service account key
	//sa := option.WithCredentialsFile("/Users/adglad/Downloads/triparific100-fa4917b5dd13.json")
	//app, err := firestore.NewClient(ctx, "triparific100", sa)
	//if err != nil {
		//log.Fatalf("Failed to create Firestore client: %v", err)
		//return nil, err
	//}
	//return app, nil
//}

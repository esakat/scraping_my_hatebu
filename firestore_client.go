package main

import (
	"cloud.google.com/go/firestore"
	"context"
	. "github.com/esakat/observe_my_hatebu/data"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

var ctx = context.Background()
var firestoreClient *firestore.Client

func createFirestoreClient() {
	var err error
	firestoreClient, err = firestore.NewClient(ctx, config.ProjectID)
	if err != nil {
		log.Fatalf("failed to create client: %v", err)
	}
}

func HasEid(eid string) bool {
	ref, err := firestoreClient.Collection(config.EntryCollectionName).Doc(eid).Get(ctx)
	if err != nil && status.Code(err) != codes.NotFound {
		log.Fatalf("failed to get %s: %v", eid, err)
	}
	return ref.Exists()
}

func AddMyHatebuEntry(entry *MyEntry) {
	_, err := firestoreClient.Collection(config.EntryCollectionName).Doc(entry.EntryID).Set(ctx, entry)
	if err != nil {
		log.Fatalf("failed to add: %v", err)
	}
}

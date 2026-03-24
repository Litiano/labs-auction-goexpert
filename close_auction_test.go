package labs_auction_goexpert

import (
	"context"
	"fullcycle-auction_go/configuration/database/mongodb"
	"fullcycle-auction_go/internal/entity/auction_entity"
	auction2 "fullcycle-auction_go/internal/infra/database/auction"
	"log"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestCloseAuction(t *testing.T) {
	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/testing.env"); err != nil {
		log.Fatal("Error trying to load env variables")
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		t.Fatal(err)
	}

	auction, internalError := auction_entity.CreateAuction("Product name", "Product category", "Product description", auction_entity.ProductCondition(1))
	if internalError != nil {
		t.Errorf("error creating auction: %v", internalError)
	}

	repository := auction2.NewAuctionRepository(databaseConnection)
	repository.CreateAuction(ctx, auction)

	createdAuction, internalError := repository.FindAuctionById(ctx, auction.Id)
	if internalError != nil {
		t.Errorf("error finding auction: %v", internalError)
	}

	if createdAuction.Status != auction_entity.AuctionStatus(0) {
		t.Errorf("auction status should be 0 but is %v", createdAuction.Status)
	}

	time.Sleep(4 * time.Second)

	createdAuction, internalError = repository.FindAuctionById(ctx, auction.Id)
	if internalError != nil {
		t.Errorf("error finding auction: %v", internalError)
	}

	if createdAuction.Status != auction_entity.AuctionStatus(1) {
		t.Errorf("auction status should be 1 but is %v", createdAuction.Status)
	}
}

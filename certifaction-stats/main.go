package main

import "math/big"

type Event struct {
	FileHash string
	ClaimHash string
	BlockNumber *big.Int
}

type Repository interface {
	Create(event Event)
	FindAll() []Event
	LastBlock() *big.Int
}

type Certifaction interface {
	EventsByBlocks(blockStart, blockEnd *big.Int)
	Listen() chan Event
}

func main() {
	// ON INIT:
	// Retrieve last block stored (first time would be 0 as we don't have cached events)
	var repository Repository
	var certifaction Certifaction
	// Listen for all new events, those should be stored async
	eventsChan := certifaction.Listen()
	// Check last current block on Ethereum
	lastBlockStored := repository.LastBlock()

	// Get all events up to the block we had in our DB and before starting to listen for new events
	events := certifaction.EventsByBlocks(lastBlockStored, nil)
	// Store all events

}

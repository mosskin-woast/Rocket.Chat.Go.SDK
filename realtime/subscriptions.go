package realtime

import (
	"fmt"
	"log"

	"github.com/gopackage/ddp"
)

// Subscribes to stream-notify-logged
// Returns a buffered channel
//
// https://rocket.chat/docs/developer-guides/realtime-api/subscriptions/stream-room-messages/
func (c *Client) Sub(name string, args ...interface{}) (chan string, error) {

	if args == nil {
		log.Println("no args passed")
		if err := c.ddp.Sub(name); err != nil {
			return nil, err
		}
	} else {
		if err := c.ddp.Sub(name, args[0], false); err != nil {
			return nil, err
		}
	}

	msgChannel := make(chan string, default_buffer_size)
	c.ddp.CollectionByName("stream-room-messages").AddUpdateListener(genericExtractor{msgChannel, "update"})

	return msgChannel, nil
}

// RoomMessage is a room message
type RoomMessage struct {
}

// SubStreamRoomMessages subscribes to stream-room-messages on an existing channel
func (c *Client) SubStreamRoomMessages(roomID string, msgChannel chan RoomMessage) error {
	streamName := "stream-room-messages"
	err := c.ddp.Sub(streamName, roomID, false)
	if err != nil {
		log.Println("### Sub err: ", err)
		return err
	}
	c.ddp.CollectionByName(streamName).AddUpdateListener(roomMessageExtractor{msgChannel, "update"})
	log.Println("SUBSCRIBED!", roomID)
	return nil
}

type genericExtractor struct {
	messageChannel chan string
	operation      string
}

func (u genericExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	if operation == u.operation {
		u.messageChannel <- fmt.Sprintf("%s -> update", collection)
	}
}

type roomMessageExtractor struct {
	messageChannel chan RoomMessage
	operation      string
}

func (u roomMessageExtractor) CollectionUpdate(collection, operation, id string, doc ddp.Update) {
	log.Printf("### OPERATION: %s\n", operation)
	if operation != u.operation {
		return
	}

	log.Printf("#### GOT MESSAGE: %+v\n", doc)
}

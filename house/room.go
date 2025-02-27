package house

import (
	"github.com/leeif/mercury/storage/data"
)

type Room struct {
	data.RoomBase
	receivceMember chan *Member
	receiveMessage chan *Message
}

func (room *Room) ReceiveMessage() {
	for {
		select {
		case message, ok := <-room.receiveMessage:
			if ok {
				room.TransferMessage(message)
			}
		}
	}
}

func (room *Room) ReceiveMember() {
	for {
		select {
		case member, ok := <-room.receivceMember:
			if ok {
				room.TransferUnReadMessage(member)
			}
		}
	}
}

func (room *Room) TransferMessage(message *Message) {
	mid := house.Store.GetMemberFromRoom(room.ID)
	entries := house.Store.GetMember(mid...)
	members := make([]*Member, len(entries))
	for i := range entries {
		members[i] = entries[i].(*Member)
	}
	for _, member := range members {
		if member != nil {
			if !member.isClosed {
				if b, err := message.json(); err == nil {
					member.conn.Send <- b
					// position increament, should be locked in the furture
					house.Store.SetRoomMemberMessage(room.ID, member.ID, message.ID)
				}
			}
		}
	}
}

func (room *Room) TransferUnReadMessage(member *Member) {
	msg_id := house.Store.GetRoomMemberMessage(room.ID, member.ID)
	messages := house.Store.GetUnReadMessage(room.ID, msg_id)
	if messages == nil {
		return
	}
	for _, v := range messages {
		message := &Message{MessageBase: *v}
		if b, err := message.json(); err == nil && !member.isClosed {
			member.conn.Send <- b
			// position increament, should be locked in the furture
			house.Store.SetRoomMemberMessage(room.ID, member.ID, v.ID)
		}
	}
}

func (room *Room) Work() {
	go room.ReceiveMessage()
	go room.ReceiveMember()
}

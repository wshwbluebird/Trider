package data

import "fmt"

type RoomCtrip struct {
	HotelName string
	HotelID string
	RoomType string
	RoomTypeName string
	RoomBed string
	RoomWIFI string
	RoomAmount string
	RoomPolicy string
	RoomPrice string
}

func NewEmptyRoomCtrip() *RoomCtrip{
	return &RoomCtrip{}
}

func (room *RoomCtrip) toString() string{
	return fmt.Sprintf("HotelName: %s\nHotelID: %s\n" +
		"RoomType: %s\nRoomTypeName: %s\n" +
		"RoomBed: %s\nRoomWIFI: %s\n" +
		"RoomAmount: %s\nRoomPolicy: %s\n" +
		"RoomPrice: %s",
		room.HotelName,room.HotelID,
		room.RoomType,room.RoomTypeName,
		room.RoomBed,room.RoomWIFI,
		room.RoomAmount,room.RoomPolicy,
			room.RoomPrice)
}
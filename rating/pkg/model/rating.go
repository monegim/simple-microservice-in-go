package model

type RecordID string

type RecordType string

const RecordTypeMovie = RecordType("Movie")

type UserID string

type RatingValue int

type Rating struct {
	RecordID   RecordID    `json:"recordID"`
	RecordType RecordType  `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}

package main

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
)

// Contact holds contact information for a single person in an address book.
type Contact struct {
	ID bson.ObjectId `bson:"_id" json:"id"`

	FirstName string   `bson:"firstName" json:"firstName"`
	LastName  string   `bson:"lastName" json:"lastName"`
	Company   string   `bson:"company" json:"company"`
	Title     string   `bson:"title" json:"title"`
	Roles     []string `bson:"roles" json:"roles"`
}

// NullString is a string value that may be null. This allows us to honor the
// contract of RFC 7396 in that sending a request like
//	"title": null
// Should unset the value of "title" rather than requiring them to send ""
type NullString struct {
	Present bool
	Value   string
}

// UnmarshalJSON implements json.Unmarshaler. Simply being called will set the
// Present field to true which would otherwise default to false.
func (ns *NullString) UnmarshalJSON(raw []byte) error {
	ns.Present = true

	// Only accepts values null or some string
	var val *string
	if err := json.Unmarshal(raw, &val); err != nil {
		return err
	}

	// Set ns.Value if provided. If val is nil we leave ns.Value the default ""
	if val != nil {
		ns.Value = *val
	}

	return nil
}

// UpdateContact is what we accept from clients to update a Contact record.
// Only fields that are provided are updated.
type UpdateContact struct {
	FirstName NullString `json:"firstName"`
	LastName  NullString `json:"lastName"`
	Company   NullString `json:"company"`
	Title     NullString `json:"title"`
	Roles     []string   `json:"roles"`
}

// Changes generates a $set operation for mongo that sets just the fields
// necessary to update a contact.
// TODO could accept a flContext.Values to generate update metadata
func (u UpdateContact) Changes() bson.M {
	fields := make(bson.M)

	if u.FirstName.Present {
		fields["firstName"] = u.FirstName.Value
	}
	if u.LastName.Present {
		fields["lastName"] = u.LastName.Value
	}
	if u.Company.Present {
		fields["company"] = u.Company.Value
	}
	if u.Title.Present {
		fields["title"] = u.Title.Value
	}
	if u.Roles != nil {
		fields["roles"] = u.Roles
	}

	return bson.M{"$set": fields}
}

package database

import (
	"SoftwareGoDay2/ent"
	"context"

	"github.com/google/uuid"
)

func (d Database) CreateContact(ctx context.Context, artistID uuid.UUID, email, phone string) (*ent.Contact, error) {
	// Create a new contact with the provided email, phone number, and artist ID
	contact, err := d.Client.Contact.Create().SetEmail(email).SetPhone(phone).SetArtistID(artistID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return contact, nil
}

func (d Database) UpdateContact(ctx context.Context, contact *ent.Contact) (*ent.Contact, error) {
	if contact != nil {
		// Update the contact with the provided data
		err := d.Client.Contact.UpdateOne(contact).Exec(ctx)
		if err != nil {
			return nil, err
		}
	}

	return contact, nil
}

func (d Database) DeleteContact(ctx context.Context, id uuid.UUID) (*ent.Contact, error) {
	// Get the contact with the provided ID
	contact, err := d.Client.Contact.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Delete the contact from the database
	err = d.Client.Contact.DeleteOne(contact).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return contact, nil
}
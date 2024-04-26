package controller

import (
	"SoftwareGoDay2/ent"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (c Controller) CreateContact(ctx context.Context, artistID string, email, phone string) (*ent.Contact, error) {
	// Parse the artist ID string into a UUID and check if it's valid
	id, err := uuid.Parse(artistID)
	if err != nil {
		return nil, errors.New("artist id is invalid")
	}

	// Check if the email is valid
	if email == "" {
		return nil, errors.New("contact email is invalid")
	}

	// Check if the phone is valid
	if phone == "" {
		return nil, errors.New("contact phone is invalid")
	}

	// Create the contact in the database
	contact, err := c.Database.CreateContact(ctx, id, email, phone)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create contact")
	}

	return contact, nil
}

func (c Controller) GetContacts(ctx context.Context) ([]*ent.Contact, error) {
	// Get all contacts
	artists, err := c.Database.GetArtists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contacts")
	}

	// Get the contacts from the artists
	var contacts []*ent.Contact
	for _, artist := range artists {
		contacts = append(contacts, artist.Edges.Contact)
	}

	return contacts, nil
}

func (c Controller) GetContactByID(ctx context.Context, id string) (*ent.Contact, error) {
	// Parse the ID string into a UUID and check if it's valid
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("contact id is invalid")
	}

	// Get the artist by ID
	artist, err := c.Database.GetArtistByID(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contact by ID")
	}

	// Get the contact from the artist
	contact := artist.Edges.Contact

	return contact, nil
}

func (c Controller) UpdateContact(ctx context.Context, id, email, phone string) (*ent.Contact, error) {
	// Get the contact by ID
	contact, err := c.GetContactByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get contact by ID")
	}

	if contact != nil {
		// Check if the email is valid and update it if it's different
		if email != "" && email != contact.Email {
			contact.Email = email
		}

		// Check if the phone is valid and update it if it's different
		if phone != "" && phone != contact.Phone {
			contact.Phone = phone
		}
	}

	// Update the contact
	contact, err = c.Database.UpdateContact(ctx, contact)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update contact")
	}

	return contact, nil
}

func (c Controller) DeleteContact(ctx context.Context, id string) (*ent.Contact, error) {
	// Parse the ID string into a UUID and check if it's valid
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("contact id is invalid")
	}

	// Delete the contact by ID
	contact, err := c.Database.DeleteContact(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete contact")
	}

	return contact, nil
}
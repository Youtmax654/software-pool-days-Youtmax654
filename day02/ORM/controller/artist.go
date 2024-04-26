package controller

import (
	"SoftwareGoDay2/ent"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)


var InvalidArtistName = errors.New("artist name is invalid")

func (c Controller) CreateArtist(ctx context.Context, name string, nationality string) (*ent.Artist, error) {
	// Check if the name is empty
	if name == "" {
		return nil, InvalidArtistName
	}

	// Create the artist in the database
	artist, err := c.Database.CreateArtist(ctx, name, nationality)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create artist")
	}
	
	return artist, nil
}

func (c Controller) GetArtists(ctx context.Context) ([]*ent.Artist, error) {
	// Get all artists
	artists, err := c.Database.GetArtists(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get artists")
	}

	return artists, nil
}

func (c Controller) GetArtistByID(ctx context.Context, id string) (*ent.Artist, error) {
	// Parse the ID string into a UUID and check if it's valid
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("artist id is invalid")
	}

	// Get the artist by ID
	artist, err := c.Database.GetArtistByID(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get artist by ID")
	}

	return artist, nil
}

func (c Controller) UpdateArtist(ctx context.Context, id, name, nationality string) (*ent.Artist, error) {
	// Get the artist by ID
	artist, err := c.GetArtistByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get artist by ID")
	}

	// Check if the name is valid and update it if it's different
	if name != "" && name != artist.Name {
		artist.Name = name
	}

	// Check if the nationality is valid and update it if it's different
	if nationality != "" && nationality != artist.Nationality {
		artist.Nationality = nationality
	}

	// Update the artist
	artist, err = c.Database.UpdateArtist(ctx, artist)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update artist")
	}

	return artist, nil
}

func (c Controller) DeleteArtist(ctx context.Context, id string) (*ent.Artist, error) {
	// Parse the ID string into a UUID and check if it's valid
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("artist id is invalid")
	}

	// Delete the artist by ID
	artist, err := c.Database.DeleteArtist(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete artist")
	}

	return artist, nil
}
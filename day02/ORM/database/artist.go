package database

import (
	"SoftwareGoDay2/ent"
	"context"

	"github.com/google/uuid"
)

func (d Database) CreateArtist(ctx context.Context, name, nationality string) (*ent.Artist, error) {
	// Create a new artist with the provided
	artist, err := d.Client.Artist.Create().SetName(name).SetNationality(nationality).Save(ctx)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (d Database) GetArtists(ctx context.Context) ([]*ent.Artist, error) {
	// Get all artists from the database
	artists, err := d.Client.Artist.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (d Database) GetArtistByID(ctx context.Context, id uuid.UUID) (*ent.Artist, error) {
	// Get the artist with the provided ID
	artist, err := d.Client.Artist.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (d Database) UpdateArtist(ctx context.Context, artist *ent.Artist) (*ent.Artist, error) {
	// Update the artist with the provided data
	err := d.Client.Artist.UpdateOne(artist).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (d Database) DeleteArtist(ctx context.Context, id uuid.UUID) (*ent.Artist, error) {
	// Get the artist with the provided ID
	artist, err := d.Client.Artist.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Delete the artist from the database
	err = d.Client.Artist.DeleteOne(artist).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
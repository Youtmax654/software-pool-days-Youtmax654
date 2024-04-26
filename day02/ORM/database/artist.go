package database

import (
	"SoftwareGoDay2/ent"
	"context"

	"github.com/google/uuid"
)

func (d Database) CreateArtist(ctx context.Context, name, nationality string) (*ent.Artist, error) {
	artist, err := d.Client.Artist.Create().SetName(name).SetNationality(nationality).Save(ctx)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (d Database) GetArtists(ctx context.Context) ([]*ent.Artist, error) {
	artists, err := d.Client.Artist.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return artists, nil
}

func (d Database) GetArtistByID(ctx context.Context, id uuid.UUID) (*ent.Artist, error) {
	artist, err := d.Client.Artist.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (d Database) UpdateArtist(ctx context.Context, artist *ent.Artist) (*ent.Artist, error) {
	err := d.Client.Artist.UpdateOne(artist).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return artist, nil
}

func (d Database) DeleteArtist(ctx context.Context, id uuid.UUID) (*ent.Artist, error) {
	artist, err := d.Client.Artist.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	err = d.Client.Artist.DeleteOne(artist).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return artist, nil
}
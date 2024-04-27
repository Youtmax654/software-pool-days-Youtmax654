package database

import (
	"SoftwareGoDay2/ent"
	"context"

	"github.com/google/uuid"
)

func (d Database) CreateRecordCompany(ctx context.Context, name string) (*ent.RecordCompany, error) {
	// Create the record company in the database
	recordCompany, err := d.Client.RecordCompany.Create().SetName(name).Save(ctx)
	if err != nil {
		return nil, err
	}

	return recordCompany, nil
}

func (d Database) GetRecordCompanies(ctx context.Context) ([]*ent.RecordCompany, error) {
	// Get all record companies from the database
	recordCompanies, err := d.Client.RecordCompany.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	return recordCompanies, nil
}

func (d Database) GetRecordCompanyByID(ctx context.Context, id uuid.UUID) (*ent.RecordCompany, error) {
	// Get the record company by ID from the database
	recordCompany, err := d.Client.RecordCompany.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return recordCompany, nil
}

func (d Database) UpdateRecordCompany(ctx context.Context, recordCompany *ent.RecordCompany) (*ent.RecordCompany, error) {
	// Update the record company in the database
	err := d.Client.RecordCompany.UpdateOne(recordCompany).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return recordCompany, nil
}

func (d Database) DeleteRecordCompany(ctx context.Context, id uuid.UUID) (*ent.RecordCompany, error) {
	// Get the record company by ID from the database
	recordCompany, err := d.Client.RecordCompany.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	// Delete the record company from the database
	err = d.Client.RecordCompany.DeleteOne(recordCompany).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return recordCompany, nil
}

func (d Database) AddArtistToRecordCompany(ctx context.Context, artistID uuid.UUID, recordCompanyID uuid.UUID) (*ent.RecordCompany, error) {
	// Add the artist to the record company in the database
	recordCompany, err := d.Client.RecordCompany.UpdateOneID(recordCompanyID).AddArtistIDs(artistID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return recordCompany, nil
}

func (d Database) RemoveArtistFromRecordCompany(ctx context.Context, artistID uuid.UUID, recordCompanyID uuid.UUID) (*ent.RecordCompany, error) {
	// Remove the artist from the record company in the database
	recordCompany, err := d.Client.RecordCompany.UpdateOneID(recordCompanyID).RemoveArtistIDs(artistID).Save(ctx)
	if err != nil {
		return nil, err
	}

	return recordCompany, nil
}
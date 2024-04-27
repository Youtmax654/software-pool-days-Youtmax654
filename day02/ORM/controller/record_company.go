package controller

import (
	"SoftwareGoDay2/ent"
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

func (c Controller) CreateRecordCompany(ctx context.Context, name string) (*ent.RecordCompany, error) {
	// Check if the name is valid
	if name == "" {
		return nil, errors.New("record company name is invalid")
	}

	// Create the record company in the database
	recordCompany, err := c.Database.CreateRecordCompany(ctx, name)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create record company")
	}

	return recordCompany, nil
}

func (c Controller) GetRecordCompanies(ctx context.Context) ([]*ent.RecordCompany, error) {
	// Get all record companies
	recordCompanies, err := c.Database.GetRecordCompanies(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get record companies")
	}

	return recordCompanies, nil
}

func (c Controller) GetRecordCompanyByID(ctx context.Context, id string) (*ent.RecordCompany, error) {
	// Parse the ID string into a UUID and check if it's valid
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("record company id is invalid")
	}

	// Get the record company by ID
	recordCompany, err := c.Database.GetRecordCompanyByID(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get record company by id")
	}

	return recordCompany, nil
}

func (c Controller) UpdateRecordCompany(ctx context.Context, id string, name string) (*ent.RecordCompany, error) {
	// Get the record company by ID
	recordCompany, err := c.GetRecordCompanyByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Check if the name is valid and update it if it's different
	if name != "" && name != recordCompany.Name {
		recordCompany.Name = name
	}

	// Update the record company in the database
	recordCompany, err = c.Database.UpdateRecordCompany(ctx, recordCompany)
	if err != nil {
		return nil, errors.Wrap(err, "failed to update record company")
	}

	return recordCompany, nil
}

func (c Controller) DeleteRecordCompany(ctx context.Context, id string) (*ent.RecordCompany, error) {
	// Parse the ID string into a UUID and check if it's valid
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("record company id is invalid")
	}

	// Delete the record company from the database
	recordCompany, err := c.Database.DeleteRecordCompany(ctx, uid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to delete record company")
	}

	return recordCompany, nil
}

func (c Controller) AddArtistToRecordCompany(ctx context.Context, artistID, recordCompanyID string) (*ent.RecordCompany, error) {
	// Parse the artist ID string into a UUID and check if it's valid
	aid, err := uuid.Parse(artistID)
	if err != nil {
		return nil, errors.New("artist id is invalid")
	}

	// Parse the record company ID string into a UUID and check if it's valid
	rid, err := uuid.Parse(recordCompanyID)
	if err != nil {
		return nil, errors.New("record company id is invalid")
	}

	// Add the artist to the record company in the database
	recordCompany, err := c.Database.AddArtistToRecordCompany(ctx, aid, rid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to add artist to record company")
	}

	return recordCompany, nil
}

func (c Controller) RemoveArtistFromRecordCompany(ctx context.Context, artistID, recordCompanyID string) (*ent.RecordCompany, error) {
	// Parse the artist ID string into a UUID and check if it's valid
	aid, err := uuid.Parse(artistID)
	if err != nil {
		return nil, errors.New("artist id is invalid")
	}

	// Parse the record company ID string into a UUID and check if it's valid
	rid, err := uuid.Parse(recordCompanyID)
	if err != nil {
		return nil, errors.New("record company id is invalid")
	}

	// Remove the artist from the record company in the database
	recordCompany, err := c.Database.RemoveArtistFromRecordCompany(ctx, aid, rid)
	if err != nil {
		return nil, errors.Wrap(err, "failed to remove artist from record company")
	}

	return recordCompany, nil
}
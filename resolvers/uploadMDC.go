package resolvers

import (
	"context"
	"fmt"

	"github.com/3dw1nM0535/deli/db/models"
	models1 "github.com/3dw1nM0535/deli/models"
	"github.com/3dw1nM0535/deli/utils"
)

func (r *mutationResolver) UploadMc(ctx context.Context, input models1.UploadMc) (*models1.File, error) {
	ctx = context.Background()

	// validate rider is registered
	rider := &models.Rider{}
	r.ORM.DB.First(&rider, "id = ?", utils.ParseUUID(input.RiderID))
	if rider.ID.String() == "00000000-0000-0000-0000-000000000000" {
		err := fmt.Errorf("no rider with id '%s' is registered with Deli", input.RiderID)
		return &models1.File{}, err
	}

	_, attr, err := utils.Upload(ctx, input.File.File, riderIDBucketName, credPath, projectID, input.File.Filename)
	if err != nil {
		return &models1.File{}, err
	}

	medicalCert := &models.MDC{
		Media:     attr.MediaLink,
		Content:   attr.ContentType,
		Size:      int(attr.Size),
		CreatedAt: attr.Created,
		UpdatedAt: attr.Updated,
		RiderID:   utils.ParseUUID(input.RiderID),
	}
	r.ORM.DB.Save(&medicalCert)
	return &models1.File{
		ID:        medicalCert.ID.String(),
		Media:     medicalCert.Media,
		Content:   medicalCert.Content,
		Size:      medicalCert.Size,
		CreatedAt: &medicalCert.CreatedAt,
		UpdatedAt: &medicalCert.UpdatedAt,
	}, nil
}
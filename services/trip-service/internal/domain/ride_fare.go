package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ride-sharing/services/trip-service/pkg/types"
)

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       string // ex: van, luxury, sedan
	TotalPriceInCents float64
	Route             *types.OsrmApiResponse
}

package ff

import (
	ff_mongo_exceptions "github.com/GabrielEstmr/ff-4-go/ff/mongo/exceptions"
	ff_resources "github.com/GabrielEstmr/ff-4-go/ff/resources"
)

type FeaturesMethods interface {
	IsEnabled(key string) (bool, ff_mongo_exceptions.LibException)
	IsDisabled(key string) (bool, ff_mongo_exceptions.LibException)
	Enable(key string) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException)
	Disable(key string) (ff_resources.FeaturesData, ff_mongo_exceptions.LibException)
}

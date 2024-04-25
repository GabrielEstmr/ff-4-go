package ff_factories

import (
	"errors"
	"github.com/GabrielEstmr/ff-4-go/ff"
	ff_mongo "github.com/GabrielEstmr/ff-4-go/ff/mongo"
)

type RegisterMethodsFactory struct {
	ffConfigData *ff.FfConfigData
}

func NewRegisterMethodsFactory(ffConfigData *ff.FfConfigData) *RegisterMethodsFactory {
	return &RegisterMethodsFactory{ffConfigData}
}

func (this *RegisterMethodsFactory) Get() (ff.RegisterMethods, error) {
	if this.ffConfigData.GetClientType() == ff.MONGO {
		return ff_mongo.NewRegisterMethodsMongoImpl(this.ffConfigData), nil
	}
	return nil, errors.New("could not instantiate a valid FeaturesData")
}

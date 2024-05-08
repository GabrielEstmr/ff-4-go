package ff_gateways_cacheddb

import (
	ff_domains "github.com/GabrielEstmr/ff-4-go/ff/domains"
	ff_domains_exceptions "github.com/GabrielEstmr/ff-4-go/ff/domains/exceptions"
	ff_gateways "github.com/GabrielEstmr/ff-4-go/ff/gateways"
	"log"
)

type RolloutCachedMongoGatewayImpl struct {
	gateway      ff_gateways.RolloutGateway
	cacheGateway ff_gateways.RolloutGateway
}

func NewRolloutCachedMongoGatewayImpl(
	gateway ff_gateways.RolloutGateway,
	cacheGateway ff_gateways.RolloutGateway,
) *RolloutCachedMongoGatewayImpl {
	return &RolloutCachedMongoGatewayImpl{
		gateway:      gateway,
		cacheGateway: cacheGateway,
	}
}

func (this *RolloutCachedMongoGatewayImpl) Save(
	rollout ff_domains.Rollout,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	savedRolloutDoc, err := this.gateway.Save(rollout)
	if err != nil {
		return *new(ff_domains.Rollout), err
	}

	go func() {
		_, err := this.cacheGateway.Save(savedRolloutDoc)
		if err != nil {
			log.Println("ff-4-go error: error to save in cache gateway. err:", err.Error())
		}
	}()
	return savedRolloutDoc, nil
}

func (this *RolloutCachedMongoGatewayImpl) Update(
	rollout ff_domains.Rollout,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	_, err := this.cacheGateway.Update(rollout)
	if err != nil {
		log.Println("ff-4-go error: error to update in cache gateway. err:", err.Error())
	}

	updatedRolloutDoc, err := this.gateway.Update(rollout)
	if err != nil {
		return *new(ff_domains.Rollout), err
	}

	return updatedRolloutDoc, nil
}

func (this *RolloutCachedMongoGatewayImpl) Delete(
	key string,
) ff_domains_exceptions.LibException {

	errC := this.cacheGateway.Delete(key)
	if errC != nil {
		return errC
	}

	errDb := this.gateway.Delete(key)
	if errDb != nil {
		return errDb
	}

	return nil
}

func (this *RolloutCachedMongoGatewayImpl) FindById(
	key string,
) (ff_domains.Rollout, ff_domains_exceptions.LibException) {

	rollout, errC := this.cacheGateway.FindById(key)
	if errC == nil && !rollout.IsEmpty() {
		return rollout, nil
	}

	rolloutDB, errDB := this.gateway.FindById(key)
	go func() {
		if errDB == nil && !rolloutDB.IsEmpty() {
			_, err := this.cacheGateway.Save(rollout)
			if err != nil {
				log.Println("ff-4-go error: error to save in cache gateway. err:", err.Error())
			}
		}
	}()
	return rolloutDB, errDB
}

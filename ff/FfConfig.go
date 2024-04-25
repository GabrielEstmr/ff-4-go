package ff

type FfConfig struct {
	configData      *FfConfigData
	featuresMethods FeaturesMethods
	registerMethods RegisterMethods
}

func NewFfConfig(
	configData *FfConfigData,
	featuresMethods FeaturesMethods,
	registerMethods RegisterMethods) *FfConfig {
	return &FfConfig{
		configData:      configData,
		featuresMethods: featuresMethods,
		registerMethods: registerMethods}
}

func (this FfConfig) GetConfigData() *FfConfigData {
	return this.configData
}

func (this FfConfig) GetFeaturesMethods() FeaturesMethods {
	return this.featuresMethods
}

func (this FfConfig) GetRegisterMethods() RegisterMethods {
	return this.registerMethods
}

package ff_configs_resources

type CacheType string

const (
	REDIS CacheType = "REDIS"
)

var CacheTypeEnum = map[CacheType]CacheType{
	REDIS: REDIS,
}

var CacheTypeEnumFromNames = map[string]CacheType{
	"REDIS": REDIS,
}

func (this CacheType) Exists() bool {
	_, exists := CacheTypeEnum[this]
	return exists
}

func (this CacheType) FromValue(value string) CacheType {
	valueMap, exists := CacheTypeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this CacheType) Name() string {
	valueMap, exists := CacheTypeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}

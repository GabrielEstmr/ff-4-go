package ff_configs_resources

type DbType string

const (
	MONGO DbType = "MONGO"
)

var dbTypeEnum = map[DbType]DbType{
	MONGO: MONGO,
}

var dbTypeEnumFromNames = map[string]DbType{
	"MONGO": MONGO,
}

func (this DbType) Exists() bool {
	_, exists := dbTypeEnum[this]
	return exists
}

func (this DbType) FromValue(value string) DbType {
	valueMap, exists := dbTypeEnumFromNames[value]
	if exists {
		return valueMap
	}
	return ""
}

func (this DbType) Name() string {
	valueMap, exists := dbTypeEnum[this]
	if exists {
		return string(valueMap)
	}
	return ""
}

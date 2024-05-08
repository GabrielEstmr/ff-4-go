# Features for Golang (ff-4-go)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/redis/go-redis/v9)](https://pkg.go.dev/github.com/redis/go-redis/v9?tab=doc)

> ff-4-go is brought to you by :star: [**Gabriel Rodrigues**](https://www.linkedin.com/in/gabrielmoraisrodrigues/).
> Gabriel Rodrigues is a back-end software developer graduated in aerospace engineering and with more than 8 years
> of experience in the area.

## Main Features

- Provide feature management in an easy way in golang applications
- Plug and use: pass only few arguments to be allowed to both use features inside the code of the host application and
  to provide endpoints to manage them as well (for front-end access to these features for instance)
- Work with boolean features, multi-values features and rollout feature control

## Compatibility

- Go 1.21.4 or above
- github.com/redis/go-redis/v9 v9.5.1 or above
- go.mongodb.org/mongo-driver v1.15.0 or above

## Installation

Make sure to initialize a Go module:

```shell
go mod init github.com/my/repo
```

Then install go-redis/v9:

```shell
go get github.com/redis/go-redis/v9
```

## Quickstart

### Client arguments instantiation:

#### Basic (no caching:)

```go
     ffClientArgs := *ff_configs_resources.NewMongoFfConfigData(
      main_configs_mongo.GetMongoDBDatabaseBean());
```

#### Full customized client arguments instantiation:

```go
    ffClientArgs := ff_configs_resources.NewMongoFfConfigData(
      main_configs_mongo.GetMongoDBDatabaseBean()).
      WithCustomFeatureFlagColName("my-ff-feature-flags-collection"). // in case of custom feature-flags collection name
      WithCustomFeatureColName("my-ff-features-collection"). // in case of custom features collection name
      WithCustomRolloutColName("my-ff-rollouts-collection"). // in case of custom rollouts collection name
      WithRedisCache(main_configs_cache.GetRedisClusterBean()). // in case of add caching to the resources
      WithCustomCachePrefix("custom-ff-caching-prefix"). // in case of multiple application sharing the same caching client (to avoid conflicts)
      WithFeatureFlagsInitialValues(featureFlagsLib).    // in case of initial state for feature-flags
      WithFeaturePropertiesInitialValues(featuresLib). // in case of initial state for feature-properties
      WithRolloutsInitialValues(rolloutsLib).          // in case of initial state for rollouts
      WithCustomBaseRoutePath("/ff") // in case of custom base uri for the endpoints routes
```

### Client instantiation:

Use the client arguments instantiated as above and the client factory to obtain the ff-4-go client

```go
    client, err := ff_configs_factories.NewClientFactory(ffClientArgs).Build()
    if err != nil {
      log.Panicf("%s: %s", "Error to instantiate ff-4-go client", err.Error())
    }
```

Through the client you have access to methods to manipulate feature-flags, feature-properties and rollouts in the code
of your host application:

```go
  isFFEnabled, err := client.GetFeaturesMethods().IsEnabled("my-feature-flag-key")

featureProperty, err := client.GetFeaturesPropertyMethods().FindById("my-feature-property-key")

isTargetInRollout, err := client.GetRolloutMethods().IsEnabledAllOrTargetInRollout("my-rollout-key")
```

Through the client you have access the functions to use in your ws handler

```go
  router:= *mux.NewRouter()
routeFns := client.GetRouteFn()
for _, v := range routeFns {
router.HandleFunc(v.URI, v.ControllerFunc).Methods(v.Method)
}
```

All available Methods

```go
type FeaturesMethods interface {
  Create(featureFlag ff_domains.FeatureFlag) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
  Delete(key string) ff_domains_exceptions.LibException
  Enable(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
  Disable(key string) (ff_domains.FeatureFlag, ff_domains_exceptions.LibException)
  IsEnabled(key string) (bool, ff_domains_exceptions.LibException)
  IsDisabled(key string) (bool, ff_domains_exceptions.LibException)
}

type FeaturePropertyMethods interface {
  Create(property ff_domains.FeatureProperty) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
  Update(property ff_domains.FeatureProperty) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
  Delete(key string) ff_domains_exceptions.LibException
  FindById(key string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
  AddValueToProperty(key string, value string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
  RemoveValueToProperty(key string, value string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
  Enable(key string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
  Disable(key string) (ff_domains.FeatureProperty, ff_domains_exceptions.LibException)
}

type RolloutMethods interface {
  Create(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException)
  Update(rollout ff_domains.Rollout) (ff_domains.Rollout, ff_domains_exceptions.LibException)
  Delete(key string) ff_domains_exceptions.LibException
  FindById(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
  AddTargetToRollout(key string, target string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
  RemoveTargetFromRollout(key string, target string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
  EnableToAll(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
  DisableToAll(key string) (ff_domains.Rollout, ff_domains_exceptions.LibException)
  IsEnabledAllOrTargetInRollout(key string, target string) (bool, ff_domains_exceptions.LibException)
}
```

## How to Use

### Via Endpoints

### Via Methods in your host application

## Available Endpoints

1. Feature Properties Endpoints

- Create Feature Property

```shell
curl --request POST \
  --url http://localhost:8081/ff/v1/feature-properties \
  --header 'Content-Type: application/json' \
  --data '{
	"key": "FEATURE_TEST",
	"group": "Group",
	"description": "description",
	"enabled": true,
	"values": {
    "value1": "value1",
    "value2": "value2",
    "value3": "value3"
  }
}'
```

- Update Feature Property

```shell
curl --request PUT \
  --url http://localhost:8081/ff/v1/feature-properties/FEATURE_TEST \
  --header 'Content-Type: application/json' \
  --data '{
	"group": "Group",
	"description": "description",
	"enabled": true,
	"values": {
    "value1": "value1",
    "value2": "value2",
    "value3": "value3",
		"value4": "value4"
  }
}'
```

- Delete Feature Property

```shell
curl --request DELETE \
  --url http://localhost:8081/ff/v1/feature-properties/FEATURE_TEST
```

- Get Feature Property by its key

```shell
curl --request GET \
  --url http://localhost:8081/ff/v1/feature-properties/FEATURE_TEST
```

- Add value to a Feature Property

```shell
curl --request PUT \
  --url http://localhost:8081/ff/v1/feature-properties/FEATURE_TEST/values/new_value/remove
```

- Remove value to a Feature Property

```shell

```

- Enable Feature Property

```shell

```

- Disable Feature Property

```shell

```

## Application Example

## Contact

Any question or improvement contact to:

```shell
gabriel.estmr@gmail.com
```


# Features for Golang (ff-4-go)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/redis/go-redis/v9)](https://pkg.go.dev/github.com/redis/go-redis/v9?tab=doc)

> ff-4-go is brought to you by :star: [**Gabriel Rodrigues**](https://www.linkedin.com/in/gabrielmoraisrodrigues/).
> Gabriel Rodrigues is a back-end software developer graduated in aerospace engineering and with more than 8 years 
> of experience in the area.


## Main Features
- Provide feature management in an easy way in golang applications
- Plug and use: pass only few arguments to be allowed to both use features inside the code of the host application and to provide endpoints to manage them as well (for front-end access to these features for instance)
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


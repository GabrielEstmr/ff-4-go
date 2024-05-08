# Features for Golang (ff-4-go)

## Main Goals

## Main Features
- 


## Compatibility

## How to Install
```shell
go get aushuahsuhas
```

## Client Examples

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


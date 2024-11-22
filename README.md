# gems-bidding-engine

## TODO 
* Determine how we want to save to transactional database
* 


## Added libraries 
Try to keep adding any new packages that we need to deliver the speed that we need for this database. 

* [GoRedis](https://pkg.go.dev/github.com/go-redis/redis/v8#section-readme)
  * State management 
* [AWS for Go](github.com/aws/aws-sdk-go)
  * Connects to AWS to talk with the services needed
* [Viper](https://pkg.go.dev/github.com/dvln/viper)
  * Environment setup - will read config files in many formats which is helpful when 
* [Zap](https://pkg.go.dev/go.uber.org/zap#section-readme)
  * Takes care of all of the logging implementations that will be needed for the routing 
  * https://betterstack.com/community/guides/logging/go/zap/

## Redis 
This will keep track of bid state and assurance that the proper functions continue to occur 


## Setting up Docker 
Docker file has already been created, will need to iterate further to ensure we have the 
right logging in place for this work. 

```
docker build -t bidding-service .
docker run -p 8080:8080 bidding-service
```


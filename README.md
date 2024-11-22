# gems-bidding-engine

## TODO 
* Determine how we want to save to transactional database
* 


## Added libraries 
Try to keep adding any new packages that we need to deliver the speed that we need for this database. 

* [GoRedis](https://pkg.go.dev/github.com/go-redis/redis/v8#section-readme)
* [AWS for Go](github.com/aws/aws-sdk-go)
* [Viper](https://pkg.go.dev/github.com/dvln/viper)
  * Environment setup - will read config files in many formats which is helpful when 

## Redis 
This will keep track of bid state and assurance that the proper functions continue to occur 


## Setting up Docker 
Docker file has already been created, will need to iterate further to ensure we have the 
right logging in place for this work. 

```
docker build -t bidding-service .
docker run -p 8080:8080 bidding-service
```


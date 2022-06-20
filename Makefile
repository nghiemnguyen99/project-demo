#gen :
#	protoc --go_out=plugins=grpc:. redis/sumpb/sumpb.proto
#run-redis :
#	go run .\cmd\redis_service\main.go

#run-main :
#	go run .\cmd\main_service\main.go

#build : run-redis run-main

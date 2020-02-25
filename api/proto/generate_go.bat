ECHO test
protoc -I v1/ --go_out=plugins=grpc:v1 v1/world.proto 
PAUSE
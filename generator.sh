protoc -I greet/ --go_out=. --go-grpc_out=. greetpb/greet.proto
protoc -I calculator/ --go_out=calculator/ --go-grpc_out=calculator/ calculatorpb/calculate.proto
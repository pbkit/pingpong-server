PROTO_ROOT=.pollapo/

intro:
	@echo ""
	@echo ">>>>>> $(RECIPE) >>>>>>"
	@echo ""

outro:
	@echo ""
	@echo "<<<<<< $(RECIPE) <<<<<<<"
	@echo ""

protoGen:
	@make intro RECIPE=protoGen
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative -I=$(PROTO_ROOT) $(PROTO_ROOT)pbkit/interface-pingpong-server/pingpong.proto
	@make outro RECIPE=protoGen

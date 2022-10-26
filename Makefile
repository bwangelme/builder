.PHONY: build serve proto

proto:
	protoc -I=. --go_out=. builder/pbs/recordpb/record.proto builder/pbs/recordpb/recordview.proto

build:
	go build -o ./tmpdir/builder .

serve: build
	./tmpdir/builder

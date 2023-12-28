gen_flatbuff_go:
	flatc -g -o ./flatbuffer_gen schema/users.fbs

gen_flatbuff_ts:
	flatc --ts -o ./web/flatbuffer_gen schema/users.fbs

run_web:
	(cd web && pnpm dev)

build_web:
	(cd web && pnpm build)

run_prod_web: build_web
	(cd web && pnpm start)

run_go:
	go run cmd/http/main.go

build_go:
	go build -o bin/http cmd/http/main.go

run_prod_go: build_go
	./bin/http

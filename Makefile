run:
	go run .

build-libs:
	@cd nmap && make libnmap.a
	@mkdir -p lib
	@for i in `find ./nmap/ -type f -name "*.a"`; do cp $$i ./lib/ ; done

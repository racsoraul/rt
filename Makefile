test:
	go test -race ./...

buildProjectile:
	go build -o ./bin/projectile cmd/projectile/main.go

buildBounce:
	go build -o ./bin/bounce cmd/bounce/main.go

buildAll: buildBounce buildProjectile

clean:
	rm -r ./bin
	rm *.ppm

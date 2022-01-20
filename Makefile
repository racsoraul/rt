test:
	go test -race ./...

buildProjectile:
	go build -o ./projectile cmd/projectile/main.go && ./projectile

clean:
	rm ./projectile

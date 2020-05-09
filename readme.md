
# Run the application
- Download migration tool goose: https://github.com/pressly/goose/releases
## Steps to run aplication:

- move to migration
```
$ cd migration
```
- run migration with database's config: "host=localhost port=5432 dbname=testgoose password=tamon12 user=postgres sslmode=disable"
```
$ goose postgres "host=localhost port=5432 dbname=testgoose password=tamon12 user=postgres sslmode=disable" up
```
- return project
```
$ cd ../
```
- run application
```
$ go run main.go
```


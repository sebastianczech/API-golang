# API for useful resources e.g. lubimyczytac.pl, imdb.com

[![BCH compliance](https://bettercodehub.com/edge/badge/sebastianczech/API-golang?branch=master)](https://bettercodehub.com/)


## Prepare API to run local and on Heroku

```
mkdir /Users/seba/go/src/github.com/sebastianczech
ln -s /Users/seba/Projects/go/api/ /Users/seba/go/src/github.com/sebastianczech/API-golang

cd /Users/seba/go/src/github.com/sebastianczech/API-golang
go get ./...

go mod init github.com/sebastianczech/API-golang
go mod tidy
```

## Start API on local machine

```
PORT=8080 go run main.go 
```

## Test curls for checking API

```
curl http://localhost:8080/books?find=cz%C5%82owiek%20nietoperz | json_pp

[
{
"Author": "Jo Nesbø",
"Title": "Człowiek nietoperz",
"Image": "https://s.lubimyczytac.pl/upload/books/240000/240128/335676-352x500.jpg",
"Website": "https://lubimyczytac.pl/ksiazka/240128/czlowiek-nietoperz"
}
]

curl http://localhost:8080/films?find=millenium:%20m%C4%99%C5%BCczyzni | json_pp

[
{
"Title": "Män som hatar kvinnor",
"Tv": "",
"Image": "https://m.media-amazon.com/images/M/MV5BMTc2Mjc0MDg3MV5BMl5BanBnXkFtZTcwMjUzMDkxMw@@._V1_UX182_CR0,0,182,268_AL_.jpg",
"Website": "https://www.imdb.com/title/tt1132620/?ref_=fn_al_tt_1"
}
]
```

# Useful go commands

## Download package

```bash
go get github.com/julienschmidt/httprouter
```

## Install package

```bash
go install golang.org/x/tools/cmd/godoc
```

## Format source code

```bash
go fmt
```

## Verify source code

```bash
go vet
```

## Build program

```bash
go build
go build -o api
```

## Clean binary

```bash
go clean
```

## Test code

```bash
go test
go test -cover
```

## Generate documentation and share it via HTTP server

```bash
go install golang.org/x/tools/cmd/godoc
godoc -http=:6060
```

## Use goimports for updating import lines

Command ``goimports`` updates your Go import lines, adding missing ones and removing unreferenced ones.

```bash
go get golang.org/x/tools/cmd/goimports
goimports -w .
```
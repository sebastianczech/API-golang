# API for useful resources e.g. lubimyczytac.pl, imdb.com

## Test curls for checking API

```
curl http://localhost:8080/books?find=cz%C5%82owiek%20nietoperz | json_pp
curl http://localhost:8080/films?find=millenium:%20m%C4%99%C5%BCczyzni | json_pp
```

## Prepare API to deploy on Heroku

```
ln -s /Users/seba/Projects/go/api/src/ com.sebastianczech.api
cd /Users/seba/go/src/com.sebastianczech.api

go get -u github.com/tools/godep
godep save ./...

go get -u github.com/kardianos/govendor
govendor init

cd /Users/seba/Projects/go/api
mv src/Godeps .
mv src/vendor .
```

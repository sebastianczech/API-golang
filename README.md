# API for useful resources e.g. lubimyczytac.pl, imdb.com

## Prepare API to deploy on Heroku

```
cd /Users/seba/go/src/com.sebastianczech.api

go get -u github.com/tools/godep
godep save ./...

go get -u github.com/kardianos/govendor
govendor init
```

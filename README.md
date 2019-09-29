# piqlit
### Mono Repo for **piqlit** Home Movie Manager

## Architecture
### Frontend
### Backend
#### Testing
Testing the piqlit backend is performed through Docker Compose.  To do this, start up the normal Docker Compose script by executing this command in a shell in the root project directory: 
``` bash
docker-compose up
```

Open up another tab and execute the command 

``` bash
docker-compose -f docker-compose.yml -f docker-compose.test.yml run --rm backend-master-api-test
```

This latter command should run the entire test suite for the piqlit backend.  Currently there is a 30 second timeout because this project is small and it is assumed that the test suite will never extend beyond that duration under normal circumstances.  If this changes, modify the command in the test-compose file to remove the 30 second timeout.
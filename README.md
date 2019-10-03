# piqlit
### Mono Repo for **piqlit** Home Movie Manager

## Architecture
### Frontend
### Backend

## Running Locally
Running piqlit locally is as easy as running `docker-compose build && docker-compose up` in your terminal.  Before you can do this, however, you must copy the example.env file with the name `.env` and modify any values that you wish to change.  

## Testing
The compose file defaults to using a specific volume called `db_data`.  When testing, it is a good idea to swap this out for a test database volume by setting the `PL_DB_VOL` environment variable to `db_test_data` when running `docker-compose up`. 
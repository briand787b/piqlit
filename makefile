run:
	docker-compose down
	docker-compose build
	docker-compose config
	docker-compose up

test:
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		config
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		down \
			--remove-orphans
	-docker volume rm piqlit_db_test_data
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		build
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		run --rm backend-test
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		run --rm postman-test
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		down

	

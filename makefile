run_dev:
	docker-compose down
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		down
	docker-compose build
	docker-compose config
	docker-compose up

run_prod:
	docker-compose down
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		down
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		build
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		config
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.prod.yml \
		up

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

	

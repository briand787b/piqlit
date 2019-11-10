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
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		build
	-docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		run --rm backend-test
	-docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		run --rm postman-test
	docker-compose \
		-f docker-compose.yml \
		-f docker-compose.test.yml \
		down

	
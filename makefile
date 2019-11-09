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
		up \
			--build \
			--force-recreate


	
production:
	docker compose -f common.yml -f services.yml up -d --build
development:
	docker compose -f common.yml -f services.development.yml up -d --build
local:
	docker compose -f common.yml -f services.local.yml up -d --build

clean-production:
	docker compose -f common.yml -f services.yml kill
	docker compose -f common.yml -f services.yml down
clean-development:
	docker compose -f common.yml -f services.development.yml kill
	docker compose -f common.yml -f services.development.yml down
clean-local:
	docker compose -f common.yml -f services.local.yml kill
	docker compose -f common.yml -f services.local.yml down -v
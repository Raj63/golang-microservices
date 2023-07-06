.PHONY: precommit-invoices precommit-portal reload

reload:
	docker-compose down
	docker image rm golang-microservices_invoices
	docker image rm golang-microservices_web-portal
	docker-compose up --build --remove-orphans

precommit-invoices:
	cd services/invoices && make precommit
	pwd

precommit-portal:
	pwd
	cd services/web-portal && make precommit
	pwd




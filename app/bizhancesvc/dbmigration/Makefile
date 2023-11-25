create:
	migrate create -ext sql -dir sql $(name)

up:
	migrate -path sql -database ${BIZHANCESVC_DBURI} up $(n)

down:
	migrate -path sql -database ${BIZHANCESVC_DBURI} down $(n)

goto:
	migrate -path sql -database ${BIZHANCESVC_DBURI} goto $(version)

force:
	migrate -path sql -database ${BIZHANCESVC_DBURI} force $(version)

#!/bin/bash

function full_test() {

	echo "####################################################################"

	curl -sX GET http://localhost:1324/milkyway/init | json_pp || return 1
	echo "#-----------------------------"

	curl -sX GET http://localhost:1324/milkyway/cpu | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://localhost:1324/milkyway/memR | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://localhost:1324/milkyway/memW | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://localhost:1324/milkyway/fioR | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://localhost:1324/milkyway/fioW | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://localhost:1324/milkyway/dbR | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://localhost:1324/milkyway/dbW | json_pp || return 1
	echo "#-----------------------------"

	curl -sX GET http://localhost:1324/milkyway/clean | json_pp || return 1
	echo "#-----------------------------"

}

full_test

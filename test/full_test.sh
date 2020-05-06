#!/bin/bash

HOST=${1-"localhost"}

function full_test() {

	if ( $2 == "install" )
	then
		echo "[Call Install]"
		curl -sX GET http://${HOST}:1324/milkyway/install | json_pp || return 1
		echo "#-----------------------------"
	fi

	echo "####################################################################"

	curl -sX GET http://${HOST}:1324/milkyway/init | json_pp || return 1
	echo "#-----------------------------"

	curl -sX GET http://${HOST}:1324/milkyway/cpu | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://${HOST}:1324/milkyway/memR | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://${HOST}:1324/milkyway/memW | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://${HOST}:1324/milkyway/fioR | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://${HOST}:1324/milkyway/fioW | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://${HOST}:1324/milkyway/dbR | json_pp || return 1
	echo "#-----------------------------"
	curl -sX GET http://${HOST}:1324/milkyway/dbW | json_pp || return 1
	echo "#-----------------------------"

	curl -sX GET http://${HOST}:1324/milkyway/clean | json_pp || return 1
	echo "#-----------------------------"

}

full_test

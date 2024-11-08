#!/bin/bash

while true; do
 	sleep $1
	psql -U griff -d postgres -c "SELECT pg_terminate_backend(pid) FROM pg_stat_activity WHERE datname = 'employeeManager';"
	psql -U griff -d postgres -c 'DROP DATABASE "employeeManager";'
	psql -U griff -d postgres -c 'CREATE DATABASE "employeeManager";'
	psql -U griff -d employeeManager -f /docker-entrypoint-initdb.d/dumpfile.sql
done

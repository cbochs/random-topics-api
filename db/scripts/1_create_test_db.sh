#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-SQL
    DROP DATABASE IF EXISTS apidb_test;
    CREATE DATABASE apidb_test;
SQL

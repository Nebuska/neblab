set -e

# Read your template and replace variables
sql=$(sed \
    -e "s|\${ACCOUNT_DB_USER}|${ACCOUNT_DB_USER}|g" \
    -e "s|\${ACCOUNT_DB_PASS}|${ACCOUNT_DB_PASS}|g" \
    -e "s|\${ACCOUNT_DB_NAME}|${ACCOUNT_DB_NAME}|g" \
    -e "s|\${TASKER_DB_USER}|${TASKER_DB_USER}|g" \
    -e "s|\${TASKER_DB_PASS}|${TASKER_DB_PASS}|g" \
    -e "s|\${TASKER_DB_NAME}|${TASKER_DB_NAME}|g" \
    /docker-entrypoint-initdb.d/init-postgres.sql.template)

# Execute the resulting SQL against Postgres
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<EOF
$sql
EOF
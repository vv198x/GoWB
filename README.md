# GoWB
Tools for Wildberries



docker-compose для pgsql
`
version: "3.9"

services:
  postgres:
    image: postgres:14-alpine
    command:
      - "postgres"
      - "-c"
      - "max_connections=200"
      - "-c"
      - "shared_buffers=512MB"
      - "-c"
      - "effective_cache_size=1536MB"
      - "-c"
      - "checkpoint_completion_target=0.9"
      - "-c"
      - "maintenance_work_mem=256MB"
      - "-c"
      - "log_line_prefix=%m [%p]  %q %u @%d %h -"
    restart: always
    environment:
      POSTGRES_DB: "pgdb"
      POSTGRES_USER: "pgsql"
      POSTGRES_PASSWORD: "secret"
      PGDATA: "/var/lib/postgresql/data/pgdata"
    volumes:
    - ./pgdata:/var/lib/postgresql/data
    ports:
    - "5432:5432"
	`
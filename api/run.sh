docker run --name fom --mount source=fom,target=/app -e POSTGRES_DB=fom -e POSTGRES_PASSWORD=password -p 5432:5432 postgis/postgis:latest

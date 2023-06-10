# Local Environment Usage

To initialize your local development environment, we will be using Docker Compose.

# Requirements

- [Docker-Compose](https://docs.docker.com/compose/install/)
# Usage

1. Obtain the .env File

Copy the provided .env.example file to create a new .env file:

```
cp .env.example .env
```

2. Configure Environment Variables

Open the .env file and configure the environment variables according to your requirements. You can use a text editor of your choice:

```
code .env
```

Example .env file:

```
#!/bin/bash
POSTGRES_USER="user"
POSTGRES_PASSWORD="password"
POSTGRES_DB="database_name"
POSTGRES_DB_HOST=postgresdb
POSTGRES_PORT="5432"
```

3. Start Docker Compose

Launch the Docker Compose process to set up the environment:


```
docker-compose up
```

After executing the above command, Docker Compose will create the frontend and psql-data directories. The main parent folder (the folder where this command is executed) will be automatically loaded for the Go application.

The frontend folder contains all the website files (similar to the /var/www/html folder) and is used for hosting the website.

Please follow the instructions above to set up your local development environment using Docker Compose.

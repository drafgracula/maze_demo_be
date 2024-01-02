# Makefile for running Go script in Docker

# Define a variable for the project name
PROJECT_NAME := maze_demo

# Define a variable for the path to the Docker Compose file
COMPOSE_FILE := docker/docker-compose.yml

# Build the Docker container
build:
	docker-compose -f $(COMPOSE_FILE) build --no-cache

# Run the Docker container
run:
	docker-compose -f $(COMPOSE_FILE) up

# Stop and remove the Docker container
stop:
	docker-compose -f $(COMPOSE_FILE) down

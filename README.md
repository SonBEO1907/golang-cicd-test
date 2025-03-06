# Golang CI/CD with Docker and AWS

This project demonstrates setting up a CI/CD pipeline for a simple Golang application using Docker, Docker Hub, GitHub Actions, and AWS ECS (Elastic Container Service).

## Overview
The project showcases the following:
- A simple Go application built with the **Gin** framework.
- Dockerfile for containerizing the application.
- CI pipeline to build and push Docker images to Docker Hub.
- CD pipeline to deploy the Docker image to AWS ECS.

## Project Structure
```bash
.
├── Dockerfile           # Dockerfile for building the application image
├── go.mod              # Go modules dependencies
├── go.sum              # Go modules checksums
├── main.go             # Main application code
├── .github             # GitHub workflows folder
│   └── workflows
│       ├── ci.yml      # CI/CD pipeline workflow
├── .aws                # AWS task definition JSON
└── README.md           # Project documentation
```

## Prerequisites
- Go installed locally (optional for local development)
- Docker installed
- Docker Hub account
- AWS account
- GitHub repository with appropriate secrets set up:
  - `DOCKER_USERNAME`
  - `DOCKER_TOKEN`
  - `AWS_ACCESS_KEY_ID`
  - `AWS_SECRET_ACCESS_KEY`
  - `AWS_REGION`
  - `AWS_CLUSTER_NAME`
  - `AWS_SERVICE_NAME`

## Running Locally
If you'd like to run the project locally:
```bash
go run main.go
```
Visit: http://localhost:8080/display

### Expected Response
```json
{
  "msg": "Hello World!!!"
}
```

## Docker Setup
1. Build Docker Image:
```bash
docker build -t golang-cicd-test .
```

2. Run Docker Container:
```bash
docker run -p 8080:8080 golang-cicd-test
```

## CI/CD Pipeline
The project uses GitHub Actions for automation:
- On push to the `main` branch, the pipeline builds and pushes the Docker image to Docker Hub.
- After a successful build, the pipeline deploys the application to AWS ECS.

## AWS ECS Setup
1. Create an ECS Cluster.
2. Define a Task Definition with the appropriate container configurations.
3. Create a Service to run the task on the cluster.
4. Configure AWS security groups and networking for public access.

## Environment Variables
| Variable              | Description       |
|----------------------|------------------|
| `AWS_ACCESS_KEY_ID`   | AWS access key   |
| `AWS_SECRET_ACCESS_KEY` | AWS secret key   |
| `AWS_REGION`          | AWS region       |
| `DOCKER_USERNAME`      | Docker Hub username |
| `DOCKER_TOKEN`        | Docker Hub token |

## License
This project is licensed under the MIT License.


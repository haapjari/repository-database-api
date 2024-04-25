# Greetings

Hello! I am Jari Haapasaari ([mail](mailto:haapjari@gmail.com)). This project is part of my thesis work. If you're interested into reproduce the research, please see: [repository-analysis-orchestration](https://github.com/haapjari/repository-analysis-orchestration) repository.

***

## About

- This component is a RESTful API that provides CRUD operations for a Repository Database.

***

## How-To

### Run

- You require `go` and `make` to run this project. Tested with `go-1.22.0`.
- Setup `PORT` as Environment Variable, and execute `make run` or just `PORT=8080 make run`.

### Build and Run as a Docker Container

- Build the Image: `docker build -t repository-database-api:latest .`
- Run the Image (On the Host, for Simplicity): `docker run -idt -e PORT=8080 --network=host repository-database-api:latest`

### Examples

#### GET /api/v1/repos

```bash
curl -X GET "http://localhost:8080/api/v1/repos" -H "Accept: application/json"
```

#### POST /api/v1/repos

```bash 
curl -X POST "http://localhost:8080/api/v1/repos" \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-d '{
"name": "new-repo",
"full_name": "username/new-repo",
"created_at": "2021-01-01T00:00:00Z",
"stargazer_count": 0,
"language": "Go",
"open_issues": 0,
"closed_issues": 0,
"open_pull_request_count": 0,
"closed_pull_request_count": 0,
"forks": 0,
"watcher_count": 0,
"commit_count": 0,
"network_count": 0,
"latest_release": "v1.0.0",
"total_releases_count": 1,
"contributor_count": 1,
"third_party_loc": 1000,
"self_written_loc": 2000
}'
```

#### GET /api/v1/repos/{id}

```bash
curl -X GET "http://localhost:8080/api/v1/repos/{id}" -H "Accept: application/json"
```

#### PUT /api/v1/repos/{id}

```bash
curl -X PUT "http://localhost:8080/api/v1/repos/{id}" \
     -H "Content-Type: application/json" \
     -H "Accept: application/json" \
     -d '{
          "name": "updated-repo",
          "full_name": "username/updated-repo",
          "created_at": "2021-01-02T00:00:00Z",
          "stargazer_count": 1,
          "language": "Python",
          "open_issues": 1,
          "closed_issues": 1,
          "open_pull_request_count": 1,
          "closed_pull_request_count": 1,
          "forks": 1,
          "watcher_count": 1,
          "commit_count": 1,
          "network_count": 1,
          "latest_release": "v1.0.1",
          "total_releases_count": 2,
          "contributor_count": 2,
          "third_party_loc": 1200,
          "self_written_loc": 2200
        }'
```

#### DELETE /api/v1/repos/{id}

```bash
curl -X DELETE "http://localhost:8080/api/v1/repos/{id}"
```

#### GET /api/v1/repos/normalized

```bash
curl -X GET "http://localhost:8080/api/v1/repos/normalized" -H "Accept: application/json"
```

#### POST /api/v1/repos/normalized

```bash
curl -X POST "http://localhost:8080/api/v1/repos/normalized" \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-d '{
    "name": "new-normalized-repo",
    "full_name": "username/new-normalized-repo",
    "created_at": 1622541600.0,
    "stargazer_count": 0.0,
    "language": "Go",
    "open_issues": 0.0,
    "closed_issues": 0.0,
    "open_pull_request_count": 0.0,
    "closed_pull_request_count": 0.0,
    "forks": 0.0,
    "watcher_count": 0.0,
    "commit_count": 0.0,
    "network_count": 0.0,
    "latest_release": 1.0,
    "total_releases_count": 1.0,
    "contributor_count": 1.0,
    "third_party_loc": 1000.0,
    "self_written_loc": 2000.0
}'
```

#### GET /api/v1/repos/normalized/{id}

```bash
curl -X GET "http://localhost:8080/api/v1/repos/normalized/1" -H "Accept: application/json"
```

#### PUT /api/v1/repos/normalized/{id}

```bash
curl -X PUT "http://localhost:8080/api/v1/repos/normalized/1" \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-d '{
    "name": "updated-normalized-repo",
    "full_name": "username/updated-normalized-repo",
    "created_at": 1622541700.0,
    "stargazer_count": 10.0,
    "language": "Python",
    "open_issues": 5.0,
    "closed_issues": 4.0,
    "open_pull_request_count": 3.0,
    "closed_pull_request_count": 2.0,
    "forks": 6.0,
    "watcher_count": 7.0,
    "commit_count": 8.0,
    "network_count": 9.0,
    "latest_release": 2.0,
    "total_releases_count": 3.0,
    "contributor_count": 4.0,
    "third_party_loc": 1200.0,
    "self_written_loc": 2200.0
}'
``` 

#### DELETE /api/v1/repos/normalized/{id}

```bash
curl -X DELETE "http://localhost:8080/api/v1/repos/normalized/1"
```
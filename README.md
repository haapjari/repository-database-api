# Greetings

Hello! I am Jari Haapasaari ([mail](mailto:haapjari@gmail.com)). This project is part of my thesis work. If you're interested into reproduce the research, please see: [repository-analysis-orchestration](https://github.com/haapjari/repository-analysis-orchestration) repository.

***

## About

- TBD

***

## How-To

### Run

- You require `go` and `make` to run this project. Tested with `go-1.22.0`.
- Setup `PORT` as Environment Variable, and execute `make run` or just `PORT=8080 make run`.

### Build and Run as a Docker Container

- Build the Image: `docker build -t repository-database-api:latest .`
- Run the Image (On the Host, for Simplicity): `docker run -idt -e PORT=8080 --network=host repository-database-api:latest`

### Examples

```bash
```

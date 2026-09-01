# github-stat-service

## What it does  
Backend service for pulling GitHub stats (stars, forks, open issues, last update, language, etc.). Built with Go, pulls repo stats using GitHub API, stores what it finds in PostgreSQL, and a small REST API to query the data.


## What I built it with 
Go, PostgreSQL, GitHub API. 


## What it shows 
A Go HTTP service that fetches data from the GitHub API, caches it in Postgres, and serves it through a REST API with a staleness-based refresh policy.


## API Endpoints

- `GET /stats?owner=&repo=` — returns GitHub stats for the given repo (stars, forks, open issues, last update, language). Returns cached data if available and not stale; otherwise fetches fresh from GitHub and caches it.
- `GET /health` — health check (returns 200 OK).


## How to run it 
TBD — Postgres via docker-compose, .env for connection info, go run .


## What "done" means for this project 
Done means a recruiter can clone the repo, run it against a local Postgres, hit the /stats endpoint with an owner/repo, get back GitHub stats, and see that a second hit returns cached data without another API call.

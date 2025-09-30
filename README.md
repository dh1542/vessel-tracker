# Vessel-Tracker
The Vessel-Tracker is a service that uses the free AIS stream web socket (https://aisstream.io/) to fetch position data of
ships and provides visual representation on a map. 

## Showcase
![screenshot](./docs/media/img.png "Vessel-Tracker-UI screenshot")

## How to deploy

### To Run Locally
#### Prerequisites

- Docker
- Minikube
- kubectl
- go
- react


#### 1. Install and deploy PostgreSQL
`kubectl apply -f deployment/namespace.yaml`

`helm repo add bitnami https://charts.bitnami.com/bitnami`
`helm repo update`

`helm install postgres bitnami/postgresql -n postgres -f databases/postgres/postgres-values.yaml`

port-forward to localhost:5432:

`kubectl port-forward svc/postgres-postgresql 5432:5432 -n postgres`

#### 2. Deploy Go Backend
1. Fill `/.env` with db values and api key
2. Generate sqlc functions: `cd ./AIS/db` and then `sqlc generate`
3. Run main: `cd ..` and `go run main.go`
#### 3. Deploy UI
`cd ./vessel-tracker-ui` and then run `npm install; npm run dev`
#### 4. Access in browser
Locate in your browser on http://localhost:5173/


## Architecture

            +-----------------------+
            |   AIS Stream (WS)     | (Global Firehose)
            +----------+------------+
                       |
                       v
            +----------+------------+
            |   Ingestion Service   | 
            |  - Connects to WS     |
            |  - Parses messages    | GO
            |  - Stores in DB       |
            |  - Scrapes for image  |
            +----------+------------+
                       |
                       v
            +----------+------------+
            |   PostgreSQL          | 
            +----------+------------+
                       |
                       v
            +----------+------------+
            |   REST API Service    | 
            |  - Handles HTTP GET   | GO
            |  - Queries DB by FOV  |
            +----------+------------+
                       |
                       v
            +----------+------------+
            |   React Frontend      |
            |  - Polls /api/position| REACT
            |  - params: bbox       |
            +-----------------------+








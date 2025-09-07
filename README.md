            +-----------------------+
            |   AIS Stream (WS)     | (Global Firehose)
            +----------+------------+
                       |
                       v
            +----------+------------+
            |   Ingestion Service   | (Go Service 1)
            |  - Connects to WS     |
            |  - Parses messages    |
            |  - Stores in DB       |
            +----------+------------+
                       |
                       v
            +----------+------------+
            |   PostgreSQL + PostGIS| (With spatial indexing)
            +----------+------------+
                       |
                       v
            +----------+------------+
            |   REST API Service    | (Go Service 2)
            |  - Handles HTTP GET   |
            |  - Queries DB by FOV  |
            +----------+------------+
                       |
                       v
            +----------+------------+
            |   React Frontend      |
            |  - Polls /api/vessels |
            |  - params: bbox       |
            +-----------------------+


# How to deploy
# Create namespace
kubectl apply -f deployment/namespace.yaml

# Install PostgreSQL with Helm
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update

helm install postgres bitnami/postgresql \
  -n postgres \
  -f databases/postgres/postgres-values.yaml

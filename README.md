# AIS Data Processing System Roadmap (Go + PostgreSQL + React)

## Phase 1: Core System Setup

### 1. WebSocket Client Service
- **Package**: `github.com/gorilla/websocket` or `nhooyr.io/websocket`
- **Tasks**:
  - Establish connection to `aisstream.io`
  - Implement reconnection logic
  - Parse AIS messages (NMEA 0183 format)
  - Validate and normalize incoming data

### 2. PostgreSQL + PostGIS Database
- **Setup**:
  - Docker container with PostGIS extension
  - Table schema for AIS data:
    ```sql
    CREATE TABLE vessels (
        mmsi INTEGER PRIMARY KEY,
        position GEOGRAPHY(POINT, 4326),
        speed FLOAT,
        heading FLOAT,
        timestamp TIMESTAMPTZ,
        ship_name TEXT,
        ship_type INTEGER
    );
    ```
  - Spatial index on `position` column

### 3. Data Processing Layer
- **Components**:
  - Message queue (channel-based in Go)
  - Batch inserter (for efficient DB writes)
  - Data validation
- **Optimizations**:
  - Connection pooling (`pgxpool`)
  - Prepared statements

## Phase 2: API Service (Standard Library Only)

### 1. HTTP Server
- **Packages**: `net/http`, `encoding/json`
- **Endpoints**:
  - `GET /api/v1/vessels` - GeoJSON of current vessels
  - `GET /api/v1/vessels/{mmsi}` - Single vessel history
  - `GET /api/v1/area?bbox=minLon,minLat,maxLon,maxLat` - Spatial query

### 2. GeoJSON Response Builder
```go
type FeatureCollection struct {
    Type     string     `json:"type"`
    Features []Feature  `json:"features"`
}

type Feature struct {
    Type       string     `json:"type"`
    Geometry   Geometry   `json:"geometry"`
    Properties Properties `json:"properties"`
}

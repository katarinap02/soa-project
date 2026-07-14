# 🗺️ SOA Travel & Tour App — Microservices Platform (Go / Golang)

This project represents an advanced microservices platform for creating, searching, and executing tourist tours, managing user connections, and writing travel blogs. The system is built entirely in **Go (Golang)** and deployed via a containerized architecture featuring isolated databases and advanced distributed systems patterns.

---

## 🏗️ Tech Stack

*   **Language:** Go (Golang)
*   **Communication:** gRPC (RPC protocol) & REST (HTTP)
*   **Databases:**
    *   **PostgreSQL** (Relational database for users, stakeholders, and core tour metadata)
    *   **MongoDB** (NoSQL document store for blogs and tour execution sessions)
    *   **Neo4j** (NoSQL graph database for the user social graph and recommendations)
*   **Orchestration & Infrastructure:** Docker, Docker Compose, API Gateway
*   **Monitoring & Observability:** Jaeger (Tracing), Grafana Loki (Logging), Prometheus & Grafana (Metrics)

---

## ⚙️ System Architecture & Microservices

The application is built around an **isolated API Gateway** that routes all client-side traffic to the appropriate microservices. Each service is an independent entity running on its own dedicated port with an isolated database instance.

### Microservices:
1.  **Stakeholders & Auth Service:** Manages user roles (Administrator, Guide, Tourist), registration, user profiles, and account blocking. *(Database: PostgreSQL)*
2.  **Blog Service:** Handles blog creation (with Markdown support), comment management, and blog post liking mechanics. *(Database: MongoDB)*
3.  **Followers Service:** A dedicated social graph service. Implemented using **Neo4j** to manage user follows, enforce commenting rules (restricting comments to followed accounts only), and provide friend-of-a-friend social graph recommendation queries.
4.  **Tour Service:** Manages the tour lifecycle (draft, published, archived states), map key points coordinates, automatic route distance calculation, and tourist reviews. *(Database: PostgreSQL)*
5.  **Purchase Service:** Implements shopping cart mechanisms, order items cost calculations, and secures checkout tokens (`TourPurchaseToken`) which unlock full tour details upon purchase.
6.  **Tour Execution Service:** Tracks active tourist run sessions (`TourExecution`). Incorporates a **Position Simulator** that checks the tourist's current coordinates against key point locations, automatically logging completion times when a key point is reached. *(Database: MongoDB)*

---

## 🚀 Key Implemented Features & Core Concepts

### 🐳 Containerization & System Orchestration
*   **Multi-Stage Configurations:** Designed customized multi-stage `Dockerfile` configurations for each Go microservice to ensure minimal container footprints.
*   **Isolated Architecture:** Configured a unified `docker-compose.yml` to spin up all services and companion databases under an isolated virtual bridge network, ensuring secure container-to-container communication.

### 📊 Advanced NoSQL Data Modeling & API Routing
*   **Graph Recommendations:** Built Cypher queries in Neo4j to dynamically suggest new accounts to follow based on deep graph connection paths and shared relationships.
*   **Document Storage:** Utilized MongoDB to efficiently manage deeply nested document structures for blog content and session tracking states.
*   **Centralized Entry Point:** Deployed an isolated API Gateway routing system to abstract the internal microservice topology from the client, presenting a unified interface.

### 🛡️ Distributed System Patterns & Observability
*   **High-Speed RPC Protocol:** Configured internal communications (Gateway-to-service and inter-service requests) using optimized **gRPC (RPC)**.
*   **SAGA Pattern:** Implemented a SAGA pattern to maintain transactional consistency across multiple microservices (such as coordinating shopping cart checkouts and generating purchase tokens).
*   **Distributed Tracing (Jaeger):** Embedded tracing across the service mesh to visualize request lifecycles and quickly identify performance bottlenecks or network latencies.
*   **Centralized Log Aggregation (Loki):** Automated log collection across all Docker containers, enabling real-time centralized querying.
*   **Prometheus & Grafana Metrics:**
    *   *Host Metrics:* Full performance tracking (CPU, RAM, disk I/O, and network throughput) of the host environment.
    *   *Container Metrics:* Real-time resource consumption tracking for each individual active microservice container.

---

## 👥 Development Team

*   **Katarina Petrović**
*   **Nataša Radmilović**
*   **Mateja Stevanović**
*   **Tanja Rizović**

---

## 🛠️ Getting Started

All system components, databases, and monitoring dashboards are configured to boot up using a single orchestrator command.

### Prerequisites:
*   Installed [Docker](https://www.docker.com/) and Docker Compose.

### Launching the platform:
Run the following command in the root directory of the project:

```bash
docker-compose up --build

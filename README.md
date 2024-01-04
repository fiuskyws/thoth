# thoth
Simple Key-Value Database written in Go.

## POC Releases

### Base - POC Release 1
- DONE!
- [x] Table
    - [x] Create
    - [x] List
- [x] Commands
    - [x] `SET`
    - [x] `GET`
- [x] Server
- [x] Makefile:
    - [x] `run`
    - [x] `test`

### Event Sourcing - POC Release 2 
- In Development
- [ ] Event Sourcing (from Thoth to Pegasus)

## Backlog
- InMemory DB Key Features (asked ChatGPT for it)

1. **Memory Storage:**
   - *Data Storage in RAM:* All data is stored in the system's main memory (RAM), enabling fast read and write operations.

2. **Data Structures:**
   - *Support for Complex Data Structures:* Provides a variety of data structures like strings, lists, sets, hashes, etc.

3. **Caching:**
   - *Caching Mechanisms:* Used as caching systems to store frequently accessed or computed data, reducing the load on backend databases.

4. **Key-Value Storage:**
   - *Key-Value Store:* Organized as a key-value store, associating each piece of data with a unique key for quick retrieval.

5. **Atomic Operations:**
   - *Atomic Operations:* Supports atomic operations, ensuring certain operations are executed as a single, indivisible unit.

6. **Persistence Options:**
   - *Persistence Mechanisms:* While primarily in-memory, some systems offer persistence options to store data on disk for recovery after system restarts.

7. **Pub/Sub Messaging:**
   - *Publish/Subscribe Mechanism:* Enables a publish/subscribe model for real-time communication between different components.

8. **Scalability:**
   - *Horizontal Scalability:* Can be scaled horizontally by adding more nodes to the cluster, distributing data and load across multiple instances.

9. **High Throughput and Low Latency:**
   - *High Throughput:* Designed to handle a large number of transactions per second.
   - *Low Latency:* Provides very low response times for data retrieval and updates.

10. **Partitioning/Sharding:**
    - *Data Partitioning:* Supports partitioning or sharding of data to distribute it across multiple nodes for improved performance and scalability.

11. **Data Expiry/TTL:**
    - *Time-To-Live (TTL):* Allows setting a time limit for how long data should be retained in memory, useful for caching scenarios.

12. **Cluster Management:**
    - *Cluster Management:* Provides tools for managing and monitoring clusters of in-memory database nodes.

13. **Security:**
    - *Access Controls:* Offers access controls and authentication mechanisms to secure the data.

14. **Replication:**
    - *Data Replication:* Supports data replication across nodes for fault tolerance and high availability.

15. **Multi-Threaded Architecture:**
    - *Multi-Threaded Processing:* Utilizes a multi-threaded architecture to efficiently handle concurrent requests.

16. **Integration with Other Systems:**
    - *Integration Capabilities:* Often integrates well with other systems and databases to complement existing data storage solutions.

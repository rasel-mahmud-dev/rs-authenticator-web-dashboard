# How to Set Up PostgreSQL Logical Replication: Use Cases and Step-by-Step Guide.



PostgreSQL replication is a powerful mechanism that ensures data consistency and availability across multiple databases. It enables data from one database (the primary) to be synchronized with one or more other databases (the replicas). This process is crucial for applications that require high availability, scalability, or distributed data sharing.

This guide explores the types of PostgreSQL replication, their benefits, and real-world use cases, followed by a step-by-step setup process for logical replication.

---

### What is PostgreSQL Replication?

PostgreSQL replication is a process where data from one database (primary) is copied and kept in sync with one or more other databases (replicas).

There are **two main types** of replication in PostgreSQL:

1. **Physical Replication**: Entire data is replicated at the block level. It’s commonly used for high availability (e.g., streaming replication).
2. **Logical Replication**: Replicates changes at the logical level (row-based). It’s more flexible and allows selective replication of data, meaning you can choose specific tables or parts of a database to replicate.

---

### Benefits of Logical Replication

1. **Selective Replication**: Replicate only specific tables or rows, improving performance and storage.
2. **Cross-Version Replication**: Allows replication between different PostgreSQL versions.
3. **Custom Transformations**: Data can be modified during replication, making it customizable.
4. **Replication to Other Systems**: Replicate data to other databases or applications, enabling cross-system data sharing.

---


### Implementing Logical Replication in Railway Reservation Systems 


**Key Features of Railway Reservation System Databases**

### Key features

#### 1. **User Management**
- Manage passenger and staff accounts, roles, and authentication.

#### 2. **Train Schedules**
- Store train routes, schedules, stops, and availability.

#### 3. **Booking and Cancellations**
- Handle ticket bookings, cancellations, and payment processing.

#### 4. **Seat Allocation**
- Track and manage seat availability in real-time.

#### 5. **Real-Time Notifications**
- Notify users about booking statuses, train delays, or schedule changes.

#### 6. **Reporting and Insights**
- Generate reports on booking trends, user activity, and system performance.

#### 7. **Offline Operations**
- Ensure uninterrupted services during connectivity issues, syncing data when online.


Railway reservation databases are designed to handle complex operations efficiently. Here are the simplified core features:


### Architecture Overview
Logical replication is configured between:

Central Database (Publisher):

Acts as the main source of truth for global data like train schedules, pricing, and user roles.
Station Databases (Subscribers):

Each station maintains a local database to ensure uninterrupted operations, even during connectivity issues.
Use Cases for Logical Replication
Train Schedules and Pricing:

**Publisher**: The central server stores the master copy of train schedules and ticket pricing.
**Subscribers**: Each station’s database receives updates to ensure consistency in schedules and ticket prices.
Seat Availability Updates:

**Subscribers**: Stations push seat allocation and cancellation changes to the central server.
**Publisher**: The central server broadcasts these updates to other stations in real time, avoiding double bookings.
User Management and Authentication:

User accounts and roles are managed centrally and replicated to all station databases for consistent access control.
Booking and Payment Status:

Stations push booking and payment updates to the central server, ensuring accurate transaction records.

### What is Logical Replication in This Use Case?

Logical replication is a powerful feature of modern database systems that allows selective replication of data across multiple servers. In the context of a railway reservation system, logical replication ensures that critical information, such as train schedules, seat availability, and bookings, remains consistent across distributed databases while enabling real-time updates and efficient synchronization.

---

#### **How Logical Replication Works**

Logical replication operates by replicating data at the *row level* based on predefined rules. It allows for:
1. **Selective Replication:** Only specific tables or parts of the database, such as train schedules or bookings, are replicated.
2. **Asynchronous Updates:** Changes made in one database are asynchronously propagated to other databases, minimizing performance overhead.
3. **Publisher-Subscriber Model:** One database (publisher) sends changes to one or more databases (subscribers).

---

#### **Why Logical Replication is Ideal for Railway Systems**

In a railway reservation system, multiple station databases operate semi-independently but need to synchronize with a central database to ensure accuracy and consistency. Logical replication addresses the following use-case scenarios:

1. **Real-Time Schedule Updates:**
    - Train schedules updated at the central database are propagated to all station databases, ensuring passengers and staff have access to the latest information.

2. **Distributed Booking Management:**
    - Bookings made at any station are replicated to the central database and other stations as needed, preventing overbooking and maintaining seat allocation accuracy.

3. **Regional Data Sharing:**
    - Logical replication can limit replication to region-specific data, ensuring that only relevant information is shared with each station.

4. **Offline Sync:**
    - Stations operating offline can use their local databases and sync changes with the central server once connectivity is restored. Logical replication handles conflict resolution during these syncs.

5. **Custom Data Sharing:**
    - Certain data, like passenger details, can be replicated selectively to comply with data privacy regulations or optimize performance.

---

#### **Benefits of Logical Replication for the Railway System**

- **Efficiency:** Reduces unnecessary data replication by focusing only on essential tables or rows.
- **Scalability:** Allows the system to scale across numerous stations without overwhelming network resources.
- **Fault Tolerance:** Local databases at stations can operate independently, ensuring uninterrupted service during central server downtime.
- **Customizability:** Flexible rules allow administrators to control what data is replicated and where.

---



### Real-World Use Case 2: E-commerce Platforms

**Use Case**: An e-commerce platform uses logical replication to keep the data in sync between the main database (handling customer transactions) and a separate reporting database (used for analytics). This ensures that the system can generate real-time reports and insights without affecting the performance of the website or app.

**Why Use Logical Replication**:

- **Better Performance**: Offloading reporting tasks to a replica keeps the primary database fast for customer-facing transactions.
- **Real-time Insights**: Business teams receive up-to-date data for decision-making, such as adjusting inventory or running promotions.
- **Scalability**: Easily add more reporting systems or tools without affecting the main database.
- **Data Accuracy**: Replication ensures that all systems have consistent and accurate data.
- **High Availability**: If the main database goes down, the replica can still provide read-only data, minimizing downtime.
- **Customizable Data Sync**: Replicate only the data that’s needed, making reporting more efficient.

---

## Setting Up PostgreSQL Logical Replication

In this guide, we’ll walk through the steps to set up **PostgreSQL logical replication** between two servers: a **primary** (publisher) and a **replica** (subscriber). This setup allows the replica server to receive real-time data changes from the primary server.

### 1. Install PostgreSQL

First, install PostgreSQL on both the **primary** and **replica** servers. On Ubuntu, you can install it by running the following commands:

```bash
sudo apt update
sudo apt install postgresql postgresql-contrib
```

Repeat this process on both the primary and replica machines.

### 2. Configure the Primary Database (Publisher)

### Edit `postgresql.conf`

On the **primary server**, you need to enable logical replication. Open the `postgresql.conf` file and update the following settings:

```
wal_level = logical
max_replication_slots = 10
max_wal_senders = 10
```

These settings ensure that PostgreSQL can send changes to the replica server.

### Edit `pg_hba.conf`

Next, configure `pg_hba.conf` to allow the **replica** server to connect to the primary for replication. Add the following line:

```
host    replication     postgres     <replica_db_ip>/32        md5
```

Replace `<replica_ip>` with the actual IP address of the replica server.

### Restart PostgreSQL

After making the changes, restart PostgreSQL on the primary server:

```bash
sudo systemctl restart postgresql
```

### 3. Configure the Replica Database (Subscriber)

### Edit `postgresql.conf`

On the **replica server**, you need to enable hot standby mode. Open the `postgresql.conf` file and add:

```
hot_standby = on
```

This setting allows the replica server to accept read queries while receiving replication data.

### Edit `pg_hba.conf`

In the `pg_hba.conf` file on the replica, allow the **primary** server to connect for replication:

```
host    replication     postgres     <primary_ip>/32       md5
```

Replace `<primary_ip>` with the actual IP address of the primary server.

### Restart PostgreSQL

After updating the configuration, restart PostgreSQL on the replica:

```bash
sudo systemctl restart postgresql
```

### 4. Set Up Logical Replication with Docker (Optional)

If you are using Docker, you can configure both the primary and replica databases using `docker-compose`.

Here's a simple `docker-compose.yml` configuration:

```yaml
services:
  central_primary_db:
    image: postgres:16-bullseye
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: 12345
      POSTGRES_DB: primary_db
    ports:
      - "5442:5432"
    volumes:
      - ./central_primary_db_data:/var/lib/postgresql/data
    restart: always
    command: [
      "postgres",
      "-c", "wal_level=logical",
      "-c", "max_replication_slots=10",
      "-c", "max_wal_senders=10"
    ]

  kamalapur_station_db:
    image: postgres:16-bullseye
    environment:
      POSTGRES_USER: postgres
      POSTGRES_PASSWORD: 12345
      POSTGRES_DB: kamalapur_station_db
    ports:
      - "5443:5432"
    volumes:
      - ./kamalapur_station_db_data:/var/lib/postgresql/data
    restart: always

    airport_station_db:
     image: postgres:16-bullseye
     environment:
        POSTGRES_USER: postgres
        POSTGRES_PASSWORD: 12345
        POSTGRES_DB: airport_station_db
     ports:
        - "5444:5432"
     volumes:
        - ./airport_station_db_data:/var/lib/postgresql/data
     restart: always  
```

### 5. Create a Publication on the Primary Database

On the **primary** database, you need to create a **publication** to define which tables will be replicated. For example, if you want to replicate the `users` table, use this SQL command:

```sql
-- Connect to the primary database
\c central_primary_db

-- Create a publication for the `users` table or multiple table vai comma sepa.
CREATE PUBLICATION central_railway_system_pub
   FOR TABLE
      users,                      
      train_schedules,            
      tickets,                    
      payments,                   
      seat_reservations           
   WITH (publish = 'insert, update, delete');
```

### 6. Create a Subscription on the Replica Database

On the **replica** database, you need to create a **subscription** that connects to the primary database and starts receiving changes. Use this SQL command:

```sql
-- Connect to the replica database
\c replica_db

-- Create a subscription to the primary database
CREATE SUBSCRIPTION central_railway_system_sub
   CONNECTION 'host=<primary_ip> port=5432 dbname=primary_db user=replicator_user password=<secure_password> sslmode=require'
   PUBLICATION central_railway_system_pub
   WITH (
      copy_data = true,   -- Copy existing data to replica db.
      create_slot = true,  
      enabled = true       
    );

```

Replace `<primary_ip>` with the IP address of your primary database.

Once the subscription is created, the replica will start receiving real-time updates from the primary.



### Summary

PostgreSQL logical replication is an essential tool for modern data-driven applications, offering flexibility, performance, and high availability. By selectively replicating tables and rows, logical replication addresses unique challenges in use cases like toll booth management and e-commerce analytics. This guide has provided a practical walkthrough for setting up logical replication, ensuring reliable and consistent data synchronization across systems.

Acknowledgment: This content from my  personal experience and ai.
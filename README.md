# NATS Streaming Service

Simple service that stores order data in database and in-memory cache

Demonstration: https://drive.google.com/file/d/1G3eLTmiovblVQZqJ7C8FdGGQ8zfvJ02K/view?usp=sharing


### Tech stack:
- database: PostgreSQL
- streaming platform: NATS Streaming
- containerization: Docker Compose


### Services:
- Publisher: connects to nuts-streaming and publishes json data
- Subscriber: once started, loads data from database to cache, then connects to nuts-streaming and stores valid data to database and cache, runs http-server and outputs data by order uid


### Running:
``` 
make
```


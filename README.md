# kafka

Apache Kafka is an open-source distributed event streaming platform. After going over its use-cases and capabilites, i am super curious to learn about it and build application taking leverage of 
this awosome platform to build a high performing dtributed application. The capabilites of the kafka system convinced me after trying it out in my microservices, kafka can be used to reduce the high latency and Increase response time even during traffic spikes or high volume data. 

i will share my experiances and learning here using basic producer and consumer in go. Hope this may be of use of others making similar jurney as me. 


## What is event streaming?

Event streaming refers to the ongoing delivery of events. Event could be a real time event originating from the sensor(like medical device tracking a patiant), it could be an transaction like playment from payment processing platform or any other system generating events (emiting some change in the system).Processing refers to taking action on events based on predefined policy or based on machine learning model. 

## Reason of using solutions like Apache Kafka
Increased Reliability
Queues make your data persistent, If one part of the system is ever unreachable, the other can still continue to interact with the queue.
Better Performance
Message queues enable asynchronous communication, Endpoints that are producing and consuming messages interact with the queue, not each other.
Scalability
When workloads peak, Message queues helps to scale precisely where you need to.

### use cases 
https://kafka.apache.org/intro#intro_usage


### Docker Compose 
https://github.com/conduktor/kafka-stack-docker-compose




Topic- 
1. What is kafka 
2. Its construct 
3. Concepts - Topic, partition, consumer group, offset
4. Golang lib.
5. Write your consumer and producer 


### Referance
https://medium.com/@jhansireddy007/how-to-parallelise-kafka-consumers-59c8b0bbc37a


cd cp-all-in-one
git checkout 5.5.4-post
# go-micro
Simple Go microservice project with a broker, mail, logging and listener service. AMQP and RabbitMQ are also used. Docker and k8s are used for deployment.

# Deployment
## Kubernetes
Before launching k8s with all its services make sure to fire up the external `postgres.yml` with `docker-compose -f .\postgres.yml up -d`
## LoadBalancer setup
`kubectl expose deployment broker-service --type=LoadBalancer --port=8080 --target-port=8080`

## Troubleshooting

### Windows: Makefile *** missing separator.  Stop.
https://www.youtube.com/watch?v=2nM6DBE0blA

### Deploying with Docker Swarm
When deploying with Docker Swarm, you'll need to edit your `hosts` file accordingly:
```
127.0.0.1       localhost backend
::1             localhost backend
```

# test-project
![Go](https://github.com/s-akhmedoff/test-project/workflows/Go/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/s-akhmedoff/test-project)](https://goreportcard.com/report/github.com/s-akhmedoff/test-project)

Server that sends n number of random messages to a telegram channel/group using telegram bot token with a rate of 1 message per second. Also, a bunch of messages have one of these priorities: Low, Medium, and High. Highest priority message should go first no matter how busy the server is. There're an API gateway and at least a microservices and they're connected using rabbitmq pup/sub, where API gateway receive post request (n, priority) and generates n random messages and publishes to a message broker where on the other side a service consume it with a constant rate (1msg/s)and sends to a telegram channel/group. It also has akc/nack. Of cource, I used DevOps skills, so Github CI provides CI/CD, Deployed on Heroku

## How to
Makefile provides automated deploying
1. make dep
2. make sure docker is running
3. make env_up
4. make api_up
5. make generator_up
6. make bot_up
7. make example
8. Enjoy

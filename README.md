# Celestia-Performance-Analyze
Exploring the Impact of a Light Node on Server Performance: A Golang Case Study
# How To Run
`go mod tidy`
that helps to install all the requirement packages


## Introduction
In today's world, servers play a critical role in running a wide range of applications, from simple websites to complex machine learning models. Understanding server performance and resource utilization is vital for optimizing your applications and ensuring a smooth user experience.

In this README, we'll discuss the impact of running a light node on a server with the following specifications: 16 vCPUs, 32 GB RAM, and 350 GB SSD. Using Golang, we conducted two different tests: one with an empty server and another with a light node running. We'll delve into the details of the test setup, results, and implications for server performance.

## Test Setup
To accurately measure the impact of the light node on server performance, we conducted two separate tests using Golang:

### Test 1 - Empty Server: This test measured the baseline CPU usage percentage of the server when it was idle, without any applications or services running.

### Test 2 - Light Node Running: In this test, we ran a light node with hardware requirements of a single-core CPU on the server, and measured the CPU usage percentage.

Both tests were run for a duration of 6 hours, and the average CPU usage percentage was calculated for each test.

## Test Results
### Test 1 - Empty Server: The average CPU usage percentage for the empty server was found to be approximately 0.016%. This low value indicates minimal background processes and services running on the server.

### Test 2 - Light Node Running: With the light node running, the average CPU usage percentage increased to around 0.7%. This increase is expected as the light node requires additional resources to function.

## Analysis and Implications
The test results demonstrate that running a light node on the server has a noticeable impact on CPU usage. While the increase from 0.016% to 0.7% may not seem significant, it represents a 43-fold increase in CPU usage percentage.

However, considering the server's specifications, the increased CPU usage is still well within acceptable limits. With 16 vCPUs available, the server has ample processing power to handle the light node's requirements, leaving plenty of resources for additional tasks and applications.

## Conclusion
Our Golang case study has demonstrated that running a light node on a server with 16 vCPUs, 32 GB RAM, and 350 GB SSD does have an impact on CPU usage, but the server can still handle the additional load comfortably. It is important to monitor and optimize server performance, especially when running resource-intensive applications or services. Understanding the implications of running a light node on server performance can help developers and system administrators make informed decisions when allocating resources and optimizing their infrastructure.

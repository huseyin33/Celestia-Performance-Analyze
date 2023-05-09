# Celestia-Performance-Analyze
![874888_analytics_512x512](https://github.com/huseyin33/Celestia-Performance-Analyze/assets/72567591/88ca27f1-d7f6-4e7a-9549-795b6d608c30)




Exploring the Impact of a Light Node on Server Performance: A Golang Case Study
# How To Run
`go mod tidy`
that helps to install all the requirement packages

# Bandwidth Test
`apt install speedtest-cli` `and speedtest-cli`

## Introduction
In today's world, servers play a critical role in running a wide range of applications, from simple websites to complex machine learning models. Understanding server performance and resource utilization is vital for optimizing your applications and ensuring a smooth user experience.

In this README, we'll discuss the impact of running a light node on a server with the following specifications: 16 vCPUs, 32 GB RAM, and 350 GB SSD. Using Golang, we conducted two different tests: one with an empty server and another with a light node running. We'll delve into the details of the test setup, results, and implications for server performance.

## Test Setup
To accurately measure the impact of the light node on server performance, we conducted two separate tests using Golang:

### Test 1 - Empty Server: This test measured the baseline CPU usage percentage of the server when it was idle, without any applications or services running.

### Test 2 - Light Node Running: In this test, we ran a light node with hardware requirements of a single-core CPU on the server, and measured the CPU usage percentage.

Both tests were run for a duration of 6 hours, and the average CPU usage percentage was calculated for each test.


## Results
### Test 1 - Empty Server: The average CPU usage percentage for the empty server was found to be approximately 0.016%. This low value indicates minimal background processes and services running on the server.
![Test1](https://github.com/huseyin33/Celestia-Performance-Analyze/assets/72567591/4553d6a6-d1a8-4a10-8599-e24d2bbbd6fa)


### Test 2 - Light Node Running: With the light node running, the average CPU usage percentage increased to around 0.7%. This increase is expected as the light node requires additional resources to function.
![Test](https://github.com/huseyin33/Celestia-Performance-Analyze/assets/72567591/f885b924-2afc-4df4-ad89-af75723a8703)

## Internet Speedtest (Bandwidth)

Internet speed is a factor that directly affects performance. Many central VPS providers, especially those in the European region, can offer high internet speeds. Although the download and upload requirements for light nodes are quite low, I thought that choosing a server location outside of Europe would contribute to decentralization. However, it seems that this decision has had a slight impact on my node's performance.

![Hız](https://github.com/huseyin33/Celestia-Performance-Analyze/assets/72567591/5a1c981e-f12c-4c4e-be0a-698f2f906d2b)

# As Celestia is updated and improved, it will become more widespread. Of course, these minor glitches will decrease over time. Celestia will be more user-friendly and efficient.

## Analysis and Implications
The test results demonstrate that running a light node on the server has a noticeable impact on CPU usage. While the increase from 0.016% to 0.7% may not seem significant, it represents a 43-fold increase in CPU usage percentage.

However, considering the server's specifications, the increased CPU usage is still well within acceptable limits. With 16 vCPUs available, the server has ample processing power to handle the light node's requirements, leaving plenty of resources for additional tasks and applications.

## Final thoughts
Our Golang case study has demonstrated that running a light node on a server with 16 vCPUs, 32 GB RAM, and 350 GB SSD does have an impact on CPU usage, but the server can still handle the additional load comfortably. It is important to monitor and optimize server performance, especially when running resource-intensive applications or services. Understanding the implications of running a light node on server performance can help developers and system administrators make informed decisions when allocating resources and optimizing their infrastructure. 

I have examined data such as CPU, internet bandwidth, and RAM in my performance analysis. I have provided information about my system requirements. Since I have a system with high specifications compared to the minimum requirements of a light node, there has not been a significant performance decrease. There have been some hiccups in the system due to updates, but these are completely normal and external factors resulting from improvements to the system.

Thank you for reading.

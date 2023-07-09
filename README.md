
# AWS Lambda Performance Comparison

This repository contains code samples and performance comparisons for [AWS Lambda](https://aws.amazon.com/lambda/) functions implemented in different programming languages. The purpose of this repo is to evaluate the performance of AWS Lambda functions written in [Go](https://go.dev/), [C#](https://learn.microsoft.com/en-us/dotnet/csharp/), and [Node.js](https://nodejs.org). 

This repository contains code samples and configuration files for saving login logs into a PostgreSQL relational database running on [AWS RDS](https://aws.amazon.com/rds/postgresql/) using AWS Lambda. 

## Introduction

AWS Lambda is a serverless computing service that allows you to run your code without provisioning or managing servers. It supports multiple programming languages, including Go, C#, and Node.js, among others. However, the performance characteristics of Lambda functions can vary depending on the programming language used.

This repository aims to compare the performance of Lambda functions implemented in Go, C#, and Node.js. By measuring factors such as cold start latency, execution time, and memory utilization, we can gain insights into the performance trade-offs between these programming languages when running on AWS Lambda.

## Setup

Create login logs table.

```sql
create table if not exists user_login_reports
(
    id              text not null primary key,
    user_pool_id    text,
    cognito_user_id text,
    region          text,
    email           text,
    user_attributes text,
    created_on_utc  timestamp default now()
);
```

Test event can be found [here](./post-auth.json)

## Testing

To test the login logging system, follow these steps:

* Invoke the Lambda function manually using sample login log events as input. You can use the AWS Management Console or the AWS CLI to trigger the function.
* Check the PostgreSQL database to ensure that the login logs have been stored correctly. Use SQL queries to retrieve and validate the data.
* Check lambda function execuiton time.


## Result

All mesasurements are in `millisecond` ( ms )

## Cold 

| Language  | 128mb  | 256mb  | 512mb  |
|-----------|--------|--------|--------|
|  Go       |  488   |   235  |   129  |
|  Node     |  297   |   128  |    52  |
|  C#       | 6990   |  3417  |  1660  |


## Warm

| Language  |  128mb |  256mb |   512mb |
|-----------|--------|--------|---------|
|  Go       |   5    |    4   |   3     |
|  Node     | 234    |  102   |  36     |
|  C#       |  62    |   42   |   8     | 



Versions:

- .NET **6.0.400**
- Go   **1.19.0**
- Node **16.17.0**


## Conclusion

Based on the performance tests conducted in this study, the following conclusions can be drawn:

* Go exhibits the lowest cold + warm start latency and generally performs well in terms of execution time and memory utilization.
* C#'s cold start latency is much higher than Go and Node.js, indicating that C# takes a longer time to initialize and start executing the code.
* Node.js also shows relatively consistent performance across different memory configurations. While it has higher warm latencies compared to Go.

The choice of programming language for AWS Lambda functions should consider the specific requirements of the application, such as response time, resource consumption, and developer familiarity.

## Disclamer

Of course, different results can be obtained by optimizing the codes, but we tested the applications that we wrote in the fastest way as soon as possible. And of course, it may be possible to get different results in different use cases.  

Java is excluded due to memory requirement. In the first test results, java gave a bad result in these memory configurations.

## License

This project is licensed under the **MIT License**. Please see the [LICENSE](./LICENSE) file for more details.

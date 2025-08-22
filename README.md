# Beego + mySql + Alibaba OSS

## Project Overview

This is a Go project based on the Beego framework, designed to interact with 
- Alibaba Cloud Object Storage Service (OSS) - The project provides an API interface that allows users to post, get and delete a string stored in Alibaba Cloud OSS.
- mySql database - demonstrates the basic Get and Post from mySql. Beego ORM automatically update the schema to the mysql database. 

## Technology Stack

- [Go](https://golang.org/) - Programming language
- [Beego v2](https://github.com/beego/beego) - Web framework
- [Bee](https://github.com/beego/bee) - Command-line tool for Beego
- [Alibaba Cloud OSS SDK](https://github.com/aliyun/aliyun-oss-go-sdk) - Alibaba Cloud Object Storage Service SDK

## Installation and Usage

### Prerequisites

- Go 1.16 or higher
- Alibaba Cloud OSS account and access credentials

### Installation Steps

1. Clone the repository
   ```bash
   git clone https://github.com/frankzhang210/alibabaoss.git
   cd alibabaoss
   ```

2. Install dependencies
   ```bash
   go mod tidy
   ```

3. Configure OSS credentials
   Set up your Alibaba Cloud OSS access credentials in `conf/app.conf`

4. Build and run

   Before starting the application, you need to configure the alibaba access key and mysql connection string in `conf/app.conf`

   ```bash
   bee run
   ```

### Docker

1. Build the Docker image
   ```bash
   sudo docker image build -t alibabaoss-alpine:2.0 -f ./alpine.Dockerfile .
   ```

2. Run the Docker container
   ```bash
      sudo docker run --rm -p 39022:8080 --name alibabaoss-container alibabaoss-alpine:2.0
   ```   

3. Access the API
   ```bash
      http://localhost:39022/
   ```   

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![LinkedIn][linkedin-shield]][https://www.linkedin.com/in/ehsansepehri/]



<!-- PROJECT LOGO -->
<br />
<div align="center">
 
  <h3 align="center">Go-Serverless-aws APIGateway</h3>

  <p align="center">
    Hi buddy 👋! <br /> Here are a few descriptions of how you can deploy and test the project
    <br />
  </p>
</div>

<!-- ABOUT THE PROJECT -->
## About The Project


It's a simple Restful API on AWS using the following tech stack:
* Serverless Framework (https://serverless.com/)
* Go language (https://golang.org/)
* AWS API Gateway
* AWS Lambda
* AWS DynamoDB
<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- GETTING STARTED -->
## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.


### Installation

_Below is an example of how you can instruct your audience on installing and setting up your app. This template doesn't rely on any external dependencies or services._

1. Clone the repo
   ```sh
   git clone https://github.com/EhsanSepehriNasab/rayka_test
   ```
2. Install NPM packages
   ```sh
   go get
   ```
 3. Build main.go
    ```sh
    make build
     ```
### Unit Test
_Below is an example of how you can run unit test._

1. Run unit tests.
   ```sh
   make test
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


### Deployment
_Below is an example of how you can deploy your changes into the aws cloud services._

1. Deploy using serverless yml file configure.
   ```sh
   serverless deploy --aws-profile ${your_profile}
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


## Usage & Test

They are 4 crud opreation method for this project `(GET POST PUT DELETE)` 

* Here is a example of how to test GET method: 
```sh
curl -X GET https://p1uf*****.amazonaws.com/dev/api/devices/0854fd2c-c56d-47de-b01d-2d29c013e4a5
```
and POST: 

```sh
curl --header "Content-Type: application/json" --request POST --data '{
"deviceModel": "/devicemodels/id4",
"name": "Sensor",
"note": "Testing a sensor.",
"serial": "A020000102"
}' https://p1uf*****.amazonaws.com/dev/api/devices
```

and DELETE: 
```sh
curl -X DELETE https://p1uf*****.amazonaws.com/dev/api/devices/0854fd2c-c56d-47de-b01d-2d29c013e4a5
```
and UPDATE: 
```sh
curl --header "Content-Type: application/json" --request PUT --data '{
"id": "0854fd2c-c56d-47de-b01d-2d29c013e4a5",
"deviceModel": "/devicemodels/id4",
"name": "Sensor",
"note": "Testing a sensor.",
"serial": "A020000102"
}' https://p1uf*****.amazonaws.com/dev/api/devices
```

The result could be `201` or `200` for succesed requests and `404` or `400` or `500` for failed ones.
DELETE method always response `nil`
<p align="right">(<a href="#readme-top">back to top</a>)</p>


## Acknowledgments

Use this space to list resources you find helpful and would like to give credit to. I've included a few of my favorites to kick things off!

* [ServerLess Framwork](https://www.serverless.com/)
* [AWS](https://aws.amazon.com/)
* [Golang](https://golang.google.cn/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>


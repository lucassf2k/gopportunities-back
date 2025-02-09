# Job Openings API

A simple API for managing job opportunities, developed in Golang using the Gin framework. The API allows creating, listing, searching, updating, and deleting job offers.

## Technologies Used

- Golang
- Gin (Web Framework)
- Makefile (Task Automation)

## Endpoints

The API follows the RESTful structure and uses version v1. The default base URL is:

**http://localhost:8080**

## Create a New Job Opening

**POST** /api/v1/opening

Request Body:

```Go
{
"role": "Senior da vida e do universo",
"company": "Ultrasoft",
"location": "Mossoró",
"remote": true,
"link": "https://vagapica.com",
"salary": 52000
}
```

## List All Job Openings

**GET** /api/v1/openings

## Get Details of a Specific Job Opening

**GET** /api/v1/opening?id={id}

## Update an Existing Job Opening

**PUT** /api/v1/opening?id={id}

Request Body:

```Go
{
"role": "CTO",
"company": "Ultrasoft",
"location": "Natal",
"remote": false,
"link": "https://vagapica.com",
"salary": 75000
}
```

## Delete a Job Opening

**DELETE** /api/v1/opening?id={id}

## How to Run

Clone the repository:

git clone https://github.com/lucassf2k/gopportunities-back.git
cd gopportunities-back

Install the dependencies:

```Bash
go mod tidy

Execute a API usando Makefile:

make run
```

License

This project is licensed under the MIT License.

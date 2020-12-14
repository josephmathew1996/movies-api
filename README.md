## Movies API

### API Documentation

## Technical Choices

- Language - Go
- MySQL Database
- Docker-compose for running the database

## Setup

### Requirements

- Ensure `Go` and `docker-compose` are installed

### Running the application

In the root directory,

- Run `docker-compose up` to start mysql and phpmyadmin services. This should create the database and the tables. This should also pull all the dependencies. Finally, API server will start on port 9000
- Upon successful start up, you should see `Started HTTP server on port 9000` message.

### Running the tests

In the root directory,

- Run `go test -v ./...`
- All the test files are located in the controllers package. Run `go test -v ./controllers/` to test the controllers package
- The tests do not depend on any external dependencies. The db interface has been mocked out in the mocks package.

## API Overview

- Performs all operations such as get all movies, add, update and delete a movie.
- Handles 2 user roles (admin/user).
- Admin has the permissions to perform all the operations while user is only able to view the movies.
- CORS is enabled
- Successfull login returns the `Bearer <access_token>` (JWT format) which is needed for accessing all secured endpoints.
- Token expiry is set as 60 minutes. Please update the value in the docker compose file if required.
- Movies searching, sorting and filtering are all implemented.

## API Methods and Status Code

- JSON is used for all requests and responses
- Includes methods - `GET`, `POST`, `PUT`, `POST`
- Status code 200 - Success
- Status code 401 - Unauthorized request
- Status code 400 - Bad Request
- Status code 404 - Requested resource not found
- Status code 500 - Server Error

### Authentication added for all secured endpoints. So provide Authorization header with the value token in the format `Bearer <tokenvalue>` for those endpoints

### 1. GET /api/v1/movies?name=&director=&page=1&count=10&sortby=id

    - Get all movies
    - Response: 200
    - Response Body:
        {
            "status": "SUCCESS",
            "statusCode": 200,
            "message": "Movies fetch success",
            "data": {
                "total": 248,
                "noOfPages": 50,
                "itemsPerPage": 5,
                "movies": [
                    {
                        "id": 1,
                        "name": "The Wizard of Oz",
                        "director": "Victor Fleming",
                        "genre": [
                            "Adventure",
                            "Fantasy",
                            "Family",
                            "Musical"
                        ],
                        "99popularity": 83,
                        "imdb_score": 8.3
                    },
                    {
                        "id": 2,
                        "name": "Star Wars",
                        "director": "George Lucas",
                        "genre": [
                            "Action",
                            "Adventure",
                            "Fantasy",
                            "Sci-Fi"
                        ],
                        "99popularity": 88,
                        "imdb_score": 8.8
                    },
                    {
                        "id": 3,
                        "name": "Cabiria",
                        "director": "Giovanni Pastrone",
                        "genre": [
                            "Adventure",
                            "Drama",
                            "War"
                        ],
                        "99popularity": 66,
                        "imdb_score": 6.6
                    },
                    {
                        "id": 4,
                        "name": "Psycho",
                        "director": "Alfred Hitchcock",
                        "genre": [
                            "Mystery",
                            "Horror",
                            "Thriller"
                        ],
                        "99popularity": 87,
                        "imdb_score": 8.7
                    },
                    {
                        "id": 5,
                        "name": "King Kong",
                        "director": "Merian C. Cooper",
                        "genre": [
                            "Adventure",
                            "Fantasy",
                            "Horror"
                        ],
                        "99popularity": 80,
                        "imdb_score": 8
                    }
                ]
            }

        }

### 2. POST /api/v1/movies

    - Creates a new movie
    - Request Body:
            {
                "name" : "Interstellar",
                "director": "Christopher Nolan",
                "genre" : ["Adventure", "Drama", "Sci-Fi"],
                "99popularity": 90.0,
                "imdb_score" : 9.0
            }
    - Response: 200
    - Response Body:
            {
                "status": "SUCCESS",
                "statusCode": 200,
                "message": "Movie created successfully",
                "data": {
                    "id": 249,
                    "name": "Interstellar",
                    "director": "Christopher Nolan",
                    "genre": [
                        "Adventure",
                        "Drama",
                        "Sci-Fi"
                    ],
                    "99popularity": 90,
                    "imdb_score": 9
                }
            }

### 3. PUT /api/v1/movies/{id}

    - Updates a movie
    - Specify the movie id as parameter
    - Request Body:
            {
                    "id": 249,
                    "name": "Interstellar Part-2",
                    "director": "Christopher Nolan",
                    "genre": [
                        "Adventure",
                        "Drama",
                        "Sci-Fi",
                        "Thriller"
                    ],
                    "99popularity": 95,
                    "imdb_score": 9.5
            }
    - Response: 200
    - Response Body:
            {
                "status": "SUCCESS",
                "statusCode": 200,
                "message": "Movie updated successfully",
                "data": null
            }

### 4. DELETE /api/v1/movies/{id}

    - Deletes a movie
    - Specify the movie id as parameter
    - Response: 200
    - Response Body:
            {
                "status": "SUCCESS",
                "statusCode": 200,
                "message": "Movie deleted successfully",
                "data": null
            }

## Progressing the artifact into production

- Separate production config file for running in production
- Jenkins can be used for testing and deploying the binary.

Steps :

1. Test and build the code using jenkins. Binary is tagged using the git tag.
2. Deploy the binary using jenkins specifying the tag of the binary. Tags allow easy rollbacks.

# Blogging Platform API Documentation

## Overview
This is a RESTful API for a blogging platform. It provides CRUD operations for blog posts and allows filtering by author and creation date.

## Endpoints

### Get All Posts
- **Endpoint:** `GET /admin/posts`
- **Description:** Retrieves a list of all posts. 
- **Parameters:**
  - `limit`: Optional. Number of items to retrieve. Default is 10.
  - `offset`: Optional. Start offset. Default is 0.
- **Responses:**
  - `200 OK`: Successful retrieval.
  - `400 Bad Request`: Invalid parameter.
  - `500 Internal Server Error`: Server-side error.


### Get All Posts by filtering
- **Endpoint:** `GET /posts/paging/`
- **Description:** Retrieves a list of all posts. Supports pagination, filtering by author, and filtering by creation date.
- **Parameters:**
  - `author`: Optional. Filter by author.
  - `created_after`: Optional. Filter by creation date (in RFC3339 format).
  - `limit`: Optional. Number of items to retrieve. Default is 10.
  - `offset`: Optional. Start offset. Default is 0.
- **Responses:**
  - `200 OK`: Successful retrieval.
  - `400 Bad Request`: Invalid parameter.
  - `500 Internal Server Error`: Server-side error.


### Get Single Post
- **Endpoint:** `GET /posts/{id}`
- **Description:** Retrieves a specific post by its ID.
- **Responses:**
  - `200 OK`: Successful retrieval.
  - `404 Not Found`: Post not found.
  - `500 Internal Server Error`: Server-side error.

### Create Post
- **Endpoint:** `POST /admin/posts`
- **Description:** Creates a new post. Requires admin permissions.
- **Request Body:** 
  - `title`: Title of the post (string).
  - `content`: Content of the post (string).
  - `author`: Author of the post (string).
- **Responses:**
  - `201 Created`: Successful creation.
  - `400 Bad Request`: Invalid input data.
  - `401 Unauthorized`: No valid authorization.
  - `500 Internal Server Error`: Server-side error.

### Update Post
- **Endpoint:** `PUT /admin/posts/{id}`
- **Description:** Updates an existing post by its ID. Requires admin permissions.
- **Request Body:** 
  - `title`: Updated title of the post (string).
  - `content`: Updated content of the post (string).
- **Responses:**
  - `200 OK`: Successful update.
  - `400 Bad Request`: Invalid parameter or input data.
  - `404 Not Found`: Post not found.
  - `500 Internal Server Error`: Server-side error.

### Delete Post
- **Endpoint:** `DELETE /admin/posts/{id}`
- **Description:** Deletes a specific post by its ID. Requires admin permissions.
- **Responses:**
  - `204 No Content`: Successful deletion.
  - `400 Bad Request`: Invalid parameter.
  - `404 Not Found`: Post not found.
  - `500 Internal Server Error`: Server-side error.


### Auth

### login Post
- **Endpoint:** `POST /auth/login`
- **Description:** Deletes a specific post by its ID. Requires admin permissions.
- **Responses:**
 - `200 OK`: Successful login.
  - `400 Bad Request`: Invalid parameter or input data.
  - `404 Not Found`: Post not found.
  - `500 Internal Server Error`: Server-side error.


### register Post
- **Endpoint:** `POST /auth/register`
- **Description:** Deletes a specific post by its ID. Requires admin permissions.
- **Responses:**
 - `201 OK`: Successful creation.
  - `400 Bad Request`: Invalid parameter or input data.
  - `404 Not Found`: Post not found.
  - `500 Internal Server Error`: Server-side error.

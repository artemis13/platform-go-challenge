# Table of contents

 1. [GlobalWebIndex Engineering Challenge](#globalwebindex-engineering-challenge)
 2. [Features](#features)
 3. [Prerequisities](#prerequisities)
 4. [Installation](#Installation)
 5. [Usage](#usage)
 6. [API Endpoints](#api-endpoints)
 7. [Running Tests](#running-tests)
 8. [Useful Resources](#useful-resources)
 9. [Contact](#contact)


# GlobalWebIndex Engineering Challenge

## Introduction

This challenge is designed to give you the opportunity to demonstrate your abilities as a software engineer and specifically your knowledge of the Go language.

On the surface the challenge is trivial to solve, however you should choose to add features or capabilities which you feel demonstrate your skills and knowledge the best. For example, you could choose to optimise for performance and concurrency, you could choose to add a robust security layer or ensure your application is highly available. Or all of these.

Of course, usually we would choose to solve any given requirement with the simplest possible solution, however that is not the spirit of this challenge.

## Challenge

Let's say that in GWI platform all of our users have access to a huge list of assets. We want our users to have a peronal list of favourites, meaning assets that favourite or “star” so that they have them in their frontpage dashboard for quick access. An asset can be one the following
* Chart (that has a small title, axes titles and data)
* Insight (a small piece of text that provides some insight into a topic, e.g. "40% of millenials spend more than 3hours on social media daily")
* Audience (which is a series of characteristics, for that exercise lets focus on gender (Male, Female), birth country, age groups, hours spent daily on social media, number of purchases last month)
e.g. Males from 24-35 that spent more than 3 hours on social media daily.

Build a web server which has some endpoint to receive a user id and return a list of all the user’s favourites. Also we want endpoints that would add an asset to favourites, remove it, or edit its description. Assets obviously can share some common attributes (like their description) but they also have completely different structure and data. It’s up to you to decide the structure and we are not looking for something overly complex here (especially for the cases of audiences). There is no need to have/deploy/create an actual database although we would like to discuss about storage options and data representations.

Note that users have no limit on how many assets they want on their favourites so your service will need to provide a reasonable response time.

A working server application with functional API is required, along with a clear readme.md. Useful and passing tests would be also be viewed favourably

It is appreciated, though not required, if a Dockerfile is included.

## Submission

Just create a fork from the current repo and send it to us!

Good luck, potential colleague!


# Features

# Prerequisities

# Installation

# Usage

# API Endpoints

The server defines the following RESTful endpoints:
``` bash
POST /users/{id}/favorites: Adds a favorite to a user favotites list.
#e.g. curl comamnd 
curl -X POST "http://localhost:1323/users/1/favorites" \
-H "Authorization: gwi-token-12345" \
-H "Content-Type: application/json" \
-d '{"Type":0,"Description":"Interested in sports","Chart":{"Title":"How interested are you in Sports?","XAxis":"Interest","YAxis":"Amount",
"Data":{"Quite interested":32,"Very interested":31.9}}}'




GET /users/{id}/favorites: Retrieve all favorities of a user with a certain id.
#e.g curl comamnd 
curl GET "http://localhost:1323/users/1/favorites" -H "Authorization:gwi-token-12345"
#curl command to retreive all user favorites using pagination
curl -X GET "http://localhost:1323/users/1/favorites?page=1&limit=1" \
-H "Authorization: Bearer gwi-token-12345" \
-H "Content-Type: application/json"


PUT /users/{id}/favorites/{id}: Edits a user\'s favorite description 
curl -X PUT "http://localhost:1323/users/1/favorites/1a80e52a-f604-41e2-9aa2-4fd0f732e649" \
     -H "Authorization: gwi-token-12345" \
     -H "Content-Type: application/json" \
     -d '{"description": "Favorite Sport Trends"}'

DELETE /users/{id}/favorites/{id}: Delete a favorite from a user\'s list.
#e.g curl comamnd
curl -X DELETE "http://localhost:1323/users/1/favorites/1a80e52a-f604-41e2-9aa2-4fd0f732e649" \
-H "Authorization:gwi-token-12345"
```
# Running Tests

To execute the automated tests for this system, run the following command in the project directory:
``` bash
go test -count=1 ./...
```
# Useful Resources

[Exploring Golang REST API Frameworks](https://dev.to/xngwng/top-5-go-rest-api-frameworks-k0e)
Tutorial: [Developing a RESTful API with Go and Gin](https://go.dev/doc/tutorial/web-service-gin)
Tutorial: [How to build an API using Go and Gorilla mux](https://dev.to/envitab/how-to-build-an-api-using-go-ffk)
Tutorial: [Build a Golang RESTful Stock API With the Echo Framework](https://betterprogramming.pub/intro-77f65f73f6d3)
[echo QuickStart](https://echo.labstack.com/docs/quick-start)
[echo CRUD](https://echo.labstack.com/docs/cookbook/crud)

# Contact

Created by [Artemis Apostolou](https://github.com/artemis13) for the purposes of GWI Go Platrform Challenge

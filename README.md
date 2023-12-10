# Pack Service

This project proposes a solution to distributing items within a package group

# Introduction

This project solves the problem of distributing items between packages, providing an API in Go that allows calculating the best distribution of items between registered package sizes, respecting the following rules:

1. Only whole packs can be sent. Packs cannot be broken open.
2. Within the constraints of Rule 1 above, send out no more items than necessary to
   fulfil the order.
3. Within the constraints of Rules 1 &amp; 2 above, send out as few packs as possible to

By default, we have the following pre-added package sizes, allowing the addition or removal of packages via API:
- 250 Items
- 500 Items
- 1000 Items
- 2000 Items
- 5000 Items

# Setup

### Prerequisites

* [**Docker**](https://www.docker.com/products/docker-desktop/) (>= 24.0.5) and
  [**Docker-compose**](https://docs.docker.com/compose/install/) (>= 2.20.3) - run containers

### Execution

To build the containers, run:
```shell
make build
```

To execute the application, run:
```shell
make run
```

If everything goes well, the API Rest server will run on port `8080`.

And to stop the application, use:
```shell
make stop
```

### Usage


To calculate the number of packages, use the following GET endpoint:
``` curl
curl --location 'localhost:8080/calculate?items=12001'
```
The API will return a JSON in the following format, showing the size of the packages used and the quantity:
``` json
[
    {
        "size": 250,
        "quantity": 1
    },
    {
        "size": 2000,
        "quantity": 1
    },
    {
        "size": 5000,
        "quantity": 2
    }
]
```


To add new packages, use the POST endpoint below:
``` curl
curl --location 'localhost:8080/pack' \
--header 'Content-Type: application/json' \
--data '{
    "size": 1
}'
```

To remove existing packages, use the DELETE endpoint below:

``` curl
curl --location --request DELETE 'localhost:8080/pack/5000'
```


To simplify validations, the application can be online accessed at:
``` curl
curl --location 'http://3.15.189.102:8080/calculate?items=12001'
```

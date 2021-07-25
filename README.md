## Ticket Sales
API deployed in a docker container that manage ticket sales from events in different countries.

The project structure is based in Domain Drive Design, organized in a few layers.

It uses a MySQL Database to store tickets information.

It has an authentication middleware that allows to use the API using the header `auth-token:1234asdf`, otherwise you will not be able to access it and will receive an unauthorized error.

## Endpoints
**POST** _/api/v1/sales_

An endpoint to register ticket sales for a global event with data about the country from where the purchase was made.

Example cURL:
```
curl --location --request POST 'localhost:8080/api/v1/sales' \
--header 'auth-token: 1234asdf' \
--header 'Content-Type: application/json' \
--data-raw '{
    "event_id": 1,
    "amount": 100,
    "sale_type": "credit_card",
    "country_id": 1,
    "country_name": "Argentina"
}'
```

**GET** _/api/v1/stats_

An endpoint to retrieve near-realtime statistics of sales totals by country.

Example cURL:
```
curl --location --request GET 'localhost:8080/api/v1/stats' \
 --header 'auth-token: 1234asdf'
```

## How to run the project
With the following command you can use the API on 8080 port.

`docker-compose up --build`
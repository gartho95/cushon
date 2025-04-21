## How To Run
rename `.env.example` to `.env` 
`go mod tidy`
`go run .`

## API Routes
`GET /funds`

Returns a list of all available funds.

Example response:

```
[
  {
    "id": "uuid",
    "name": "Fund Name"
  },
  {
    "id": "uuid",
    "name": "Fund Name"
  }
]
```

`GET /account?fundId={fundId}&userId={userId}`

Returns the details of a specific account based on the provided fundId and userId query parameters.

Request Parameters:
fundId: The ID of the fund to retrieve the account from.

userId: The ID of the user whose account details are being requested.

Example response:
```
{
  "id": "uuid",
  "balance": "£100.00",
  "fund": "Fund A",
  "user": "John Doe"
}
```

`POST /deposit`

This route allows for depositing funds into an account if the user and requested account exist.

Example payload:
```
{
  "fundBalances": [
    {
      "fundId": "uuid",
      "value": 100000
    }
  ],
  "userId": "uuid"
}
```

## Assumptions

This code was written under the assumption that a user who is inputting the information is doing so on a page/screen that renders a list of funds, and from that they are able to query an account or send a deposit to an account.

I am assuming that the request sent by the front end will be a single integer (perhaps the user inputs in £12.34 format, but the front end will push this as 1234

## Approach

I decided to use GORM as the database framework to save considerable time on building database logic and queries, as well as to be able to use it's migration and to make use of it's Dialector interface to easily allow for different database types to be used and added.

Initially I took the approach of having each database entity use it's own generic route that allows for post/get etc... and then have the actual handler logic included under a `/custom` route prefix, but went with the direct custom approach in the end to save time.

I kept the logic of seperation between retail/employer users simple to allow for extendability, and worked with the assumption that a `user` will have a `type` to tell us which to use, so essentially the database itself is the source of truth. 

As the database is the source of truth, it needs to be populated before you can provide an example body for the `deposit` route, to achieve this I run a migration and seeder (the latter for testing and demonstration purposes) to populate mock data, in a real world scenario I would likely either have the migration as part of some CI-CD implementation. 

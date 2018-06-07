# Punkt Bohater

Build a web application that allows a user to enter a date time range and get
back the rate at which they would be charged to park for that time span.

## Requirements

- API will need documentation and a contract published.
- Support JSON & XML over HTTP
- API computes price for a specified datetime range given a JSON file of rates.
- Datetime ranges specified in isoformat.
- Rate must cover a datetime range for it to be available.
- Rates will never overlap.

### Sample File

```json
{
  "rates": [
    {
      "days": "mon,tues,wed,thurs,fri",
      "times": "0600-1800",
      "price": 1500
    },
    {
      "days": "sat,sun",
      "times": "0600-2000",
      "price": 2000
    }
  ]
}
```

### Sample Result

- `2015-07-01T07:00:00Z` to `2015-07-01T12:00:00Z` should yield `1500`
- `2015-07-04T07:00:00Z` to `2015-07-04T12:00:00Z` should yield `2000`
- `2015-07-04T07:00:00Z` to `2015-07-04T20:00:00Z` should yield `unavailable`

### Test Data

```json
{
  "rates": [
    {
      "days": "mon,tues,thurs",
      "times": "0900-2100",
      "price": 1500
    },
    {
      "days": "fri,sat,sun",
      "times": "0900-2100",
      "price": 2000
    },
    {
      "days": "wed",
      "times": "0600-1800",
      "price": 1750
    },
    {
      "days": "mon,wed,sat",
      "times": "0100-0500",
      "price": 1000
    },
    {
      "days": "sun,tues",
      "times": "0100-0700",
      "price": 925
    }
  ]
}
```

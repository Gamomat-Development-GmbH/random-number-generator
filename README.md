# random-number-generator
This project is a http endpoint for serving random numbers within a given range.
The Mersenne Twister implementation of Jochen Voss (https://github.com/seehuhn/mt19937) is used.

# URL Parameters

min - lower bounce for random numbers
maxExclusive - upper bounce exclusive for random numbers
count - count of random numbers to get

# Response
```json
{
    "min": 1,
    "maxExclusive": 10,
    "count": 5,
    "randomNumbers": [4, 2, 7, 8, 1]
}
```
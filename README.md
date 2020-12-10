# random-number-generator
This project is a http endpoint for serving random numbers within a given range.
The Fortuna implementation of Jochen Voss (https://github.com/seehuhn/fortuna) is used.

# Usage
# URL Parameters
* min - lower bounce for random numbers
* maxExclusive - upper bounce exclusive for random numbers
* count - count of random numbers to get
* format - optional, possible values: 
    * json (default)
    * int32_le
    * int32_be    

# Response
In case for json format:

```json
{
    "min": 1,
    "maxExclusive": 10,
    "count": 5,
    "randomNumbers": [4, 2, 7, 8, 1]
}
```

In case for int32_le or int32_be format the response is binary.

# Cryptographic strength

The RNG Service is using an implementation of Fortuna.
To ensure the cryptographical strength cycling and reseeding was implemented in addition.

## Cycling
The cycling causes the RNG to skip a random amount of random numbers (1 up to 9).

* A background process is executing the cycling in a random time period (every 1000 up to 10000 milliseconds).
* When requesting random numbers then cycling is interrupting the process in a random step (every 100 up to 600 random numbers).

## Testing
See [README.md](test/README.md) in folder test for instructions.

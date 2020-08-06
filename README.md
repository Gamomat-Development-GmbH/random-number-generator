# random-number-generator
This project is a http endpoint for serving random numbers within a given range.
The Mersenne Twister implementation of Jochen Voss (https://github.com/seehuhn/mt19937) is used.

# Usage
# URL Parameters

min - lower bounce for random numbers
maxExclusive - upper bounce exclusive for random numbers
count - count of random numbers to get
format - optional: json, int32_le, int32_be

# Response

## format json
```json
{
    "min": 1,
    "maxExclusive": 10,
    "count": 5,
    "randomNumbers": [4, 2, 7, 8, 1]
}
```

# Run dieharder (https://webhome.phy.duke.edu/~rgb/General/dieharder.php) test suite:

## Cycling
The cycling causes the RNG to skip a random amount of random numbers (10 up to 99).

* A background process is executing the cycling in a random time period (every 1000 up to 10000 milliseconds).
* When requesting random numbers then cycling is interrupting the process in a random step (every 100 up to 600 random numbers).

## Reseeding
The reseeding is updating the random number generator seed.
A background process is executing the reseeding in a random time period (every 1000 up to 10000 milliseconds).

## Testing
See [README.md](test/README.md) in folder test for instructions.

# Run dieharder test suite:

For dieharder documentation see [here](https://webhome.phy.duke.edu/~rgb/General/dieharder.php)

## Execution

To run the dieharder test on the RNG run `docker-compose up` in this directory.
This will build a container named "rng" and a container named "dh" and execute the dieharder suite of tests.

The test execution takes about 3-4 hours on a modern device (e.g. Mac 2,6 GHz 6-Core Intel Core i7). 

Sample results from a test run 08/06/2020: 

```
[output from docker build steps omitted]
Successfully tagged test_dieharder:latest
WARNING: Image for service dieharder was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating rng ... done
Creating dh  ... done
Attaching to rng, dh
rng          | Server started... GOOS linux GOARCH amd64
dh           | #=============================================================================#
dh           | #            dieharder version 3.31.1 Copyright 2003 Robert G. Brown          #
dh           | #=============================================================================#
dh           |    rng_name    |rands/second|   Seed   |
dh           | stdin_input_raw|  7.72e+06  | 137237819|
dh           | #=============================================================================#
dh           |         test_name   |ntup| tsamples |psamples|  p-value |Assessment
dh           | #=============================================================================#
dh           |    diehard_birthdays|   0|       100|     100|0.90982049|  PASSED  
dh           |       diehard_operm5|   0|   1000000|     100|0.31404086|  PASSED  
dh           |   diehard_rank_32x32|   0|     40000|     100|0.76332385|  PASSED  
dh           |     diehard_rank_6x8|   0|    100000|     100|0.62906564|  PASSED  
dh           |    diehard_bitstream|   0|   2097152|     100|0.37943694|  PASSED  
dh           |         diehard_opso|   0|   2097152|     100|0.87216514|  PASSED  
dh           |         diehard_oqso|   0|   2097152|     100|0.42859031|  PASSED  
dh           |          diehard_dna|   0|   2097152|     100|0.85831576|  PASSED  
dh           | diehard_count_1s_str|   0|    256000|     100|0.12295043|  PASSED  
dh           | diehard_count_1s_byt|   0|    256000|     100|0.18916646|  PASSED  
dh           |  diehard_parking_lot|   0|     12000|     100|0.26323148|  PASSED  
dh           |     diehard_2dsphere|   2|      8000|     100|0.93967880|  PASSED  
dh           |     diehard_3dsphere|   3|      4000|     100|0.01229065|  PASSED  
dh           |      diehard_squeeze|   0|    100000|     100|0.32376311|  PASSED  
dh           |         diehard_sums|   0|       100|     100|0.14942409|  PASSED  
dh           |         diehard_runs|   0|    100000|     100|0.05670137|  PASSED  
dh           |         diehard_runs|   0|    100000|     100|0.03126270|  PASSED  
dh           |        diehard_craps|   0|    200000|     100|0.87906861|  PASSED  
dh           |        diehard_craps|   0|    200000|     100|0.85818553|  PASSED  
dh           |  marsaglia_tsang_gcd|   0|  10000000|     100|0.71367146|  PASSED  
dh           |  marsaglia_tsang_gcd|   0|  10000000|     100|0.59942014|  PASSED  
dh           |          sts_monobit|   1|    100000|     100|0.27738579|  PASSED  
dh           |             sts_runs|   2|    100000|     100|0.18720845|  PASSED  
dh           |           sts_serial|   1|    100000|     100|0.30845876|  PASSED  
dh           |           sts_serial|   2|    100000|     100|0.90679409|  PASSED  
dh           |           sts_serial|   3|    100000|     100|0.51853152|  PASSED  
dh           |           sts_serial|   3|    100000|     100|0.41413840|  PASSED  
dh           |           sts_serial|   4|    100000|     100|0.49138556|  PASSED  
dh           |           sts_serial|   4|    100000|     100|0.23350447|  PASSED  
dh           |           sts_serial|   5|    100000|     100|0.46445898|  PASSED  
dh           |           sts_serial|   5|    100000|     100|0.79807699|  PASSED  
dh           |           sts_serial|   6|    100000|     100|0.05266728|  PASSED  
dh           |           sts_serial|   6|    100000|     100|0.99536083|   WEAK   
dh           |           sts_serial|   7|    100000|     100|0.05702913|  PASSED  
dh           |           sts_serial|   7|    100000|     100|0.50178865|  PASSED  
dh           |           sts_serial|   8|    100000|     100|0.15561185|  PASSED  
dh           |           sts_serial|   8|    100000|     100|0.57629666|  PASSED  
dh           |           sts_serial|   9|    100000|     100|0.97859609|  PASSED  
dh           |           sts_serial|   9|    100000|     100|0.63043786|  PASSED  
dh           |           sts_serial|  10|    100000|     100|0.26475321|  PASSED  
dh           |           sts_serial|  10|    100000|     100|0.09168046|  PASSED  
dh           |           sts_serial|  11|    100000|     100|0.60676425|  PASSED  
dh           |           sts_serial|  11|    100000|     100|0.58115230|  PASSED  
dh           |           sts_serial|  12|    100000|     100|0.84758232|  PASSED  
dh           |           sts_serial|  12|    100000|     100|0.31849957|  PASSED  
dh           |           sts_serial|  13|    100000|     100|0.14537269|  PASSED  
dh           |           sts_serial|  13|    100000|     100|0.50726524|  PASSED  
dh           |           sts_serial|  14|    100000|     100|0.88320563|  PASSED  
dh           |           sts_serial|  14|    100000|     100|0.99989920|   WEAK   
dh           |           sts_serial|  15|    100000|     100|0.56923096|  PASSED  
dh           |           sts_serial|  15|    100000|     100|0.26179453|  PASSED  
dh           |           sts_serial|  16|    100000|     100|0.90892804|  PASSED  
dh           |           sts_serial|  16|    100000|     100|0.59023820|  PASSED  
dh           |          rgb_bitdist|   1|    100000|     100|0.36165991|  PASSED  
dh           |          rgb_bitdist|   2|    100000|     100|0.05009940|  PASSED  
dh           |          rgb_bitdist|   3|    100000|     100|0.19543818|  PASSED  
dh           |          rgb_bitdist|   4|    100000|     100|0.21654417|  PASSED  
dh           |          rgb_bitdist|   5|    100000|     100|0.27770860|  PASSED  
dh           |          rgb_bitdist|   6|    100000|     100|0.22722791|  PASSED  
dh           |          rgb_bitdist|   7|    100000|     100|0.07142145|  PASSED  
dh           |          rgb_bitdist|   8|    100000|     100|0.27702907|  PASSED  
dh           |          rgb_bitdist|   9|    100000|     100|0.43717432|  PASSED  
dh           |          rgb_bitdist|  10|    100000|     100|0.27465355|  PASSED  
dh           |          rgb_bitdist|  11|    100000|     100|0.14705514|  PASSED  
dh           |          rgb_bitdist|  12|    100000|     100|0.81344171|  PASSED  
dh           | rgb_minimum_distance|   2|     10000|    1000|0.25900104|  PASSED  
dh           | rgb_minimum_distance|   3|     10000|    1000|0.97475849|  PASSED  
dh           | rgb_minimum_distance|   4|     10000|    1000|0.95548546|  PASSED  
dh           | rgb_minimum_distance|   5|     10000|    1000|0.01496241|  PASSED  
dh           |     rgb_permutations|   2|    100000|     100|0.89938874|  PASSED  
dh           |     rgb_permutations|   3|    100000|     100|0.76971460|  PASSED  
dh           |     rgb_permutations|   4|    100000|     100|0.68704233|  PASSED  
dh           |     rgb_permutations|   5|    100000|     100|0.40927832|  PASSED  
dh           |       rgb_lagged_sum|   0|   1000000|     100|0.77464891|  PASSED  
dh           |       rgb_lagged_sum|   1|   1000000|     100|0.46307783|  PASSED  
dh           |       rgb_lagged_sum|   2|   1000000|     100|0.92571336|  PASSED  
dh           |       rgb_lagged_sum|   3|   1000000|     100|0.71142493|  PASSED  
dh           |       rgb_lagged_sum|   4|   1000000|     100|0.77328214|  PASSED  
dh           |       rgb_lagged_sum|   5|   1000000|     100|0.03399692|  PASSED  
dh           |       rgb_lagged_sum|   6|   1000000|     100|0.53224601|  PASSED  
dh           |       rgb_lagged_sum|   7|   1000000|     100|0.80700358|  PASSED  
dh           |       rgb_lagged_sum|   8|   1000000|     100|0.76719661|  PASSED  
dh           |       rgb_lagged_sum|   9|   1000000|     100|0.30510205|  PASSED  
dh           |       rgb_lagged_sum|  10|   1000000|     100|0.87828729|  PASSED  
dh           |       rgb_lagged_sum|  11|   1000000|     100|0.67421256|  PASSED  
dh           |       rgb_lagged_sum|  12|   1000000|     100|0.47101105|  PASSED  
dh           |       rgb_lagged_sum|  13|   1000000|     100|0.89094804|  PASSED  
dh           |       rgb_lagged_sum|  14|   1000000|     100|0.17939373|  PASSED  
dh           |       rgb_lagged_sum|  15|   1000000|     100|0.90770425|  PASSED  
dh           |       rgb_lagged_sum|  16|   1000000|     100|0.11492168|  PASSED  
dh           |       rgb_lagged_sum|  17|   1000000|     100|0.13420753|  PASSED  
dh           |       rgb_lagged_sum|  18|   1000000|     100|0.74547014|  PASSED  
dh           |       rgb_lagged_sum|  19|   1000000|     100|0.27739110|  PASSED  
dh           |       rgb_lagged_sum|  20|   1000000|     100|0.24369976|  PASSED  
dh           |       rgb_lagged_sum|  21|   1000000|     100|0.12504574|  PASSED  
dh           |       rgb_lagged_sum|  22|   1000000|     100|0.55463866|  PASSED  
dh           |       rgb_lagged_sum|  23|   1000000|     100|0.42553500|  PASSED  
dh           |       rgb_lagged_sum|  24|   1000000|     100|0.97181598|  PASSED  
dh           |       rgb_lagged_sum|  25|   1000000|     100|0.63377181|  PASSED  
dh           |       rgb_lagged_sum|  26|   1000000|     100|0.97969538|  PASSED  
dh           |       rgb_lagged_sum|  27|   1000000|     100|0.24547396|  PASSED  
dh           |       rgb_lagged_sum|  28|   1000000|     100|0.92089785|  PASSED  
dh           |       rgb_lagged_sum|  29|   1000000|     100|0.70852911|  PASSED  
dh           |       rgb_lagged_sum|  30|   1000000|     100|0.74111213|  PASSED  
dh           |       rgb_lagged_sum|  31|   1000000|     100|0.81127708|  PASSED  
dh           |       rgb_lagged_sum|  32|   1000000|     100|0.99907975|   WEAK   
dh           |      rgb_kstest_test|   0|     10000|    1000|0.32384134|  PASSED  
dh           |      dab_bytedistrib|   0|  51200000|       1|0.77634574|  PASSED  
dh           |              dab_dct| 256|     50000|       1|0.56430163|  PASSED  
dh           | Preparing to run test 207.  ntuple = 0
dh           |         dab_filltree|  32|  15000000|       1|0.11463595|  PASSED  
dh           |         dab_filltree|  32|  15000000|       1|0.77520305|  PASSED  
dh           | Preparing to run test 208.  ntuple = 0
dh           |        dab_filltree2|   0|   5000000|       1|0.82191291|  PASSED  
dh           |        dab_filltree2|   1|   5000000|       1|0.62082741|  PASSED  
dh           | Preparing to run test 209.  ntuple = 0
dh           |         dab_monobit2|  12|  65000000|       1|0.94290205|  PASSED  
```
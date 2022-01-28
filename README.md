# fpjs_coding_task

###FingerprintJS task

This problem is similar to knapsack problem. I tried to map the transaction to a simple transaction to have only the
time (weight in a standard KP)
Then use DP to solve it using the best with this generic subproblem:

    Matrix[N][T] = Best profit available without exceeding T time and using items from 0 to N.

Profits obtained for given times where:

````go
    bestProfitExpected50ms := 4139.43
    bestProfitExpected60ms := 4675.71
    bestProfitExpected90ms := 6972.29
    bestProfitExpected1000ms := 35471.812
````
(I uploaded the tests I ran to get those numbers in `transaction_test.go`)

Next step should be to implement some greedy algorithms and check how they perform.

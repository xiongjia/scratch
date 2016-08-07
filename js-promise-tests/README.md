# js-promise-tests

The simple async patterns for JS Promise (ES6 Promise).

## Series - `node test/test_series.js`
- All passed:
```
Task A [Pass] -> Task B [Pass] -> Task C [Pass] -> Final
```
- Waterfall
```
Task A [Pass] -> Task B [Err] -X-> Task C [Skip]
                    |                           
                    +--> Final [Error]          
```
- Ignore errors
```
Task A [Ignore] -> Task B [Ignore] -> Task C [Ignore] -> Final
```

## Parallel - `node test/test_parallel.js`
- All passed
```
Task A [Pass] -+                      
Task B [Pass] -|-> Final [A,B,C Pass] 
Task C [Pass] -+                      
```
- Task C error; Task A, B passed
```
Task A [Pass] -+                             
Task B [Pass] -|-> Final [A,B Pass; C Error] 
Task C [Err ] -+                             
```

## Race - `node test/test_race.js`
```
Task A [Pass (3seconds)] -+
                          |-> Final [B]
Task B [Pass (1seconds)] -+
```

## Retry - `node test/test_retry.js`
- Retry 3 times and task passed
```
A [Err; 1st] -> A [Err; 2nd] -> A [Pass; 3rd] -> Final
```

- Retry 3 times but task still failed
```
 A [Err; 1st] -> A [Err; 2nd] -> A [Err; 3rd] -> Final [Err]
```

## Until - `node test/test_until.js`
- `exports.untilPassed()` :    
  - Set the max retry times to 1000, retry delay 10ms.
  - After 100 calls, the retryTask will be switched to success. 
- `exports.untilErr()` :    
  - Set the max retry times to 1000, retry delay 10ms.
  - The retryTask always retrun error. 
  - This task will be terminated with error after 1000 times retry.


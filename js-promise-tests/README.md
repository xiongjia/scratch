# js-promise-tests

The simple async patterns for JS Promise (ES6 Promise).

## Series - `node test/test_parallel.js`
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
Task A [Ignore] -> Task B [Ignore] -X-> Task C [Ignore] -> Final
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


# gazlowe

It's my algorithm tests.

## The execution engine is Boost unit Test

- Run all tests: `gazlowe --log_level=test_suite`
- Run a specific test: `gazlowe --log_level=test_suite --run_test=<test name>`   
  <test name> is the algorithm test (e.g. --run_test=sorting/sorting_qsort )
- Get more options: `gazlowe --help`

## The list of the tests

- sorting_qsort : The QuickSort testing ( g_sorting_qsort.cxx )
- some leetcode problems : ( g_leetcode.cxx )
  - single_number: Single Number
  - max_depth_btree: Maximum Depth of Binary Tree


# gazlowe

It's my algorithm tests.

## The execution engine is Boost unit Test

- Run all tests: `gazlowe --log_level=test_suite`
- Run a specific test: `gazlowe --log_level=test_suite --run_test=<test name>`   
  <test name> is the algorithm test (e.g. --run_test=sorting/sorting_qsort )
- Get more options: `gazlowe --help`

## The list of the tests

- sorting_qsort : The QuickSort testing ( `g_sorting_qsort.cxx` )
- some leetcode problems : ( g_leetcode_{???}.cxx )
  - `g_leetcode_single_number.cxx` - Single Number
  - `g_leetcode_max_depth_btree.cxx` - Maximum Depth of Binary Tree
  - `g_leetcode_invert_btree.cxx` - Invert a binary tree
  - `g_leetcode_del_node_in_a_linked_list.cxx` - Delete Node in a Linked List
  - `g_leetcode_add_digits.cxx` - Add Digits
  - `g_leetcode_same_tree.cxx` - Same Tree
  - `g_leetcode_linked_list_cycle.cxx` - Linked List Cycle I ~ II
  - `g_leetcode_contains_duplicate.cxx` - Contains Duplicate I ~ III
  - `g_leetcode_majority_element.cxx` - Majority Element I ~ II


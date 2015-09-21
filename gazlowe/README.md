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
  - `g_leetcode_single_number.cxx` - Single Number I ~ III
  - `g_leetcode_max_depth_btree.cxx` - Maximum Depth of Binary Tree
  - `g_leetcode_invert_btree.cxx` - Invert a binary tree
  - `g_leetcode_del_node_in_a_linked_list.cxx` - Delete Node in a Linked List
  - `g_leetcode_add_digits.cxx` - Add Digits
  - `g_leetcode_same_tree.cxx` - Same Tree
  - `g_leetcode_linked_list_cycle.cxx` - Linked List Cycle I ~ II
  - `g_leetcode_contains_duplicate.cxx` - Contains Duplicate I ~ III
  - `g_leetcode_majority_element.cxx` - Majority Element I ~ II
  - `g_leetcode_queue_using_stacks.cxx` - Implement Queue using Stacks
  - `g_leetcode_stack_using_queues.cxx` - Implement Stack using Queues
  - `g_leetcode_reverse_linked_list.cxx` - Reverse Linked List
  - `g_leetcode_merge_two_sorted_lists.cxx` - Merge Two Sorted Lists
  - `g_leetcode_valid_anagram.cxx` - Valid Anagram
  - `g_leetcode_min_stack.cxx` - Min Stack
  - `g_leetcode_remove_element.cxx` - Remove Element
  - `g_leetcode_rm_dup_from_sorted_list.cxx` - Remove Duplicates from Sorted List
  - `g_leetcode_mv_zeroes.cxx` - Move Zeroes
  - `g_leetcode_implement_strstr.cxx` - Implement strStr()
  - `g_leetcode_find_min_in_rotated_sorted_array.cxx` - Find Minimum in Rotated Sorted Array


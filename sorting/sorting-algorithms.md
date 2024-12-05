# Essential Sorting Algorithms for Coding Interviews

## Quick Sort
* Time Complexity: Average O(n log n), Worst O(n²)
* Space Complexity: O(log n) average case for recursion stack
* Advantages:
    * Best practical performance in average case
    * In-place sorting (low space overhead)
    * Cache-friendly due to good locality of reference
    * Quickselect variant useful for finding kth smallest element
* Disadvantages:
    * Unstable sort (doesn't preserve relative order of equal elements)
    * Poor worst-case performance O(n²)
    * Performance depends heavily on pivot selection

## Merge Sort
* Time Complexity: O(n log n) in all cases
* Space Complexity: O(n)
* Advantages:
    * Stable sort (preserves order of equal elements)
    * Guaranteed O(n log n) performance
    * Excellent for sorting linked lists
    * Predictable performance regardless of input
    * Parallelizes well
* Disadvantages:
    * Requires O(n) extra space
    * Not in-place
    * Overkill for small arrays
    * Cache performance not as good as quicksort

## Heap Sort
* Time Complexity: O(n log n) in all cases
* Space Complexity: O(1)
* Advantages:
    * In-place sorting
    * Guaranteed O(n log n) performance
    * Excellent for finding k largest/smallest elements
    * No extra space needed unlike merge sort
* Disadvantages:
    * Unstable sort
    * Slower in practice than quicksort
    * Poor cache performance due to jumping around in memory
    * Complex implementation compared to other algorithms

## Bubble Sort
* Time Complexity: O(n²)
* Space Complexity: O(1)
* Advantages:
    * Very simple to implement
    * Stable sort
    * Adaptive (fast for nearly sorted arrays)
    * In-place algorithm
* Disadvantages:
    * Very inefficient for large datasets
    * Rarely used in practice
    * Primarily used for educational purposes

## Insertion Sort
* Time Complexity: O(n²)
* Space Complexity: O(1)
* Advantages:
    * Excellent for small arrays (< 50 elements)
    * Adaptive (very fast for nearly sorted data)
    * Stable sort
    * In-place algorithm
    * Simple implementation
* Disadvantages:
    * Quadratic time complexity makes it inefficient for large datasets
    * Not suitable for large-scale sorting

## Selection Sort
* Time Complexity: O(n²)
* Space Complexity: O(1)
* Advantages:
    * Simple implementation
    * Minimal memory usage
    * Makes minimum number of swaps (O(n))
* Disadvantages:
    * Always O(n²) even if array is sorted
    * Unstable sort
    * Not practical for large datasets

## Count sort
* Time Complexity: O(n+k) where k is the range
* Space Complexity: O(n) for unstable, O(n+k) for stable
* Advantages:
    * Good for integers or data that can map to integers in small range
    * Where range is not signifcantly larger than length
* Disadvantages:
    * Large range
    * Doesn't work for float values
    * Memory is a constraint if range is large
    * If k > nlogn => O(n+k) > O(nlogn) may as well use quicksort
        * If n = 1000, nlogn = 10000, if k>10000 quicksort probably better
    * Data with sparse distribution over range

## Interview Tips

1. Must-Know Algorithms:
   * Quick Sort and Merge Sort are the most commonly asked
   * Heap Sort is often discussed in context of priority queues
   * Understanding when to use each algorithm is crucial

2. Key Points to Remember:
   * Quick Sort is usually the best practical choice for arrays
   * Merge Sort is best for linked lists and stable sorting
   * Insertion Sort is best for small arrays or nearly sorted data
   * Built-in sort functions usually use hybrid approaches

3. Common Interview Questions:
   * Implementing Quick Sort or Merge Sort
   * Modifying sorting algorithms for specific requirements
   * Choosing the right algorithm for given constraints
   * Analyzing space/time complexity tradeoffs

# go_bubble_sort
The Bubble Sort coded in Go

Someone challenged me recently to code a bubble sort in Python without referring to any documentation. 

Luckily, I have a history with the bubble sort. I explained the bubble sort to first-year coding students while working as a T/A in college.  

That said, my college days were some time ago - maybe you've seen my photo. :) - I had to think through how the algorithm works as I wrote the code.  

I decided to take another stab at it in Go. The algorithm is the same, but the syntax is less flexible and forgiving.  

I added a verbose option to the sort function. Then verbose == true, it prints a step by step account of how the numbers are shuffled around in the array.  

One of the test cases is already sorted. The algorithm identifies it as a sorted list on the first pass through the array and avoids further processing.  

For the last test case, I generated an array of random integers between -100 and 100 and had the program sort that array with verbose turned off.  

<b>To run the program:</b>

go run bubble.go

<b>Program output:</b>

*** Bubble Sort Demo in Go ***  

*** Array before sorting: [5 3 2 4 1]  
*** Starting Sort At 2024-04-07 21:52:35.631262 -0400 EDT m=+0.000192292  
Outer loop: 0  
(i,j) = (0 0) 5 moves past 3 [3 5 2 4 1]  
(i,j) = (0 1) 5 moves past 2 [3 2 5 4 1]  
(i,j) = (0 2) 5 moves past 4 [3 2 4 5 1]  
(i,j) = (0 3) 5 moves past 1 [3 2 4 1 5]  
Outer loop: 1  
(i,j) = (1 0) 3 moves past 2 [2 3 4 1 5]  
(i,j) = (1 2) 4 moves past 1 [2 3 1 4 5]  
Outer loop: 2  
(i,j) = (2 1) 3 moves past 1 [2 1 3 4 5]  
Outer loop: 3  
(i,j) = (3 0) 2 moves past 1 [1 2 3 4 5]  
*** Finished Sort At 2024-04-07 21:52:35.631538 -0400 EDT m=+0.000468082  
*** Array after sorting: [1 2 3 4 5]  

==> Sorted array:  [1 2 3 4 5]  

*** Array before sorting: [10 20 30 40]  
*** Starting Sort At 2024-04-07 21:52:35.631594 -0400 EDT m=+0.000523795  
Outer loop: 0  
*** Finished Sort At 2024-04-07 21:52:35.631598 -0400 EDT m=+0.000528514  
*** Array after sorting: [10 20 30 40]  

==> Sorted array:  [10 20 30 40]  

*** Array before sorting: [123 99 66 11 0 -22]  
*** Starting Sort At 2024-04-07 21:52:35.631605 -0400 EDT m=+0.000535380  
Outer loop: 0  
(i,j) = (0 0) 123 moves past 99 [99 123 66 11 0 -22]  
(i,j) = (0 1) 123 moves past 66 [99 66 123 11 0 -22]  
(i,j) = (0 2) 123 moves past 11 [99 66 11 123 0 -22]  
(i,j) = (0 3) 123 moves past 0 [99 66 11 0 123 -22]  
(i,j) = (0 4) 123 moves past -22 [99 66 11 0 -22 123]  
Outer loop: 1  
(i,j) = (1 0) 99 moves past 66 [66 99 11 0 -22 123]  
(i,j) = (1 1) 99 moves past 11 [66 11 99 0 -22 123]  
(i,j) = (1 2) 99 moves past 0 [66 11 0 99 -22 123]  
(i,j) = (1 3) 99 moves past -22 [66 11 0 -22 99 123]  
Outer loop: 2  
(i,j) = (2 0) 66 moves past 11 [11 66 0 -22 99 123]  
(i,j) = (2 1) 66 moves past 0 [11 0 66 -22 99 123]  
(i,j) = (2 2) 66 moves past -22 [11 0 -22 66 99 123]  
Outer loop: 3  
(i,j) = (3 0) 11 moves past 0 [0 11 -22 66 99 123]  
(i,j) = (3 1) 11 moves past -22 [0 -22 11 66 99 123]  
Outer loop: 4  
(i,j) = (4 0) 0 moves past -22 [-22 0 11 66 99 123]  
*** Finished Sort At 2024-04-07 21:52:35.631672 -0400 EDT m=+0.000602181  
*** Array after sorting: [-22 0 11 66 99 123]  

==> Sorted array:  [-22 0 11 66 99 123]  

Generate an array of random numbers...  
[63 39 -13 92 -8 89 61 -78 -3 14 79 -12 -73 -93 -66 -46 -51 -54 22 45 -69 -57 47 -53 5 98 -41 1 -78 60 64 18 -73 35 47 -90 54 34 -13 96]  

*** Array before sorting: [63 39 -13 92 -8 89 61 -78 -3 14 79 -12 -73 -93 -66 -46 -51 -54 22 45 -69 -57 47 -53 5 98 -41 1 -78 60 64 18 -73 35 47 -90 54 34 -13 96]  
*** Starting Sort At 2024-04-07 21:52:35.6317 -0400 EDT m=+0.000629926  
*** Finished Sort At 2024-04-07 21:52:35.631711 -0400 EDT m=+0.000641219  
*** Array after sorting: [-93 -90 -78 -78 -73 -73 -69 -66 -57 -54 -53 -51 -46 -41 -13 -13 -12 -8 -3 1 5 14 18 22 34 35 39 45 47 47 54 60 61 63 64 79 89 92 96 98]  

==> Sorted array:  [-93 -90 -78 -78 -73 -73 -69 -66 -57 -54 -53 -51 -46 -41 -13 -13 -12 -8 -3 1 5 14 18 22 34 35 39 45 47 47 54 60 61 63 64 79 89 92 96 98]  

Author: Daniel R. South, April 2024

using System;
using System.Linq;

public class Program
{
  public static void Main()
  {
    var array = new int[] { 1, 12, -5, -6, 50, 3 };
    var tableCases = new[]
    {
      new {
        subLength = 4,
        expectedValue = 12.75
      },
      new {
        subLength = 5,
        expectedValue = 10.8
      },
      new {
        subLength = 6,
        expectedValue = 9.16667
      },
    };
    Console.WriteLine("*** MAX AVERAGE IN CONTIGUOUS SUBARRAYS *** \n");
    Console.WriteLine($"ARRAY: [{string.Join(',', array)}]\n");

    foreach (var @case in tableCases)
    {
      Console.WriteLine($"SUB MIN LENGTH (k): {@case.subLength}");

      var result = MaxAverage(array, @case.subLength, 0);

      Console.WriteLine($"MAX AVERERAGE: {result}");
      Console.WriteLine($"EXPECTED VALUE: {@case.expectedValue}\n");
    }
  }

  static double MaxAverage(int[] array, int k, double average)
  {
    if (array.Length < k) return average;
    var subAverage = MaxAverageTaking(array, k, average);
    return MaxAverage(array.Skip(1).ToArray(), k, subAverage > average ? subAverage : average);
  }

  static double MaxAverageTaking(int[] array, int k, double average)
  {
    if (array.Length < k) return average;
    var subAverage = array.Average();
    // Console.WriteLine($"Average for [{string.Join(',', array)}] is: {subAverage}");
    return MaxAverageTaking(array.Take(array.Length - 1).ToArray(), k, subAverage > average ? subAverage : average);
  }
}

// Given an array consisting of n integers, 
// find the contiguous subarray whose length is greater than or equal to k that has the maximum average value. 
// And you need to output the maximum average value.

// Example:

// Input: [1,12,-5,-6,50,3], k = 4
// Output: 12.75
// Explanation:
// when length is 5, maximum average value is 10.8,
// when length is 6, maximum average value is 9.16667.
// Thus return 12.75.
// Note:

// 1 <= k <= n <= 10,000.
// Elements of the given array will be in range [-10,000, 10,000].
// The answer with the calculation error less than 10-5 will be accepted.
/*
Valid contiguous subarrays for k=4
1,12,-5,-6,50,3
1,12,-5,-6,50
1,12,-5,-6

12,-5,-6,50,3
12,-5,-6,50

-5,-6,50,3
*/
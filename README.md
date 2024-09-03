# HopperTest

HopperTest is a Go program that calculates the optimal number of hops needed for a hopper to move from a starting point to an endpoint on a grid while avoiding obstacles.

## Requirements

- Go 1.16 or later

## How to Run

1. **Clone the repository**:

   ```bash
   git clone https://github.com/mericdeser1/HopperTest.git
   cd HopperTest



## How to Run

2. **Prepare the input file**:

   Ensure you have an `input.txt` file in the `HopperTest` directory containing your test cases. Below is an example format for `input.txt`:

   ```
   2
   5 5
   4 0 4 4
   1
   1 4 2 3
   3 3
   0 0 2 2
   2
   1 1 0 2
   0 2 1 1
   ```

3. **Run the program**:

   Use the terminal to run the program:

   ```bash
   go run main.go
   ```

   This command will read the input from `input.txt`, compute the results, print them to the console, and also write them to `output.txt`.

## Author

**Meriç Deşer** - [mericdeser1](https://github.com/mericdeser1)

Feel free to reach out for any questions.

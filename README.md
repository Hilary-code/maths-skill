# Math Skills

This Go program analyzes statistical data from a text input and calculates various metrics such as Average, Median, Variance, and Standard Deviation stored in a file named `data.txt`.

## Features
- Calculate Average
- Calculate Median
- Calculate Variance
- Calculate Standard Deviation
- Handles overflow gracefully

## Prerequisites
- Go installed on your machine

## Installation
1. Clone the repository:
   ```bash
    git clone https://github.com/Hilary-code/maths-skill.git

 2. Navigate to the project directory:
    ```bash
    cd math-skills
# Usage

Run the script with `go build` then   `./math-skills data.txt`, then execute the program.
## How it Works

1. **Input Parsing:**
    - The program parses text input from the command-line arguments using `os.Args[1]`.

2. **Reader:**
    - The `reader` function reads data from a file named `data.txt` and passes it to the calculation functions.
    - After calculation, the results are printed to the console, displaying the analyzed statistical data.

3. **Calculation:**
    - Calculation for Average, Median, Variance, and Standard Deviation takes place within this module.
    - These functions are called from the main program to execute the analysis.


The program will output the statistics in the following manner (the following numbers are only examples).
    
    User$ `./math-skills data.txt`
    Average: 234
    Median: 146
    Variance: 769
    Standard Deviation: 45
All values are rounded to the nearest integer.    


## File Structure

- `main.go`: Main source code file
- `calculation/calculations.go`: Contains modules for average, median, variance, and standard deviation calculations
- `reader/`: Contains `reader.go` which reads data input from `os.Args[1]` and prints the results.

## Contribution Guidelines:

Contributions to this project are encouraged and welcomed! Here's how you can get involved:

1. **Fork the Project:** Interested in making substantial changes or experimenting with new features? Fork the project and start working on your own branch.

2. **Create a Branch:** Once you've forked the project, create a new branch for your feature or bugfix. Use a clear and descriptive name for your branch (`e.g., feature/AmazingFeature or bugfix/FixIssue123`).

3. **Commit Your Changes:** After making your modifications, commit your changes to your branch with a descriptive message summarizing what you've done.

4. **Push to the Branch:** Finally, push your changes to your branch to share them with the community.

5. **Create a Pull Request:** If you have changes you'd like to propose, simply create a pull request with your modifications. Whether they're bug fixes, feature enhancements, or other improvements.

6. **Open an Issue:** Noticed a bug or have a suggestion for making this project the best it can be.

Thank you for considering contributing to our project! We look forward to seeing your contributions.

## License

This project is licensed under the MIT License.


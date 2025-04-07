# euler-go

[![Go Version](https://img.shields.io/badge/Go-1.XX+-blue.svg)](https://golang.org/) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT) 

Knocking out two birds with one stone by learning Go and solving Project Euler Problems. This repository contains my personal solutions to these problems, developed as an exercise in learning the Go programming language. Currently, solutions aim to use only the Go standard library.

## What is Project Euler?

[Project Euler](https://projecteuler.net/) is a website dedicated to a series of challenging computational problems intended to be solved with computer programs. The problems range in difficulty and primarily involve mathematics and computer science concepts. It's a popular platform for learning new programming languages and sharpening problem-solving skills.

## What is Go?

[Go](https://golang.org/) (often referred to as Golang) is an open-source programming language developed at Google. It is known for its simplicity, efficiency, strong standard library, and built-in support for concurrency. These features make it well-suited for a variety of tasks, from web development and systems programming to data science and, in this case, solving computational problems.

## Project Goals

* **Learn Go:** Gain proficiency in Go syntax, idioms, and standard library usage.
* **Solve Euler Problems:** Systematically work through Project Euler problems.
* **Standard Library Focus:** Explore the capabilities of Go's standard library by using it exclusively for solutions (where feasible).
* **Clean Code:** Write clear, readable, and reasonably efficient Go code.

## Project Structure

The project is organized as a single Go module.

* `main.go`: The entry point for running solutions. It parses the command-line argument to determine which problem's solution to execute.
* `problems/`: Contains the Go code implementing the solutions for individual problems. Each problem has a dedicated function `SolveProblem5()`.

## Dependencies

* **Go:** A recent version of the Go compiler and tools (e.g., Go 1.18+).
* **Go Standard Library:** No external Go packages are required.

## How to Run

Ensure you have Go installed and configured on your system.

1.  **Clone the repository:**
    ```shell
    git clone https://github.com/Washiil/euler-go.git
    cd euler-go
    ```

2.  **Run a specific problem solution:**
    Use `go run main.go` followed by the problem number you wish to execute.

    ```shell
    foo@bar:~/euler-go$ go run main.go <problem_number>
    ```

    **Example (running Problem 5):**
    ```shell
    foo@bar:~/euler-go$ go run main.go 5
    ```

    **Example Output (solution for Problem 5):**
    ```shell
    Problem [5] : 777777
    ```

## Adding New Solutions

1.  Implement the logic for the new problem (e.g., `problem_N.go` containing a function `SolveProblemN()`).
2.  Update `main.go` to include a case or mapping for the new problem number, directing it to call your new `SolveProblemN()` function.
3.  Ensure the new function prints the result in the standard format: `Problem [N] : <answer>`.
4.  Test your solution!

## Contributing

This is primarily a personal learning project. However, suggestions, feedback, or bug reports are welcome! Please feel free to open an Issue on GitHub. Pull requests might be considered if they align with the project's goals (learning Go, standard library usage, clarity).

## License

This project is currently unlicensed and intended for personal educational purposes.
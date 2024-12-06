# Simple Calculator (CLI)

## Description
A command-line tool that performs basic arithmetic operations like addition, subtraction, multiplication, and division. This project is designed to help learn Go programming through building a simple yet practical application.

---

## Features
- Supports four basic operations:
  - Addition (`+`)
  - Subtraction (`-`)
  - Multiplication (`*`)
  - Division (`/`)
- Validates user input for correctness (e.g., ensures numbers are valid).
- Provides clear error messages for invalid inputs or unsupported operations.

---

## Getting Started

### Prerequisites
- Go (latest stable version recommended)
- A terminal or command prompt

### Installation
1. Clone the repository:
   ```bash
   git clone https://git-repo/your-username/simple-calculator.git
   cd simple-calculator
   ```
2. Initialize the Go module:
   ```bash
   go mod init simple-calculator
   ```

---

## Usage

### Running the Calculator
Use the following format to run the calculator:
```bash
go run main.go <number1> <operator> <number2>
```

### Examples
1. Addition:
   ```bash
   go run main.go 10 + 5
   ```
   Output:
   ```
   The result of 10 + 5 is 15
   ```

2. Subtraction:
   ```bash
   go run main.go 15 - 3
   ```
   Output:
   ```
   The result of 15 - 3 is 12
   ```

3. Multiplication:
   ```bash
   go run main.go 7 * 8
   ```
   Output:
   ```
   The result of 7 * 8 is 56
   ```

4. Division:
   ```bash
   go run main.go 20 / 4
   ```
   Output:
   ```
   The result of 20 / 4 is 5
   ```

### Error Handling
- If fewer than 3 arguments are provided:
  ```bash
  go run main.go
  ```
  Output:
  ```
  Usage: go run main.go <number1> <operator> <number2>
  ```

- If the operator is unsupported:
  ```bash
  go run main.go 10 % 3
  ```
  Output:
  ```
  Unsupported operator: %
  ```

---

## Roadmap
- [ ] Add support for advanced operations (e.g., modulus, power).
- [ ] Handle interactive input (prompting the user instead of command-line arguments).
- [ ] Add unit tests for all features.

---

## License
This project is licensed under the MIT License.

---

## Contributing
Feel free to fork the repository and submit pull requests for improvements or bug fixes!

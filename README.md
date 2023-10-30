# Go Hello World Module

## Introduction

Welcome to the Go Hello World Module! If you're new to Go and looking to create your first Go project, you're in the right place. In this project, we've created a simple Go module that greets you based on the time of day, and we'll guide you through the process of understanding how it's done.

## Using Existing Go Module/Project

To use and modify the code from this project, follow these steps:

1. **Clone the Repository**: Start by cloning this project's repository. Open your terminal or command prompt and navigate to the directory where you want to store the project. Then, run the following command:

   ```bash
   git clone https://github.com/AbdullahCheema35/go-hello-world-module.git
   ```

2. **Clone the Repository**: After cloning, navigate into the project directory:

  ```bash
   cd go-hello-world-module
  ```

1. **Running the Project**: To run the project and see the code in action, use the following command:

  ```bash
  go run main.go
  ```

  This command executes the program and prints Hello World along with the greeting message.



## Creating Your Own Go Project/Module

### Step 1: Go Installation

Before you start, make sure you have Go installed on your system. If not, you can download and install it from the [official website](https://golang.org/doc/install).

### Step 2: Initialize Your Go Module

To create your own Go project, follow these steps:

1. Create a new directory anywhere on your system.

2. Navigate into the directory and open a terminal or command prompt in that directory.

3. Execute the following command to initialize your Go module:

```bash
go mod init <your module name>
```
Make sure to replace  ```<your module name>``` with your preferred module name (preferably your directory name).

### Step 3: Getting Started with Your Go Project

Before diving into your Go programming journey, follow these steps to get started and create your first Go project from scratch, including setting up packages, naming modules, and running your code.

1. **Create a File**: Start by creating a new file and name it `main.go`. This is where your Go program will begin.

2. **Set the Package**: At the top of your `main.go` file, write `package main`. This tells Go that this file is the entry point of your program. Also, don't forget to define the ```main()``` function in this file.

3. **Creating Other Packages**: If you want to create additional packages in your project (these are like separate parts of your project), create a new directory (folder) within your project directory and give it a name that represents your new package.

4. **Inside Package Files**: In that new package's directory, create Go source files (with a `.go` extension). In each of these files, you should start with `package new_package_name`, where `new_package_name` is the name of your new package.

5. **Exported Functions**: For functions that you want to call from other packages, make sure to begin their names with a capital letter. This is important in Go.

6. **Adding Imports**: If you're using code from other packages in your module, you need to tell Go that you're using them. To do this, use the `import` statement at the beginning of your Go file. If the code comes from another package in your project, you can import it using `"module_name/packagename"`.

7. **Keep go.mod Updated**: Your project uses the `go.mod` file to keep track of all the packages and dependencies it needs. To make sure your `go.mod` file is up to date, open your terminal or command prompt, navigate to the directory where your project is, and run the command `go mod tidy`.

8. **Running Your Program**: To run your program, use the command `go run main.go` in your terminal. This tells Go to start your program from the `main.go` file.

In summary, you create a Go program by creating a `main.go` file, organizing code into different packages if needed, and following some conventions, like capitalizing the first letter of functions you want to use elsewhere. You also make sure to import other packages and keep your `go.mod` file updated. Finally, you can run your program using `go run main.go`. Enjoy your coding journey in Go!

- You're all set to go. Gophers!

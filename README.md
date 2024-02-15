



# Dining Philosophers Problem

This is a solution to the classic Dining Philosophers Problem implemented in Go. The problem involves five philosophers sitting at a table and dining together, with each philosopher requiring two forks to eat.

## Overview

This project is part of a series of short CLI-based programs inspired by classic computer science problems. This series is to demonstrate using Goroutines and Go concurrency features to write solutions to these classic problems.

## Table of Contents

- [Introduction](#introduction)
- [Philosophers](#philosophers)
- [Code Details](#code-details)
- [Running the Program](#running-the-program)
- [Testing](#testing)

## Introduction

The Dining Philosophers Problem is a classic synchronization problem in computer science. The goal is to design a concurrent algorithm to allow the philosophers to alternate between eating and thinking, preventing deadlock and starvation.

## Philosophers

The philosophers in this implementation are named:
1. Jean-Paul Sartre
2. Simone de Beauvoir
3. Albert Camus
4. Martin Heidegger
5. Søren Kierkegaard

Each philosopher is represented by a struct with a name, a left fork, and a right fork.


## Code Details


-   **Philosopher Struct:**

    -   The `Philosopher` struct stores information about each philosopher, including their name, the index of their left fork, and the index of their right fork.
-   **Philosophers List:**

    -   A list of philosophers is defined, each with their specific left and right fork indices.
-   **Variables:**

    -   `hunger`: Number of times each philosopher eats until full.
    -   `eatTime`, `thinkTime`, `sleepTime`: Durations for eating, thinking, and sleep.
    -   `orderFinished`: Stores the order in which philosophers finish eating.
-   **Main Function:**

    -   The `main` function initiates the dining process, displaying welcome and end messages.
-   **Dine Function:**

    -   The `dine` function sets up variables and creates goroutines for each philosopher to simulate dining.
-   **Dining Problem Function:**

    -   The `diningProblem` function represents the behavior of an individual philosopher, including picking up and putting down forks, eating, and thinking.

## Example Output 

Here is an example of the terminal output after running the program:

```bash
$ go run .
Dining Philosophers Problem
---------------------------
The dining table is empty
Simone de Beauvoir is seated at the table.
Søren Kierkegaard is seated at the table.
Martin Heidegger is seated at the table.
Jean-Paul Sartre is seated at the table.
Albert Camus is seated at the table.
        Albert Camus takes the left fork.
        Søren Kierkegaard takes the left fork.
        Søren Kierkegaard takes the right fork.
        Søren Kierkegaard has both forks and has started eating.
        Søren Kierkegaard is thinking.
        Søren Kierkegaard has put down the forks.
        Søren Kierkegaard takes the left fork.
        Søren Kierkegaard takes the right fork.
        Søren Kierkegaard has both forks and has started eating.
        Martin Heidegger takes the left fork.
        Simone de Beauvoir takes the left fork.
        Søren Kierkegaard is thinking.
        Martin Heidegger takes the right fork.
        Martin Heidegger has both forks and has started eating.
        Martin Heidegger is thinking.
        Martin Heidegger has put down the forks.
        Søren Kierkegaard has put down the forks.
        Søren Kierkegaard takes the left fork.
        Søren Kierkegaard takes the right fork.
        Søren Kierkegaard has both forks and has started eating.
        Søren Kierkegaard is thinking.
        Søren Kierkegaard has put down the forks.
Søren Kierkegaard is satisfied.
Søren Kierkegaard left the table.
        Albert Camus takes the right fork.
        Albert Camus has both forks and has started eating.
        Albert Camus is thinking.
        Albert Camus has put down the forks.
        Simone de Beauvoir takes the right fork.
        Simone de Beauvoir has both forks and has started eating.
        Simone de Beauvoir is thinking.
        Simone de Beauvoir has put down the forks.
        Martin Heidegger takes the left fork.
        Martin Heidegger takes the right fork.
        Martin Heidegger has both forks and has started eating.
        Martin Heidegger is thinking.
        Martin Heidegger has put down the forks.
        Martin Heidegger takes the left fork.
        Martin Heidegger takes the right fork.
        Martin Heidegger has both forks and has started eating.
        Martin Heidegger is thinking.
        Martin Heidegger has put down the forks.
Martin Heidegger is satisfied.
Martin Heidegger left the table.
        Albert Camus takes the left fork.
        Albert Camus takes the right fork.
        Albert Camus has both forks and has started eating.
        Albert Camus is thinking.
        Albert Camus has put down the forks.
        Albert Camus takes the left fork.
        Albert Camus takes the right fork.
        Albert Camus has both forks and has started eating.
        Albert Camus is thinking.
        Albert Camus has put down the forks.
Albert Camus is satisfied.
Albert Camus left the table.
        Jean-Paul Sartre takes the right fork.
        Jean-Paul Sartre takes the left fork.
        Jean-Paul Sartre has both forks and has started eating.
        Jean-Paul Sartre is thinking.
        Jean-Paul Sartre has put down the forks.
        Simone de Beauvoir takes the left fork.
        Simone de Beauvoir takes the right fork.
        Simone de Beauvoir has both forks and has started eating.
        Simone de Beauvoir is thinking.
        Simone de Beauvoir has put down the forks.
        Jean-Paul Sartre takes the right fork.
        Jean-Paul Sartre takes the left fork.
        Jean-Paul Sartre has both forks and has started eating.
        Jean-Paul Sartre is thinking.
        Jean-Paul Sartre has put down the forks.
        Simone de Beauvoir takes the left fork.
        Simone de Beauvoir takes the right fork.
        Simone de Beauvoir has both forks and has started eating.
        Simone de Beauvoir is thinking.
        Simone de Beauvoir has put down the forks.
Simone de Beauvoir is satisfied.
Simone de Beauvoir left the table.
        Jean-Paul Sartre takes the right fork.
        Jean-Paul Sartre takes the left fork.
        Jean-Paul Sartre has both forks and has started eating.
        Jean-Paul Sartre is thinking.
        Jean-Paul Sartre has put down the forks.
Jean-Paul Sartre is satisfied.
Jean-Paul Sartre left the table.
The dining table is empty

```

## Running the Program

To run the program, execute the following command:

```bash
go run main.go
```

The program will display the progress of the philosophers as they dine.

## Testing

The project includes a test file (`main_test.go`) to verify the correctness of the implementation. To run the tests, use the following command:

```bash
go test
```
Be sure to comment out this section of main.go first in order to run the tests effectively:
```
    eatTime = 0 * time.Second
	sleepTime = 0 * time.Second
	thinkTime = 0 * time.Second
```
The tests include scenarios with zero delays and varying delays for eating, sleeping, and thinking.

---

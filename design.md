# Real-Time Operating System (RTOS) Simulator

The RTOS simulator is a command-line application that simulates a real-time operating system. It manages tasks with different priorities and runs them according to a priority-based scheduling algorithm.

## Components

1. Task: A basic unit of execution with a priority and a duration.
2. Scheduler: Responsible for managing tasks and running them according to their priorities.
3. Command Line Interface (CLI): A user interface to interact with the simulator, add tasks, and start the simulation.

## Scheduling Algorithm

The simulator uses a priority-based scheduling algorithm. Tasks with higher priorities run before tasks with lower priorities. Tasks with the same priority run in a first-come, first-served order.

## Usage

To build and run the application, use the following commands:

go build -o rtos_simulator
./rtos_simulator

To add tasks and run the scheduler, use the following commands:

./rtos_simulator add -p <priority> -d <duration>
./rtos_simulator run


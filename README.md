# Git control

It is a CLI application that selects the 10 most used repositories globally and can be filtered by the user from a CSV file.

## About The Project
An algorithm that will give each repository an activity score.
The NumberOfResults controls the number of results.

```bash
const NumberOfResults = 10
```
In this application, the score was determined by the number of commits.  
it is possible to identify the repository with the most commits and in which 10 repositories a given user made the most commits.

## Getting Started

```bash
# Clone the repository
$ git clone https://github.com/marielesf/git-control.git
$ cd git-control
$ go mod tidy
$ go run main.go
```

## Usage
### 1 - 10 Repositories with the most commits
![image](https://github.com/user-attachments/assets/df361142-cebe-4398-b4cc-767f7968718a)
### 2 - 10 Repositories most Used by an informed user
![image](https://github.com/user-attachments/assets/7a278d1a-a646-4db1-9b87-5e45bdda5e8c)

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)
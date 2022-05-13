# Go-Shop Auto Parts
- Project attempt for the **Qlik Hackathon May 2022**

## Description
The **Go-Shop Auto Parts** is an American retailer of aftermarket automotive parts and accessories, the second largest in the United States. Founded in 2022. <br>
The **Go-Shop CLI** is a CLI program used by cashiers in the Go-Shop car parts to help them:
- Put item's name and price
- Add the payment
- Calculate the total of purchase, and the change to return
- View and save a receipt 

When the cashier finishes the process, a receipt with the client's name is saved in a `.txt` file under the `/receipts` folder.

## Worklflow
![](/docs/diagrams/UML%20Diagram.png)

## Development
### Prerequisites ###
1. You need [Go](https://go.dev/doc/install) installed

### Run program ###
```
cd src
go run main.go receipt.go
```
### Testing ###
```
go test ./src/
```

## knowledge acquired ##
Learn more in-depth the **Go** language:
- Arrays & Slices
- Maps
- Pointers
- Multiple returns
- Structs
- Receiver functions
- Loops
- Formatting strings
- `fmt` package
- testing in Go

# Shopping Mall Checkout

This repo contains the code behind powering the checkout system at the local shopping mall.

Below are the instructions on how to run the app and how some of the files work

- [Running the app](#running-the-app)
- [How the file structure works](#how-the-file-structure-works)

---

- `products.csv` - This file contains the Product Code, Name, and Price of a certain product
- `discounts.csv` - This file contains any discounts to be applied to products.

**Explanation**:
If a product is buy one get one free then the discount applies as soon as you have two items or more therefore it will apply from the start

## Running the App

To run the app you can just run the binary in your console. In the root of the project just do `./shopping-mall`. 

Or if you have [Go](https://golang.org) installed you go use the standard Go commands like `run` or `build`
```
go run ./cmd/cli/main.go --products=./files/products.csv --discounts=./files/discounts.csv
```

We have 2 flags to customise the location of where the `products.csv` file and the `discounts.csv` file found the defaults are listed below.
```
Usage of ./shopping-mall:
  -discount
    	location of the discount list(default "files/discounts.csv")
  -product
    	location of the product list (default "files/products.csv")
```  

## How the file structure works

This project comes with 2 `csv` files for reading in a list of products and a list of discounts. The structure of these files are important for how the information
is read into the app so we've laid it out below

`products.csv` file structure

| Product Code | Product Name | Price |
| ------------ | ------------ | ----- |
| FR1          | Fruit Tea    | 3.11  |
| SR1          | Strawberries | 5.00  |
| CF1          | Coffee       | 11.23 |

Based on the table above this is how the CSV file should be structured
```
FR1,Fruit Tea,3.11
SR1,Strawberries,5.00
CF1,Coffee,11.23
```

`discounts.csv` file structure

| Product Code | Price Deduction | Buy one get one free offer? | How many items till the discount applied | Discount Description                       |
| ------------ | --------------- | --------------------------- | ---------------------------------------- | ------------------------------------------ |
| FR1          | 3.11            | yes                         | 0                                        | Buy one get on free                        |
| SR1          | 0.50            | no                          | 3                                        | Buy 3 or more for a reduced price of £4.50 |

Based on the table above this is how the CSV file should be structured
```
FR1,3.11,yes,0,Buy one get one free
SR1,0.50,no,3,Buy 3 or more for a reduced price of £4.50
```
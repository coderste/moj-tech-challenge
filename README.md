# Shopping Mall Checkout

This repo contains the code behind powering the checkout system at the local shopping mall.

Below are the instructions on how to run the app and how some of the files work

---

- `products.csv` - This file contains the Product Code, Name, and Price of a certain product
- `discounts.csv` - This file contains any discounts to be applied to products. The following struct is setup to help the system read in the discounts

| Product Code | Price Deduction | Buy one get one free offer? | How many items till the discount applied |
| ------------ | --------------- | --------------------------- | ---------------------------------------- |
| FR1          | 3.11            | yes                         | 0                                        |
| SR1          | 0.50            | no                          | 3                                        |

**Explanation**:
If a product is buy one get one free then the discount applies as soon as you have two items or more therefore it will apply from the start 

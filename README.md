# E-Commerce API
i make this api using golang with gin

## How to use this
Please create database go-ecommerce, or if u not it will error. For table, it will auto add(auto migrate) automatic.

And after that, download this repo, and copy this text, and run in terminal. and its done.

    go run main.go

## Implementation in android 
https://github.com/achmadrizkin/kotlin_e-commerce_MVVM

## API 

please change the localhost:3000 to 10.0.2.2:3000 if u want to use emulator

* List of All Products:
    * POST http://localhost:3000/v1/products/
    * GET http://localhost:3000/v1/products/
    * GET BY ID http://localhost:3000/v1/products/:id
    * GET BY Category http://localhost:3000/v1/products/category/:category
    * GET BY User Email http://localhost:3000/v1/products/user/:user_email
    * PUT http://localhost:3000/v1/products/:id
    * DELETE http://localhost:3000/v1/products/:id

* List of Transaction:
    * POST http://localhost:3000/v1/transaction/
    * GET http://localhost:3000/v1/transaction/
    * GET BY ID http://localhost:3000/v1/transaction/:id
    * PUT http://localhost:3000/v1/transaction/:id
    * DELETE http://localhost:3000/v1/transaction/:id

* List of Products Hoodie:
    * POST http://localhost:3000/v1/products/hoodie/
    * GET http://localhost:3000/v1/products/hoodie/
    * GET BY ID http://localhost:3000/v1/products/hoodie/:id
    * PUT http://localhost:3000/v1/products/hoodie/:id
    * DELETE http://localhost:3000/v1/products/hoodie/:id

* List of Products Book:
    * POST http://localhost:3000/v1/products/book/
    * GET http://localhost:3000/v1/products/book/
    * GET BY ID http://localhost:3000/v1/products/book/:id
    * PUT http://localhost:3000/v1/products/book/:id
    * DELETE http://localhost:3000/v1/products/book/:id

* List of Products Laptop:
    * POST http://localhost:3000/v1/products/laptop/
    * GET http://localhost:3000/v1/products/laptop/
    * GET BY ID http://localhost:3000/v1/products/laptop/:id
    * PUT http://localhost:3000/v1/products/laptop/:id
    * DELETE http://localhost:3000/v1/products/laptop/:id

## API Response Products
All data have same response, the difference is just All Products have GET DATA BY UserEmail, and Category. i make this so u can use it more easy. but, if use use GET DATA BY ID, it just return data by what u are search.

    {
        "data": [
            {
                "id": 10,
                "name_product": "Hoodie Overpriced",
                "image_url": "https://s1.bukalapak.com/img/19067587152/large/data.jpeg",
                "description": "this Hoodie was ...",
                "price": 23000,
                "name_user": "rizki",
                "email_user": "iamhandsome@gmail.com"
            },
            {
                "id": 11,
                "name_products": "Sherlock Holmes",
                "image_url": "https://s1.bukalapak.com/img/19067587152/large/data.jpeg",
                "description": "Best book in the world",
                "price": 12000,
                "name_user": "kevin",
                "email_user": "kevinisnothandsome@gmail.com"
            }
        ]
    }   
    
## API POST
this is how to post. Just add category if u use All Products

    {
        "name_product": "Hoodie Overpriced",
        "image_url": "https://s1.bukalapak.com/img/19067587152/large/data.jpeg",
        "description": "this Hoodie was ...",
        "price": 23000,
        "name_user": "rizki",
        "email_user": "iamhandsome@gmail.com"
    }

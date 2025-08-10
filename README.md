# go-mvc-assignment

- create a copy of .env.sample.
```bash
    cp .env.sample .env
```
- Fill in with proper details.

- Run the `main.go` file

```bash
    go run cmd/main.go
```

- Know the Live routes and more on Postman : Click [here](https://web.postman.co/workspace/My-Workspace~b9d5d650-5956-4307-8b7a-a5e57a5c5bb7/collection/44035623-0d31ba9a-354e-4b6a-8bb1-3510bfea7dc3?action=share&source=copy-link&creator=44035623).

# Know the API

## `POST`

### Account Creation : **POST** `/account/signup`
- Content-Type: `application/x-www-form-urlencoded` 
- **Body Params**
    - `user_name` : string
    - `name` : string
    - `pwd` : string
    - `re_pwd` : string
    - `role` : string (admin, cook or customer)

### User Login : **POST** `/account/login`
- Content-Type: `application/x-www-form-urlencoded` 
- **Body Params**
    - `user_name` : string
    - `password` : string

### Add Item : **POST** `/api/item`
- Content-Type: `application/x-www-form-urlencoded` 
- **Body Params**
    - `item_name` : string
    - `cook_time_min` : int (optional)
    - `price` : int
    - `cat_id` : int
    - `subcat_id` : int (optional)

### Add Category :  **POST** `/api/categories`
- Content-Type: `application/x-www-form-urlencoded` 
- **Body Params**
    - `cat_name` : int
    - `cat_description` : int (optional)

### Place Order : **POST** `/api/order`
- Content-Type: `application/json` 
- **Expected Input format**
```json
[
    {
        "item_id": 1,
        "quantity": 1,
        "total_price": 100,
        "instruction": "some instruction"
    },
    {
        "item_id": 5,
        "quantity": 2,
        "total_price": 25,
        "instruction": ""
    },
    {
        "item_id": 12,
        "quantity": 1,
        "total_price": 10,
        "instruction": "no onions"
    }
]
```
### Kitchen Actoins : **POST** `/api/cook`
- to change the status of an order
- Content-Type: `application/x-www-form-urlencoded` 
- **Body Params**
    - `order_id` : int
    - `item_id` : int 
    - `is_complete` : string ("taken" or "complete")



### Order Payment : **POST** `/api/pay`
- to make payment for an order
- Content-Type: `application/x-www-form-urlencoded` 
- **Body Params**
    - `order_id` : int
    - `customer_review` : int 

### User Logout : **POST** `/account/logout`
- Logout

## `DELETE`

### Delete User : **DELETE** `/api/user/{userID}`
- Deletes user by userID

### Delete Item from Menu **DELETE** `/api/item/{item_id}`
- Content-Type: `application/x-www-form-urlencoded` 
- **Query Params**
    - `item_id` : int

## `GET` 
### Get User List : **GET** `/api/users`
- Displays all the users 
- Expected Output
```json
[
    {
        "user_id": "2154fc56-7543-11f0-9486-6c0f08f8c4ef",
        "user_name": "haldi",
        "name": "adi",
        "pwd_hash": "",
        "role": "admin"
    }
]
```

### Get Specific User info : **GET** `/api/user/{userID}`
- Displays user by userID
- Expected Output
```json
{
    "user_id": "2154fc56-7543-11f0-9486-6c0f08f8c4ef",
    "user_name": "haldi",
    "name": "adi",
    "pwd_hash": "",
    "role": "admin"
}
```

### Display Menu : **GET** `/api/item`
- Displays Menu
- Expected Output
```json
[
    {
        "item_id": "19",
        "item_name": "Fafda Jalebi",
        "cook_time_min": "20",
        "price": "110",
        "display_pic": {
            "String": "",
            "Valid": false
        },
        "cat_id": "3",
        "category": "Gujarati",
        "subcat_id": "8",
        "subcategory": "Sweet"
    },
    {
        "item_id": "42",
        "item_name": "Pani Puri",
        "cook_time_min": "10",
        "price": "50",
        "display_pic": {
            "String": "",
            "Valid": false
        },
        "cat_id": "9",
        "category": "Tangy",
        "subcat_id": "13",
        "subcategory": "Street Food"
    }
]
```

### Display Categories : **GET** `/api/categories`
- Displays Categories
- Expected Output
```json
[
    {
        "cat_id": 1,
        "cat_name": "South Indian",
        "cat_description": {
            "String": "Traditional dishes from South India including dosas, idlis, sambars and coconut-based curries",
            "Valid": true
        }
    },
    {
        "cat_id": 2,
        "cat_name": "North Indian",
        "cat_description": {
            "String": "Rich and flavorful dishes from North India with creamy gravies, rotis and tandoor items",
            "Valid": true
        }
    }
]
```

### Cook Dashboard : **GET** `/static/cook`
- Expected Output
```json
[
    {
        "order_id": 1,
        "item_id": 1,
        "item_name": "Masala Dosa",
        "quantity": 1,
        "instruction": {
            "String": "some instruction",
            "Valid": true
        },
        "is_complete": "pending",
        "cook_id": {
            "String": "",
            "Valid": false
        },
        "table_no": {
            "Int64": 1,
            "Valid": true
        },
        "order_at": "2025-08-09T23:14:40+05:30"
    },
    {
        "order_id": 1,
        "item_id": 5,
        "item_name": "Coconut Rice",
        "quantity": 2,
        "instruction": {
            "String": "",
            "Valid": true
        },
        "is_complete": "pending",
        "cook_id": {
            "String": "",
            "Valid": false
        },
        "table_no": {
            "Int64": 1,
            "Valid": true
        },
        "order_at": "2025-08-09T23:14:40+05:30"
    }
]
```

### All Orders : **GET** `/static/admin`
- Displays all Orders (default current date)
- Expected Output
```json
[
    {
        "order_id": 1,
        "order_at": "2025-08-09T23:14:40+05:30",
        "table_no": {
            "Int64": 1,
            "Valid": true
        },
        "customer_id": "614e1ae0-7548-11f0-9486-6c0f08f8c4ef",
        "customer_name": "cust",
        "status": "received",
        "total_price": 135
    }
]
```

### All Orders by Date : **GET** `/static/admin/{date}`
- Displays all Orders of a Date
- **Query Params**
    - `date` : date string (fomat yyyy-mm-dd eg: 2025-08-09)
- Expected Output
```json
[
    {
        "order_id": 1,
        "order_at": "2025-08-09T23:14:40+05:30",
        "table_no": {
            "Int64": 1,
            "Valid": true
        },
        "customer_id": "614e1ae0-7548-11f0-9486-6c0f08f8c4ef",
        "customer_name": "cust",
        "status": "received",
        "total_price": 135
    }
]
```

### Customer Order Display : **GET** `/static/order`
- Displays orders by customer (default for the same day)
- Expected Output
```json
[
    {
        "order_id": 1,
        "order_at": "2025-08-09T23:14:40+05:30",
        "table_no": {
            "Int64": 1,
            "Valid": true
        },
        "customer_id": "614e1ae0-7548-11f0-9486-6c0f08f8c4ef",
        "customer_name": "cust",
        "status": "received",
        "total_price": 135
    }
]
```

### Customer Orders by Date : **GET** `/static/order/{date}`
- Displays user's odrders by date.
- **Query Params**
    - `date` : date string (fomat yyyy-mm-dd eg: 2025-08-09)
- Expected Output
```json
[
    {
        "order_id": 1,
        "order_at": "2025-08-09T23:14:40+05:30",
        "table_no": {
            "Int64": 1,
            "Valid": true
        },
        "customer_id": "614e1ae0-7548-11f0-9486-6c0f08f8c4ef",
        "customer_name": "cust",
        "status": "received",
        "total_price": 135
    }
]
```
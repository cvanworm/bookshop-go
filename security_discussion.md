## Problems:
- `IsPOShipped` was querying the Books table instead of PurchaseOrders
- Database Error:
    - Need to wrap scan in a for loop when multiple results could match
    - Ensure multiple entries of the same customer and book cannot be added to PurchaseOrders

## Security:
- It is possible to change the balance of someone's account.
    -This information should only be changed if a book is ordered.
- You can see and potentially take all shipping information.
    - This information should be made private to the bookstore business only. Not everyone should have access to this information.
- Multiple books of same name and author can be added with different price. This can be an issue when trying to get the correct price.
    - Maybe add a check to determine if a book is already in the system. If the price is different, create a function that uses a PUT request to reach out and change the price of the book (This must only be availible to the book store, so the prices cannot be changed willingly by anyone).

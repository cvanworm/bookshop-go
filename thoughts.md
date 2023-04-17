## Problems:
- `IsPOShipped` was querying the Books table instead of PurchaseOrders
- Database Error:
    - Need to wrap scan in a for loop when multiple results could match
    - Ensure multiple entries of the same customer and book cannot be added to PurchaseOrders

## Security:



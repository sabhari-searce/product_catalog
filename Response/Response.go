package response

var Marshal_error string = "Marshal error in conversion!"
var ProductInErr string = "Product not inserted"
var ProductInDone string = "Insert to Product done"
var ProductIdErr string = "Invalid product id found as a input"
var ProductNameErr string = "Invalid product name found as a input"
var CategoryIdErr string = "Invalid category id found as a input"
var PriceErr string = "Invalid price found as a input"
var GetProductErr map[string]string = map[string]string{"response": "NO DATA FOUND!"}
var WritingErr string = "Error on writing response to screen"
var ProductGetErr string = "Error in getting the specified product"
var GetRowErr string = "Error in getting the next row"
var AtoiErr string = "Error in converting the string to int with atoi"
var ProductDelErr map[string]string = map[string]string{"response": "DATA NOT FOUND FOR DELETING!!"}
var ProductDel map[string]string = map[string]string{"response": "DELETE ON PRODUCT DONE"}
var ProductUpErr string = "Error in updating the product table"
var ProductKeyErr string = "Entered key not found in product table for updating"
var ProductUp map[string]string = map[string]string{"response": "UPDATED ON PRODUCT DONE"}
var QuantityNeg string = "Quantity cannot be negative"
var InventoryIn string = "Insert to Inventory done"
var InventoryInErr string = "Error in inserting the record to Inventory table"
var InventoryGetErr string = "Error occured in getting inventory"
var InventoryDel map[string]string = map[string]string{"response": "DELETE ON INVENTORY DONE!!"}
var InventoryDelErr map[string]string = map[string]string{"response": "Error occured in deleting to inventory"}
var UpdateInvIdErr map[string]string = map[string]string{"response": "ENTERED ID NOT FOUND FOR UPDATING"}
var UpdateInvErr string = "Error in updating the inventory table"
var UpdateInventoryDone map[string]string = map[string]string{"response": "UPDATE ON INVENTORY DONE"}
var UpdateInvKey string = "Entered key not found in inventory table for updating"
var CategoryNameErr string = "Invalid category name found as an input"
var CategoryInErr string = "Error in inserting record to category table"
var CategoryIn string = "Insert to Category done"
var CategoryGetErr string = "Error in getting the category"
var CategoryDelNotFound map[string]string = map[string]string{"response": "DATA NOT FOUND FOR DELETING!!"}
var CategoryDelError string = "Error in deleting the content in category"
var CategoryDel map[string]string = map[string]string{"response": "DELETE ON CATEGORY DONE!!"}
var CategoryIdUpdateErr map[string]string = map[string]string{"response": "ENTERED ID NOT FOUND FOR UPDATING"}
var CategoryUpdate map[string]string = map[string]string{"response": "UPDATED ON CATEGORY DONE!"}
var CategoryUpdateErr string = "Error on updating cateory table"
var CategoryUpdateKeyErr string = "The entered key not found in the category table for updating"
var CreateCartReferenceErr string = "Error in creating and inserting the reference to the reference table"
var CartInEmptyErr string = "EMPTY QUERY DATA FOUND FOR INSERTING IN CART"
var CartInInvalid string = "Invalid argument found as an input"
var ReferenceGetErr string = "Error in getting the specified reference"
var ReferenceNotFound string = "THE ENTERED REFERENCE ID NOT CREATED"
var ProductNotFound string = "The specified product id not found"
var CartItemGetErr string = "Error occured in getting the cart item"
var GetCartErr string = "Error in getting the specified cart item"
var RefErr string = "Invalid reference found as an input "
var DeleteCart map[string]string = map[string]string{"response": "DELETE ON CART DONE!!"}
var DeleteCartErr string = "Error in deleting the specified cart item"
var CartIdErr string = "Entered id not found in the cart table"
var UpdateCart = map[string]string{"response": "UPDATED ON CART DONE"}
var UpdateCartErr string = "Error in updating the cart"
var UpdateKeyCart string = "Entered key not found in the cart table"

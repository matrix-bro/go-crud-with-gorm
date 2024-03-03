package controllers

import (
	"example/go-crud/initializers"
	"example/go-crud/models"
	"example/go-crud/serializers"

	"github.com/gin-gonic/gin"
)

func CreateAuthor(c *gin.Context) {
	var authorSerializer serializers.AuthorSerializer

	if err := c.ShouldBindJSON(&authorSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	author := models.Author{FirstName: authorSerializer.FirstName, LastName: authorSerializer.LastName}

	result := initializers.DB.Create(&author)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error creating new author."})
		return
	}

	c.IndentedJSON(201, gin.H{"data": authorSerializer, "message": "Created new author successfully."})
}

func CreateBook(c *gin.Context) {
	var author models.Author
	err := CheckByID(c.Param("id"), &author)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var bookSerializer serializers.BookSerializer

	if err := c.ShouldBindJSON(&bookSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	newBook := models.Book{Title: bookSerializer.Title, Description: bookSerializer.Description, AuthorID: author.ID}
	result := initializers.DB.Create(&newBook)

	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error creating new book."})
		return
	}

	c.IndentedJSON(201, gin.H{"data": bookSerializer, "message": "Created new book successfully."})

}

func GetAllAuthors(c *gin.Context) {
	var authors []serializers.AuthorSerializer
	result := initializers.DB.Model(&models.Author{}).Find(&authors)
	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving authors."})
		return
	}
	c.IndentedJSON(201, gin.H{"data": authors, "message": "Author retrieved successfully."})
}

func GetAllBooks(c *gin.Context) {
	var books []serializers.BookSerializer
	result := initializers.DB.Model(&models.Book{}).Find(&books)
	if result.Error != nil {
		c.IndentedJSON(400, gin.H{"message": "Error retrieving books."})
		return
	}
	c.IndentedJSON(201, gin.H{"data": books, "message": "Books retrieved successfully."})
}

func GetAuthorDetails(c *gin.Context) {
	var author models.Author
	err := CheckByID(c.Param("id"), &author)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var authorDetails serializers.AuthorDetailsSerializer
	result := initializers.DB.Preload("Books").Find(&author).Error

	if result != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": result.Error()})
		return
	}

	authorDetails.FirstName = author.FirstName
	authorDetails.LastName = author.LastName

	for _, book := range author.Books {
		book_serializer := serializers.BookSerializer{
			Title:       book.Title,
			Description: book.Description,
		}

		authorDetails.Books = append(authorDetails.Books, book_serializer)
	}

	c.IndentedJSON(201, gin.H{"data": authorDetails, "message": "Author details retrieved successfully."})
}

func GetBookDetails(c *gin.Context) {
	var book models.Book
	err := CheckByID(c.Param("id"), &book)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var bookDetails serializers.BookDetailsSerializer
	result := initializers.DB.Preload("Author").Find(&book).Error

	if result != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": result.Error()})
		return
	}

	bookDetails.Title = book.Title
	bookDetails.Description = book.Description
	bookDetails.Author.FirstName = book.Author.FirstName
	bookDetails.Author.LastName = book.Author.LastName

	c.IndentedJSON(201, gin.H{"data": bookDetails, "message": "Book details retrieved successfully."})
}

func UpdateAuthorDetails(c *gin.Context) {
	var author models.Author
	err := CheckByID(c.Param("id"), &author)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	var authorSerializer serializers.AuthorSerializer
	if err := c.ShouldBindJSON(&authorSerializer); err != nil {
		c.IndentedJSON(400, gin.H{"message": "Invalid Data"})
		return
	}

	author.FirstName = authorSerializer.FirstName
	author.LastName = authorSerializer.LastName

	err = initializers.DB.Save(&author).Error
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"data": authorSerializer, "message": "Author details updated successfully."})

}

func DeleteAuthor(c *gin.Context) {
	var author models.Author
	err := CheckByID(c.Param("id"), &author)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	// Soft delete Author and related Books
	err = initializers.DB.Select("Books").Delete(&author).Error
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Author deleted successfully."})
}

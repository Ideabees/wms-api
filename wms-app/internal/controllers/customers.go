package controllers

import (
	"fmt"
	//"mime/multipart"
	"net/http"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/request"
	"wms-app/internal/services"
	"wms-app/internal/utils"

	"github.com/gin-gonic/gin"
)

func CreateCustomer(c *gin.Context) {
	userID := c.GetString("user_id")
	//email := c.GetString("email")

	var req request.CreateCustomerRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	cust_id := utils.CreateUUID()
	custModel := dbModels.Customer{
		CustomerId:   cust_id,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MobileNumber: req.MobileNumber,
		UserId:       userID,
	}

	// call service layer to insert the customer
	msg, err := services.CreateCustomer(&custModel)
	if err != nil {
		fmt.Println("DB insertion Failed", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "DB insertion Failed",
			"status":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Customer succefully created",
			"status":  msg,
		})
	}
}

func GetCustomers(c *gin.Context) {
	userID := c.GetString("user_id")
	firstName := c.GetString("first_name")
	lastName := c.GetString("last_name")

	// call service layer to insert the customer
	data, msg, err := services.GetCustomers(userID, firstName, lastName)
	if err != nil {
		fmt.Println("DB insertion Failed", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "DB Operation Failed",
			"status":  msg,
			"data":    data,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "",
			"status":  msg,
			"data":    data,
		})
	}
}


func DeleteCustomers(c *gin.Context) {
	//userID := c.GetString("user_id")
	//email := c.GetString("email")

	var req request.DeleteCustomer
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("request is", req)
	// call service layer to delete the customers
	msg, err := services.DeleteCustomers(&req)
	if err != nil {
		fmt.Println("Deletion Failed", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "Deletion Failed",
			"status":  msg,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "Customer succefully deleted",
			"status":  msg,
		})
	}
}

func CreateBulkCustomers(c *gin.Context) {
	file, _, err := c.Request.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Upload csv File"})
        return
    }
    defer file.Close()

    // Get file extension
    /*if err := processFile(file); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

	cust_id := utils.CreateUUID()
	custModel := dbModels.Customer{
		CustomerId:   cust_id,
		FirstName:    req.FirstName,
		LastName:     req.LastName,
		MobileNumber: req.MobileNumber,
		UserId:       userID,
	}

	// call service layer to insert the customer
	msg, err := services.CreateCustomer(&custModel)
	if err != nil {
		fmt.Println("DB insertion Failed", err)
		c.JSON(http.StatusOK, gin.H{
			"message": "DB insertion Failed",
			"status":  msg,
		})
	} else { 
		c.JSON(http.StatusOK, gin.H{
			"message": "Customer succefully created",
			"status":  msg,
		})
	}*/
	c.JSON(http.StatusOK, gin.H{
		"message": "Customer succefully created",
		"status":  "ok",
	})
}

/*func processFile(file multipart.File) error {
	 // Detect MIME type dynamically
	 mimeType := detectFileType(file)

	 // Reset file pointer to the beginning after detecting MIME type
	 file.Seek(0, 0)

    switch mimeType {
    case "csv":
        return processCSV(file)
    case "xlsx":
        return processExcel(file)
    default:
        return fmt.Errorf("unsupported file type")
    }
}

// Detect file type based on MIME type
func detectFileType(file http.File) string {
    // Read the first few bytes of the file to determine MIME type
    buf := make([]byte, 512)
    _, err := file.Read(buf)
    if err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    // Detect file MIME type
    mimeType := mimetype.Detect(buf)
    return mimeType.String()
}

// Process CSV file
func processCSV(file http.File) error {
    // Create a temporary file for the CSV content
    filePath := "./upload.csv"
    outFile, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer outFile.Close()

    // Write the file contents to outFile
    _, err = outFile.ReadFrom(file)
    if err != nil {
        return err
    }

    // Read CSV file
    var users []User
    f, err := os.Open(filePath)
    if err != nil {
        return err
    }
    defer f.Close()

    if err := gocsv.UnmarshalFile(f, &users); err != nil {
        return err
    }

    // Insert data into the database
    for _, user := range users {
        if err := db.Create(&user).Error; err != nil {
            return err
        }
    }

    return nil
}

// Process Excel file
func processExcel(file http.File) error {
    // Create a temporary file for the Excel content
    filePath := "./upload.xlsx"
    outFile, err := os.Create(filePath)
    if err != nil {
        return err
    }
    defer outFile.Close()

    // Write the file contents to outFile
    _, err = outFile.ReadFrom(file)
    if err != nil {
        return err
    }

    // Read the Excel file
    xlFile, err := xlsx.OpenFile(filePath)
    if err != nil {
        return err
    }

    // Iterate through the sheets and rows
    for _, sheet := range xlFile.Sheets {
        for _, row := range sheet.Rows {
            if len(row.Cells) < 2 {
                continue
            }
            name := row.Cells[0].String()
            email := row.Cells[1].String()

            // Insert data into the database
            user := User{Name: name, Email: email}
            if err := db.Create(&user).Error; err != nil {
                return err
            }
        }
    }

    return nil
}*/
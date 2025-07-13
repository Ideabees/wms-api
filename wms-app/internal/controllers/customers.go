package controllers

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"wms-app/config"
	"wms-app/internal/models/dbModels"
	"wms-app/internal/models/request"
	"wms-app/internal/services"
	"wms-app/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
	"github.com/tealeg/xlsx"
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
		EmailID:      req.EmailID,
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
	userID := c.GetString("user_id")
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Upload a CSV or Excel file"})
		return
	}
	defer file.Close()

	filename := header.Filename
	var customers []dbModels.Customer
	var parseErr error

	if strings.HasSuffix(strings.ToLower(filename), ".csv") {
		customers, parseErr = parseCSVCustomers(file, userID)
	} else if strings.HasSuffix(strings.ToLower(filename), ".xlsx") {
		customers, parseErr = parseExcelCustomers(file, userID)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported file type. Please upload .csv or .xlsx"})
		return
	}

	if parseErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": parseErr.Error()})
		return
	}

	var success, failed int
	for _, cust := range customers {
		msg, err := services.CreateCustomer(&cust)
		if err != nil {
			failed++
		} else if msg == "Success" {
			success++
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Bulk upload completed",
		"success_count": success,
		"failed_count":  failed,
	})
}

// Helper to parse CSV file into []Customer
func parseCSVCustomers(file multipart.File, userID string) ([]dbModels.Customer, error) {
	tmpFile, err := os.CreateTemp("", "bulk_upload_*.csv")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		return nil, err
	}
	tmpFile.Seek(0, 0)

	type csvCustomer struct {
		FirstName    string `csv:"FirstName"`
		LastName     string `csv:"LastName"`
		MobileNumber string `csv:"MobileNumber"`
		EmailID      string `csv:"EmailID"`
		City         string `csv:"City"`
		Pincode      string `csv:"Pincode"`
	}
	var csvCustomers []csvCustomer
	if err := gocsv.UnmarshalFile(tmpFile, &csvCustomers); err != nil {
		return nil, err
	}
	var customers []dbModels.Customer
	for _, c := range csvCustomers {
		customers = append(customers, dbModels.Customer{
			CustomerId:   utils.CreateUUID(),
			FirstName:    c.FirstName,
			LastName:     c.LastName,
			MobileNumber: c.MobileNumber,
			EmailID:      c.EmailID,
			City:         c.City,
			Pincode:      c.Pincode,
			UserId:       userID,
		})
	}
	return customers, nil
}

// Helper to parse Excel file into []Customer
func parseExcelCustomers(file multipart.File, userID string) ([]dbModels.Customer, error) {
	tmpFile, err := os.CreateTemp("", "bulk_upload_*.xlsx")
	if err != nil {
		return nil, err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()
	_, err = io.Copy(tmpFile, file)
	if err != nil {
		return nil, err
	}
	tmpFile.Seek(0, 0)

	xlFile, err := xlsx.OpenFile(tmpFile.Name())
	if err != nil {
		return nil, err
	}
	var customers []dbModels.Customer
	for _, sheet := range xlFile.Sheets {
		for idx, row := range sheet.Rows {
			// Assume first row is header
			if idx == 0 {
				continue
			}
			var firstName, lastName, mobile, email, city, pincode string
			if len(row.Cells) > 0 {
				firstName = row.Cells[0].String()
			}
			if len(row.Cells) > 1 {
				lastName = row.Cells[1].String()
			}
			if len(row.Cells) > 2 {
				mobile = row.Cells[2].String()
			}
			if len(row.Cells) > 3 {
				email = row.Cells[3].String()
			}
			if len(row.Cells) > 4 {
				city = row.Cells[4].String()
			}
			if len(row.Cells) > 5 {
				pincode = row.Cells[5].String()
			}
			if firstName == "" && lastName == "" && mobile == "" {
				continue
			}
			customers = append(customers, dbModels.Customer{
				CustomerId:   utils.CreateUUID(),
				FirstName:    firstName,
				LastName:     lastName,
				MobileNumber: mobile,
				EmailID:      email,
				City:         city,
				Pincode:      pincode,
				UserId:       userID,
			})
		}
	}
	return customers, nil
}

func processFile(file multipart.File) error {
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
func detectFileType(file multipart.File) string {
	// Read the first few bytes of the file to determine MIME type
	buf := make([]byte, 512)
	_, err := file.Read(buf)
	if err != nil {
		// You may want to handle this error differently
		return ""
	}

	// Detect file MIME type
	// Replace mimetype.Detect with your own logic or use a package that works with []byte
	// For demonstration, let's check for CSV or XLSX by magic numbers
	if len(buf) >= 4 && buf[0] == 0x50 && buf[1] == 0x4B && buf[2] == 0x03 && buf[3] == 0x04 {
		return "xlsx" // XLSX files are ZIP archives
	}
	if strings.Contains(string(buf), ",") {
		return "csv"
	}
	return "unknown"
}

// Process CSV file
func processCSV(file multipart.File) error {
	// Create a temporary file for the CSV content
	filePath := "./upload.csv"
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write the file contents to outFile
	_, err = io.Copy(outFile, file)
	if err != nil {
		return err
	}

	// Read CSV file
	var customers []dbModels.Customer
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err := gocsv.UnmarshalFile(f, &customers); err != nil {
		return err
	}

	// Insert data into the database
	for _, customer := range customers {
		if err := config.DB.Create(&customer).Error; err != nil {
			return err
		}
	}

	return nil
}

// Process Excel file
func processExcel(file multipart.File) error {
	// Create a temporary file for the Excel content
	filePath := "./upload.xlsx"
	outFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer outFile.Close()

	// Write the file contents to outFile
	_, err = io.Copy(outFile, file)
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
			FirstName := row.Cells[0].String()
			LastName := row.Cells[1].String()
			Mobile := row.Cells[2].String()
			Email := row.Cells[3].String()
			City := row.Cells[4].String()
			Pincode := row.Cells[5].String()

			// Insert data into the database
			customer := dbModels.Customer{
				FirstName:    FirstName,
				LastName:     LastName,
				MobileNumber: Mobile,
				EmailID:      Email,
				City:         City,
				Pincode:      Pincode,
			}
			if err := config.DB.Create(&customer).Error; err != nil {
				return err
			}
		}
	}
	return nil
}

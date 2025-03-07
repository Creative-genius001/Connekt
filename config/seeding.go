package config

import (
	"math/rand"
	"time"

	"github.com/Creative-genius001/Connekt/cmd/models"
	"github.com/google/uuid"
)

var titles = []string{"Software Engineer", "Data Scientist", "Product Manager", "DevOps Engineer", "UX Designer"}
var summaries = []string{"Developing scalable applications", "Analyzing large datasets", "Managing product lifecycle", "Ensuring system reliability", "Designing user interfaces"}
var locations = []string{"New York, NY", "San Francisco, CA", "Austin, TX", "Boston, MA", "Chicago, IL"}
var industries = []string{"Technology", "Finance", "Healthcare", "Education", "Retail"}
var salary = models.Salary{
	MinValue: 70000,
	MaxValue: 50000,
	Currency: "USD",
}

func Seeding() ([]models.Job, error) {

	rand.New(rand.NewSource(time.Now().UnixNano()))
	var jobs []models.Job

	for i := 0; i < 100; i++ {
		job := models.Job{
			Id:          uuid.New().String(),
			CompanyId:   "20a88f2c-8b21-45bc-8a01-e252b002a999",
			Title:       titles[rand.Intn(len(titles))],
			Salary:      &salary,
			Description: summaries[rand.Intn(len(summaries))],
			Remote:      []bool{true, false}[rand.Intn(2)],
			IsActive:    rand.Intn(2) == 1,
			Industry:    industries[rand.Intn(len(industries))],
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
}

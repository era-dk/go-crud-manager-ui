package manager

import (
	"errors"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/Oudwins/zog"
)

type UserGender int
const (
    UserGenderUnknown UserGender = iota
    UserGenderMale
    UserGenderFemale
)
var UserGenderLabels = map[UserGender]string{
	UserGenderUnknown: "Unknown",
	UserGenderMale: "Male",
	UserGenderFemale: "Female",
}

type UserTag int
const (
    UserTagPrimary UserTag = iota + 1
    UserTagSecondary
    UserTagExtend
)
var UserTagLabels = map[UserTag]string{
	UserTagPrimary: "Primary",
	UserTagSecondary: "Secondary",
	UserTagExtend: "Extend",
}

type User struct {
	ID int
	Name string
	Email string
	Gender UserGender
	Tags []UserTag
}

func (u User) TagsLabel() string {
	var tags []string
	for _, tag := range u.Tags {
		tags = append(tags, UserTagLabels[tag])
	}
	return strings.Join(tags, ", ")
}

var testData = []User{
	{ID: 1, Name: "Oliver", Email: "oliver@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagPrimary, UserTagSecondary}},
	{ID: 2, Name: "Jack", Email: "jack@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagPrimary, UserTagSecondary}},
	{ID: 3, Name: "Jacob", Email: "jacob@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagSecondary}},
	{ID: 4, Name: "Thomas", Email: "thomas@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagSecondary}},
	{ID: 5, Name: "Emily", Email: "emily@email.com", Gender: UserGenderFemale, Tags: []UserTag{UserTagSecondary, UserTagExtend}},
	{ID: 6, Name: "Joanne", Email: "joanne@email.com", Gender: UserGenderFemale, Tags: []UserTag{UserTagSecondary, UserTagExtend}},
	{ID: 7, Name: "James", Email: "james@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagSecondary}},
	{ID: 8, Name: "Elizabeth", Email: "elizabeth@email.com", Gender: UserGenderFemale, Tags: []UserTag{UserTagPrimary, UserTagSecondary, UserTagExtend}},
	{ID: 9, Name: "Jessica", Email: "jessica@email.com", Gender: UserGenderFemale, Tags: []UserTag{UserTagPrimary, UserTagExtend}},
	{ID: 10, Name: "Robert", Email: "robert@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagSecondary}},
	{ID: 11, Name: "Richard", Email: "richard@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagSecondary, UserTagExtend}},
	{ID: 12, Name: "Sophie", Email: "sophie@email.com", Gender: UserGenderFemale, Tags: []UserTag{UserTagSecondary}},
	{ID: 13, Name: "Megan", Email: "megan@email.com", Gender: UserGenderFemale, Tags: []UserTag{UserTagPrimary, UserTagSecondary}},
	{ID: 14, Name: "Susan", Email: "susan@email.com", Gender: UserGenderFemale, Tags: []UserTag{UserTagSecondary, UserTagExtend}},
	{ID: 15, Name: "Jennifer", Email: "jennifer@email.com", Gender: UserGenderFemale, Tags: []UserTag{UserTagPrimary, UserTagSecondary}},
	{ID: 16, Name: "Michael", Email: "michael@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagSecondary}},
	{ID: 17, Name: "Joseph", Email: "joseph@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagExtend}},
	{ID: 18, Name: "Charlie", Email: "charlie@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagPrimary, UserTagSecondary}},
	{ID: 19, Name: "George", Email: "george@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagSecondary}},
	{ID: 20, Name: "Oscar", Email: "oscar@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagPrimary, UserTagSecondary}},
	{ID: 21, Name: "Kyle", Email: "kyle@email.com", Gender: UserGenderMale, Tags: []UserTag{UserTagExtend}},
}

func getTestDataList(limit int, offset int) []User {
	var results []User
	for idx, item := range testData {
		if idx < offset {
			continue
		}
		results = append(results, item)
		if len(results) == limit {
			break
		}
	}
	return results
}

func getTestDataRecord(id int) (int, *User) {
	for idx, item := range testData {
		if item.ID == id {
			return idx, &item
		}
	}
	return 0, nil
}

func saveTestDataRecord(user User) {
	if user.ID != 0 {
		idx, _ := getTestDataRecord(user.ID)
		testData[idx].Name = user.Name
		testData[idx].Email = user.Email
		testData[idx].Gender = user.Gender
		return
	}
	
	maxId := 0
	for _, item := range testData {
		if item.ID > maxId {
			maxId = item.ID
		}
	}
	maxId++

	user.ID = maxId
	testData = append(testData, user)
}

func deleteTestDataRecord(id int) {
	for idx, item := range testData {
		if item.ID == id {
			testData = slices.Delete(testData, idx, idx + 1)
			break
		}
	}
}

func TestExample(t *testing.T) {
	Grid.
		AddColumn("ID", 10).
		AddColumn("Name", 20).
		AddColumn("Email", 40).
		AddColumn("Gender", 20).
		AddColumn("Tags", 30)
	Grid.FetchFn = func (limit int, offset int) (RecordCollection, int, error) {
		results := getTestDataList(limit, offset)
		
		records := NewRecordCollection()
		for _, item := range results {
			records.Add(item.ID).
				Set("ID", strconv.Itoa(item.ID)).
				Set("Name", item.Name).
				Set("Email", item.Email).
				Set("Gender", UserGenderLabels[item.Gender]).
				Set("Tags", item.TagsLabel())
		}
		return records, len(testData), nil
	}
	Edit.Handle = func (id int) (Record, error) {
		_, user := getTestDataRecord(id)
		if user != nil {
			record := NewRecord(id)
			record.
				Set("name", user.Name).
				Set("email", user.Email).
				Set("gender", user.Gender)
			return record, nil
		}
		return NewRecord(0), errors.New("record not found")
	}
	Delete.Handle = func (id int) error {
		deleteTestDataRecord(id)
		return nil
	}
	View.Handle = func (id int) (Record, error) {
		_, user := getTestDataRecord(id)
		if user != nil {
			record := NewRecord(id)
			record.
				Set("Name", user.Name).
				Set("Email", user.Email).
				Set("Gender", UserGenderLabels[user.Gender]).
				Set("Custom 1", "my custom value 1").
				Set("Custom 2", "my custom value 2")
			return record, nil
		}
		return NewRecord(0), errors.New("record not found")
	}
	Form.Fields = []FormFieldInterface{
		&FormFieldInput{
			Name: "name",
			Label: "Name",
			ValidateFn: func(value string) error {
				errsMap := zog.String().Trim().Required().Validate(&value)
				for _, err := range errsMap {
					return fmt.Errorf(`field %s`, err.Message)
				}
				return nil
			},
		},
		&FormFieldInput{
			Name: "email",
			Label: "Email",
			ValidateFn: func(value string) error {
				errsMap := zog.String().Trim().Required().Email().Validate(&value)
				for _, err := range errsMap {
					return fmt.Errorf(`field %s`, err.Message)
				}
				return nil
			},
		},
		&FormFieldSelect{
			Name: "gender",
			Label: "Gender",
			Items: []FormFieldSelectItem{
				{ID: int(UserGenderUnknown), Label: UserGenderLabels[UserGenderUnknown]},
				{ID: int(UserGenderMale), Label: UserGenderLabels[UserGenderMale]},
				{ID: int(UserGenderFemale), Label: UserGenderLabels[UserGenderFemale]},
			},
			Required: true,
		},
		&FormSubmit{
			Label: "Submit",
			Handle: func (id int) error {
				saveTestDataRecord(User{
					ID: Form.RecordID,
					Name: Form.GetField("name").GetValue().GetString(),
					Email: Form.GetField("email").GetValue().GetString(),
					Gender: UserGender(Form.GetField("gender").GetValue().GetInt()),
				})
				return nil
			},
		},
	}

	if err := Run(); err != nil {
		log.Fatal(err)
	}
}
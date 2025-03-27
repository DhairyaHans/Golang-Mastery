package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

// Model for course - file
type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// Fake DB
var courses []Course

// MiddleWare or Helper Functions - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{
		CourseId:    "2",
		CourseName:  "React Js",
		CoursePrice: 69,
		Author: &Author{
			Fullname: "Dhairya Hans",
			Website:  "xxx.com",
		},
	})
	courses = append(courses, Course{
		CourseId:    "4",
		CourseName:  "MERN Stack",
		CoursePrice: 129,
		Author: &Author{
			Fullname: "Dhairya Hans",
			Website:  "xnxx.com",
		},
	})

	// Routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// Listen to a port
	log.Fatal(http.ListenAndServe(":4000", r))
}

// Controllers - file

// Serve Home Route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to HTTP API Routes</h1>"))
}

// Get All Courses Route
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get One Course Route
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course by Course Id")
	w.Header().Set("Content-Type", "application/json")

	// Grab id from request
	params := mux.Vars(r) // params = key:value store

	fmt.Println(params)
	fmt.Printf("Params Type %T\n", params)

	// loop through courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with the given id, " + params["id"])
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if: body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what about - {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No Data inside the JSON")
		return
	}

	// If coursename already exists, don't allow to create new one
	for _, crs := range courses {
		if crs.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course with same course name already exists")
			return
		}
	}

	// Generate a unique course id and convert it to string
	source := rand.NewSource(time.Now().UnixNano())
	rand.New(source)
	course.CourseId = strconv.Itoa(rand.Intn(100))

	// append course into courses
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// First - Grab courseId from the request
	params := mux.Vars(r)

	// loop through Courses, remove the course, and add updated data again
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TODO: Send a response when Id not found
	json.NewEncoder(w).Encode("No Course found with the Id")
	return
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete One Course")
	w.Header().Set("Content-Type", "application/json")

	// First - Get the params
	params := mux.Vars(r)

	// Loop through couses, and remove the course with matching courseId
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode("Course Deleted")
			return
		}
	}

	json.NewEncoder(w).Encode("No Course found with the specified course Id")
	return
}

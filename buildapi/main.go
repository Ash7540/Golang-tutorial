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

// fake DB
var courses []Course

// middleware, helper - file
func (c *Course) IsEmpty() bool {
	// return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

// controllers - file

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to Ashu Wbe</h1>"))
}

// Get all Courses
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// Get Single Course
func getSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get single course")
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)

	// loop through the courses, find matching id and return the response
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No Course found with given id")
	return
}

// create course
func createSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create single course")
	w.Header().Set("Content-Type", "application/json")

	// what if : body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	// what if : body is {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("No data inside JSON")
		return
	}

	//TODO: check only if title is duplicate
	// loop, title matches with course.coursename, JSON

	for _, c := range courses {
		if c.CourseName == course.CourseName {
			json.NewEncoder(w).Encode("Course already exists")
			return
		}
	}

	// generate unique id, and converting it into string
	// append course into courses

	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return

}

// update the course
func updateSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update single course")
	w.Header().Set("Content-Type", "application/json")

	// first - grab id from req
	params := mux.Vars(r)

	// loop, id, remove, add with my ID
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
	//TODO: send a response when id is not found
	json.NewEncoder(w).Encode("No Course found with given id")
	return

}

// Delete the Course
func deleteSingleCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete single course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	//loop, id, remove (index, index+1)
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			// TODO: send a confirm or deny response
			json.NewEncoder(w).Encode("Course deleted successfully")
			return
		}
	}
}

// Delete All Course
func deleteAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete all courses")
	w.Header().Set("Content-Type", "application/json")

	courses = []Course{}
	json.NewEncoder(w).Encode("All courses deleted successfully")
	return
}

func main() {
	fmt.Println("Building API in Golang")
	r := mux.NewRouter()

	// seeding
	courses = append(courses, Course{CourseId: "3", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	// routing
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getSingleCourse).Methods("GET")
	r.HandleFunc("/course", createSingleCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateSingleCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteSingleCourse).Methods("DELETE")
	r.HandleFunc("/courses", deleteAllCourses).Methods("DELETE")

	// listen to a port
	log.Fatal(http.ListenAndServe(":8000", r))

}

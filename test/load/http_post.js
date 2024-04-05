import http from 'k6/http';
import { check, sleep } from 'k6';
import { SharedArray } from 'k6/data';

// Define your student emails in a shared array to avoid data being duplicated across VUs
const studentEmails = new SharedArray('student emails', function () {
  return ['student1@example.com', 'student2@example.com','student3@example.com','student4@example.com','student5@example.com','student6@example.com','student7@example.com','student8@example.com','student9@example.com','student10@example.com'];
});

// List of course IDs
const courseIds = ['CS101', 'CS102', 'CS103', 'CS104'];

export let options = {
    vus: 10,
    duration: '1m',
};

function randomElement(array) {
  // Helper function to select a random element from an array
  return array[Math.floor(Math.random() * array.length)];
}

export default function () {
    // Select a random email and course
    const randomEmail = randomElement(studentEmails);
    const randomCourseId = randomElement(courseIds);
    
    // Define the URL
    const url = `http://localhost:8000/students/${encodeURIComponent(randomEmail)}/courses`;
    
    // Define the request body
    const payload = JSON.stringify({
        courseId: randomCourseId,
    });

    // Define the request headers
    const params = {
        headers: {
            'Content-Type': 'application/json',
        },
    };

    // Send a POST request
    let response = http.post(url, payload, params);

    // Use 'check' to add some checks for pass/fail criteria
    check(response, {
        'is status 200': (r) => r.status === 200,
        'response body contains courseId': (r) => JSON.parse(r.body).courseId == randomCourseId,
    });

    // Print out the response body for debugging purposes
    console.log(response.body);

    // Sleep for a random period between iterations to simulate real user behavior
    sleep(Math.random() * 3);
}
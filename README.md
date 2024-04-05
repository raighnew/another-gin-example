# Courses Service

## How to run

### Build Docker Image first

Run docker build command:

```bash
docker build -t 1.1.1.1:5000/demo-api:v1 -f deploy/build/Dockerfile --build-arg APP_CONF=config/prod.yml --build-arg  APP_RELATIVE_PATH=./cmd/server/...  .
```

### Using Docker Compose to run the service

```bash
docker compose -f deploy/docker-compose/docker-compose.yml up -d
```

After few mins, the demo-api will up
 
## API Specification

### List Courses

URL: GET /courses

#### Response

```json
{
  "code": 0,
  "message": "ok",
  "data": [
    {
      "courseId": "CS101",
      "name": "Intro to Computer Science",
      "lessons": 24
    },
    {
      "courseId": "CS102",
      "name": "Intro to Computer Science 2",
      "lessons": 24
    },
    {
      "courseId": "MATH255",
      "name": "Calculus I",
      "lessons": 10
    },
    {
      "courseId": "PHYS150",
      "name": "General Physics",
      "lessons": 15
    }
  ]
}
```

### Sign Up Courses

URL: POST /students/:studentEmail/courses

#### Parameters

| name         | location   | descriptions          |
| ------------ | ---------- | --------------------- |
| studentEmail | Parameters | The id of the student |
| courseId     | Body       | The id of the course  |

#### Response

Success:

```json
{
  "code": 0,
  "message": "ok",
  "data": {
    "studentId": "mail3@test.com",
    "courseId": "CS102"
  }
}
```

Error: 409, 500

```
{
  "code": 409,
  "message": "Status Conflict",
  "data": "Student has already signed up for this course."
}
```

### Delete Signed Courses

URL: DELETE /students/:studentEmail/courses/:courseId

#### Parameters

| name         | location   | descriptions          |
| ------------ | ---------- | --------------------- |
| studentEmail | Parameters | The id of the student |
| courseId     | Parameter  | The id of the course  |

#### Response

Success:

```json
{
  "code": 0,
  "message": "ok"
}
```

Error: 404, 500

```
{
  "code": 404,
  "message": "Resource Not Found",
  "data": "Error deleting signed up courses"
}
```

### Get My Courses

URL: GET /students/:studentEmail/courses

#### Parameters

| name         | location   | descriptions          |
| ------------ | ---------- | --------------------- |
| studentEmail | Parameters | The id of the student |

#### Response

Success:

```json
{
  "code": 0,
  "message": "ok",
  "data": [
    {
      "courseId": "CS102",
      "name": "Intro to Computer Science 2",
      "lessons": 24
    }
  ]
}
```

Error: 500

```
{
  "code": 500,
  "message": "Internal Server Error",
  "data": "Error Internal Server Error"
}
```

### Get My Course Classmates

URL: GET /students/:studentEmail/courses/:courseId/classmates

#### Parameters

| name         | location   | descriptions          |
| ------------ | ---------- | --------------------- |
| studentEmail | Parameters | The id of the student |
| courseId     | Parameter  | The id of the course  |

#### Response

Success:

```json
{
  "code": 0,
  "message": "ok",
  "data": [
    {
      "studentId": "mail1@test.com",
      "courseId": "CS102"
    },
    {
      "studentId": "mail2@test.com",
      "courseId": "CS102"
    },
    {
      "studentId": "mail3@test.com",
      "courseId": "CS102"
    }
  ]
}
```

Error: 500

```
{
  "code": 500,
  "message": "Internal Server Error",
  "data": "Error deleting signed up courses"
}
```

# Unit Tests

# Load Tests

# TODO Feature

## Pagination Response

If courses number goes up to 50, can pagination the get courses and list classmates API.

## Use Redis

Can load courses infos in Redis when server setup.

## Use Redis + Message Queue

If the course has a seating limit, you can use redis and message queues to handle concurrent requests.

## Swagger Inline Documentation
Can use inline codes auto generate API documentation, so when codes review it only review codes itself, but also can review the documentation.

Example:

```txt
// courses godoc
// @Summary List Course
// @Schemes
// @Description List all courses
// @Tags List Courses
// @Accept json
// @Produce json
// @Param request body v1.RegisterRequest true "params"
// @Success 200 {object} v1.Response
// @Router /courses [get]
```
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

I used MockGen for mocking and use the included test framework for testing, all mocked functions are located under `/test/mocks/`.

## Handler Layer Tests

Get function Example codes:
```go
func TestCourseHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	courseService := mock_service.NewMockCourseService(ctrl)
	courseService.EXPECT().ListCourses(gomock.Any()).Return(nil, nil)

	coursesHandler := handler.NewCourseHandler(hdl, courseService)

	router.GET("/courses", coursesHandler.ListCourses)

	req, _ := http.NewRequest("GET", "/courses", nil)

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}
```

POST function Example codes:
```go
func TestCourseHandler_SignUp(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	params := v1.SignUpRequest{
		CourseID: "CS101",
	}

	courseService := mock_service.NewMockCourseService(ctrl)
	courseService.EXPECT().IfCourseExists(gomock.Any(), "CS101").Return(true, nil)
	courseService.EXPECT().SignUpCourse(gomock.Any(), gomock.Any(), "CS101").Return(&model.Enrollment{
		StudentID: "test@mail.com",
		CourseID:  "CS101",
	}, nil)

	coursesHandler := handler.NewCourseHandler(hdl, courseService)

	router.POST("/students/:studentEmail/courses", coursesHandler.SignUpCourse)

	paramsJson, _ := json.Marshal(params)

	req, _ := http.NewRequest("POST", "/students/test@mail.com/courses", bytes.NewBuffer(paramsJson))

	resp := httptest.NewRecorder()

	router.ServeHTTP(resp, req)
	assert.Equal(t, resp.Code, http.StatusOK)
}

```

## Service Layer Tests
```go
func TestUserService_GetSignedUpCourses(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCourseRepo := mock_repository.NewMockCourseRepository(ctrl)
	mockEnrollmentRepo := mock_repository.NewMockEnrollmentRepository(ctrl)
	srv := service.NewService(logger)

	courseService := service.NewCourseService(srv, mockCourseRepo, mockEnrollmentRepo)

	ctx := context.Background()

	mockCourseRepo.EXPECT().ListSignedUpCourses(ctx, "test@mail.com").Return(nil, nil)

	_, err := courseService.GetSignedUpCourses(ctx, "test@mail.com")

	assert.NoError(t, err)
}
```

## Test Coverage

In terms of time, I've just finished unit testing for each layer, coverage is not 100%, but just need more time to write more test cases. Here are the basic test results:


```
go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

	course-sign-up/api/v1		coverage: 0.0% of statements
	course-sign-up/cmd/server		coverage: 0.0% of statements
	course-sign-up/internal/model		coverage: 0.0% of statements
	course-sign-up/internal/handler		coverage: 0.0% of statements
	course-sign-up/internal/repository		coverage: 0.0% of statements
	course-sign-up/internal/server		coverage: 0.0% of statements
	course-sign-up/pkg/http		coverage: 0.0% of statements
	course-sign-up/internal/service		coverage: 0.0% of statements
	course-sign-up/pkg/config		coverage: 0.0% of statements
	course-sign-up/test/mocks/repository		coverage: 0.0% of statements
	course-sign-up/pkg/helper/resp		coverage: 0.0% of statements
	course-sign-up/pkg/log		coverage: 0.0% of statements
	course-sign-up/test/mocks/service		coverage: 0.0% of statements
ok  	course-sign-up/test/server/handler	0.164s	coverage: 30.4% of statements
ok  	course-sign-up/test/server/service	0.269s	coverage: 22.8% of statements
```

# Load Tests

Before pushing service to the production, we need do load tests and stress tests on staging env. I choose using k6/http to test the service. the test file locates on `test/load`

Commands:
```bash
k6 run test/load/http_post.js
```

Results:

```Bash
     scenarios: (100.00%) 1 scenario, 10 max VUs, 1m30s max duration (incl. graceful stop):
              * default: 10 looping VUs for 1m0s (gracefulStop: 30s)

INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student5@example.com","courseId":"CS102"}}  source=console
INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student10@example.com","courseId":"CS101"}}  source=console
INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student2@example.com","courseId":"CS103"}}  source=console
INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student5@example.com","courseId":"CS104"}}  source=console
INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student9@example.com","courseId":"CS104"}}  source=console
INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student3@example.com","courseId":"CS103"}}  source=console
INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student7@example.com","courseId":"CS103"}}  source=console
INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student4@example.com","courseId":"CS103"}}  source=console
INFO[0000] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
INFO[0000] {"code":0,"message":"ok","data":{"studentId":"student8@example.com","courseId":"CS101"}}  source=console
INFO[0001] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
INFO[0001] {"code":0,"message":"ok","data":{"studentId":"student8@example.com","courseId":"CS102"}}  source=console
INFO[0001] {"code":0,"message":"ok","data":{"studentId":"student1@example.com","courseId":"CS104"}}  source=console
INFO[0001] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
INFO[0002] {"code":0,"message":"ok","data":{"studentId":"student7@example.com","courseId":"CS104"}}  source=console
INFO[0002] {"code":0,"message":"ok","data":{"studentId":"student9@example.com","courseId":"CS102"}}  source=console
INFO[0002] {"code":0,"message":"ok","data":{"studentId":"student10@example.com","courseId":"CS102"}}  source=console
INFO[0002] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
INFO[0002] {"code":0,"message":"ok","data":{"studentId":"student9@example.com","courseId":"CS103"}}  source=console
INFO[0002] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
INFO[0002] {"code":0,"message":"ok","data":{"studentId":"student4@example.com","courseId":"CS102"}}  source=console
INFO[0002] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
INFO[0002] {"code":0,"message":"ok","data":{"studentId":"student2@example.com","courseId":"CS104"}}  source=console
INFO[0003] {"code":0,"message":"ok","data":{"studentId":"student9@example.com","courseId":"CS101"}}  source=console
INFO[0003] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
INFO[0003] {"code":0,"message":"ok","data":{"studentId":"student4@example.com","courseId":"CS104"}}  source=console
INFO[0003] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
INFO[0003] {"code":409,"message":"Status Conflict","data":"Student has already signed up for this course."}  source=console
 ✗ is status 200
      ↳  12% — ✓ 40 / ✗ 278
     ✗ response body contains courseId
      ↳  0% — ✓ 0 / ✗ 318

     checks.........................: 6.28%  ✓ 40       ✗ 596
     data_received..................: 71 kB  1.5 kB/s
     data_sent......................: 61 kB  1.3 kB/s
     http_req_blocked...............: avg=50.13µs min=1µs     med=7µs    max=1.58ms  p(90)=15µs     p(95)=17.14µs
     http_req_connecting............: avg=17.22µs min=0s      med=0s     max=627µs   p(90)=0s       p(95)=0s
     http_req_duration..............: avg=8.3ms   min=3.91ms  med=7.12ms max=36.34ms p(90)=10.35ms  p(95)=11.74ms
       { expected_response:true }...: avg=14.07ms min=5.34ms  med=8.93ms max=36.34ms p(90)=34.6ms   p(95)=34.67ms
     http_req_failed................: 87.42% ✓ 278      ✗ 40
     http_req_receiving.............: avg=89.95µs min=9µs     med=83µs   max=335µs   p(90)=152.29µs p(95)=180.44µs
     http_req_sending...............: avg=52.46µs min=12µs    med=46µs   max=332µs   p(90)=79.3µs   p(95)=90.44µs
     http_req_tls_handshaking.......: avg=0s      min=0s      med=0s     max=0s      p(90)=0s       p(95)=0s
     http_req_waiting...............: avg=8.16ms  min=3.86ms  med=7.03ms max=36.2ms  p(90)=10.16ms  p(95)=11.54ms
     http_reqs......................: 318    6.664265/s
     iteration_duration.............: avg=1.52s   min=35.65ms med=1.48s  max=3s      p(90)=2.65s    p(95)=2.87s
     iterations.....................: 308    6.454697/s
     vus............................: 10     min=10     max=10
     vus_max........................: 10     min=10     max=10


running (0m47.7s), 00/10 VUs, 308 complete and 10 interrupted iterations
default ✗ [=============================>--------] 10 VUs  0m47.7s/1m0s
```

# TODO Feature

## Security issue
We need to authenticate the user so that if they change the :studentEmail in the url, they can view other courses. Therefore, it's a good idea to introduce Auth middleware to check if someone has permission to view the requested resource.

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
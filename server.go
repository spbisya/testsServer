package main

import (
  "log"
  "strconv"
  "strings"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

)

var dbmap = initDb()

func main() {
	r := gin.Default()
  r.LoadHTMLGlob("*.html")
	r.Use(Cors())

	v1 := r.Group("api/tests")
	{
		v1.GET("", GetTests)
		v1.GET("/:id", GetTestById)
    v1.POST("", AddTest)
    // v1.DELETE("/:id", DeleteTest)
	}

  v2 := r.Group("api/filters")
	{
		v2.GET("", GetFilters)
    v2.POST("", AddFilter)
    // v2.DELETE("/:id", DeleteFilter)
	}

	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", gin.H{
				"title": "Main website",
		})
	})
	r.Run()
}


func GetTests(c *gin.Context) {
  var tests []Test
	_, err := dbmap.Select(&tests, "SELECT * FROM tests")

  for i:=0;i<len(tests);i=i+1{

  var questions []QuestionModel
  _, err=  dbmap.Select(&questions, "SELECT * FROM questions WHERE uid=?", tests[i].Uid)
  for k:=0;k<len(questions);k=k+1{
    var answers []AnswerModel
  _, err=  dbmap.Select(&answers, "SELECT * FROM answers WHERE uid=?", tests[i].Uid)
    questions[k].Answers = answers
  }
  tests[i].Questions=questions
  var results []ResultInfo
_, err=    dbmap.Select(&results, "SELECT * FROM results WHERE uid=?", tests[i].Uid)
tests[i].ResultInfos = results
  }

	if err == nil {
		c.JSON(200, tests)
	} else {
    checkErr(err, "error")
		c.JSON(404, gin.H{"error": "no test(s) into the table"})
	}
	// curl -i http://localhost:8080/api/tests
}


func GetTestById(c *gin.Context) {
  id := c.Params.ByName("id")
  var test Test
	err := dbmap.SelectOne(&test, "SELECT * FROM tests WHERE uid=? LIMIT 1", id)
  var questions []QuestionModel
  _, err=  dbmap.Select(&questions, "SELECT * FROM questions WHERE uid=?", id)
  for k:=0;k<len(questions);k=k+1{
    var answers []AnswerModel
  _, err=  dbmap.Select(&answers, "SELECT * FROM answers WHERE uid=?", id)
    questions[k].Answers = answers
  }
  test.Questions=questions
  var results []ResultInfo
_, err=    dbmap.Select(&results, "SELECT * FROM results WHERE uid=?", id)
test.ResultInfos = results


	if err == nil {
		c.JSON(200, test)
	} else {
    checkErr(err, "error")
		c.JSON(404, gin.H{"error": "no test(s) into the table"})
	}
	// curl -i http://localhost:8080/api/tests
}

func GetFilters(c *gin.Context) {
	var filters []Filter
	_, err := dbmap.Select(&filters, "SELECT * FROM filters")
  for i:=0;i<len(filters);i=i+1{
    result := strings.Split(filters[i].TestsString, ",")
    log.Println(filters)
    filters[i].Tests, err = sliceAtoi(result)
  }
	if err == nil {
		c.JSON(200, filters)
	} else {
    checkErr(err, "showing failed")
		c.JSON(404, gin.H{"error": "no filters(s) into the table"})
	}
	// curl -i http://volhgroup.tk/api/v1/waifus
}

func sliceAtoi(sa []string) ([]int64, error) {
    si := make([]int64, 0, len(sa))
    for _, a := range sa {
        i, err := strconv.ParseInt(a, 10, 64)
        if err != nil {
            return si, err
        }
        si = append(si, i)
    }
    return si, nil
}

/*
{ "name": "Who R u from Hex?", "Description": "Who are you, waifu bro?", "image": "http://lorempixel.com/400/200/", "questions": [{"questionText": "Question 2", "answers":[{"answer":"Answer 2", "points": 10}]}], "results": [{"start":15, "end":20, "description": "Cat likes milk", "image": "http://lorempixel.com/400/200/"}] }

*/

func AddFilter(c *gin.Context) {
	var filter Filter
	c.Bind(&filter)
	log.Println(filter)
  testsString := strconv.FormatInt(filter.Tests[0], 10)
  for i:=1;i<len(filter.Tests);i=i+1{
    testsString = testsString+","+strconv.FormatInt(filter.Tests[i], 10)
  }
  if filter.Title != "" && filter.Description != "" && testsString != ""{
    if insert, err := dbmap.Exec(`INSERT INTO filters (title, description, tests) VALUES (?, ?, ?)`,
    filter.Title , filter.Description, testsString); insert != nil {
  if(err==nil){
  content := &Filter{
    Title: filter.Title,
    Description:  filter.Description,
    Tests: filter.Tests,
    TestsString: filter.TestsString,
  }
  c.JSON(201, content)
  } else {
  checkErr(err, "Insert failed")
  }
  }
  } else {
  c.JSON(400, gin.H{"error": "Fields are empty"})
  }
  }

func AddTest(c *gin.Context) {
	var test Test
	c.Bind(&test)
	log.Println(test)
	if test.Name != "" && test.Description != "" && test.Image != "" && len(test.Questions)>0 && len(test.ResultInfos)>0{
		if insert, _ := dbmap.Exec(`INSERT INTO tests (name, description, image) VALUES (?, ?, ?)`,
    test.Name, test.Description, test.Image); insert != nil {
      test_id, err := insert.LastInsertId()
			if err == nil {
        for i:=0; i<len(test.Questions);i=i+1{
          log.Println(test.Questions[i])
           dbmap.Exec(`INSERT INTO questions (uid, question) VALUES (?, ?)`, test_id, test.Questions[i].Question);
           for k:=0; k<len(test.Questions[i].Answers); k=k+1{
             dbmap.Exec(`INSERT INTO answers (uid, answer, points) VALUES (?, ?, ?)`,
              test_id, test.Questions[i].Answers[k].Answer, test.Questions[i].Answers[k].Points);
           }
        }
        for i:=0; i<len(test.ResultInfos);i=i+1{
           dbmap.Exec(`INSERT INTO results (uid, start, end, description, image) VALUES (?, ?,?,?,?)`,
            test_id, test.ResultInfos[i].Start, test.ResultInfos[i].End, test.ResultInfos[i].Description, test.ResultInfos[i].Image);
        }
				content := &Test{
          Uid: test_id,
          Name: test.Name,
          Description:  test.Description,
          Image: test.Image,
          Questions: test.Questions,
          ResultInfos: test.ResultInfos,
				}
				c.JSON(201, content)
			} else {
				checkErr(err, "Insert failed")
			}
		}
	} else {
		c.JSON(400, gin.H{"error": "Fields are empty"})
	}
}
	//  curl -i -X POST -H "Content-Type: application/json" -d "{ \"name\": \"Who R u from Winx?\", \"Description\": \"Who are you, bro? Ur life is shit\", \"image\": \"http://lorempixel.com/400/200/\" }" http://localhost:8080/api/tests}

# triviaGo api documentation

### introduction
##### the triviaGo api
*TODO:overview*

##### What API commands are used by this API?

|command|purpose|
|---|---|
|get|fetch one or more objects from a given endpoint|
|put|not yet in service|
|post|fetch one or more objects from a given endpoint with specific attributes|
|delete|not yet in service|
 
`root url`: https://triviagoserver.herokuapp.com/api
***

### endpoints

---
#### questions
* **url**: https://triviagoserver.herokuapp.com/api/questions 
* **overview**: this endpoint is used to retrieve Question objects
* **request type**: post

**request example**

|attribute|type|description|
|----|----|----|
|category|integer|should be derived from the 'ID' field of each Category|
|difficulty|string|indicating question difficulty: 'any', 'easy', 'medium', 'hard'|
|question_quantity|integer|representing number of questions to return|
|question_type|string|indicating question type: 'any', 'multiple', 'boolean'|

```javascript
axios.post("https://triviagoserver.herokuapp.com/api/question",
  {
      "category": 0,
      "difficulty": "easy",
      "question_quantity": 2,
      "question_type": "any"
  },
  { 
    headers: {
      "Content-Type": "application/x-www-form-urlencoded"
    }
  });
```

**response example**

|attribute|type|description|
|----|----|----|
|Answer|string|the answer to the question|
|Category|string|category of question|
|Difficulty|string|difficulty of question|
|Text|string|body of the question|

```json
[
    {
        "Answer": "Richard Branson",
        "Category": "General Knowledge",
        "Difficulty": "easy",
        "Text": "Virgin Trains, Virgin Atlantic and Virgin Racing, are all companies owned by which famous entrepreneur?"
    },
    {
        "Answer": "Eight",
        "Category": "General Knowledge",
        "Difficulty": "easy",
        "Text": "How many furlongs are there in a mile?"
    }
]
```
---

#### categories
* **url**: https://triviagoserver.herokuapp.com/api/categories 
* **overview**: this endpoint is used to retrieve Category objects
* **request type**: get

**request example**

```javascript
axios.get("https://triviagoserver.herokuapp.com/api/categories")
```

**response example**

|attribute|type|description|
|----|----|----|
|ID|integer|value representing the category ID|
|Name|string|containing the category name"|

```json
[
    {
        "ID": 9,
        "Name": "General Knowledge"
    },
    {
        "ID": 10,
        "Name": "Entertainment: Books"
    }
]
```
---
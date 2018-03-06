# Go Programming Challenge

Hello 👋
This is our challenge for potential new developer team members.

## Instructions

Fork this repository, and implement the scenario described below.

* You *shall* complete this challenge using **Go language**,
* You *should* try to show your development process to present a **production-ready** code,
* Please, describe your approach, your technical choices, including architectural, and anything you want us to know in a results.md file,

## The Scenario

Your mission is to implement a single endpoint that will be integrated in a larger **microservice** architecture.

* the endpoint is exposed at `/api/v1/trivia`
* the endpoint allows to fetch data from a local database (see explainations below)
* the query may takes parameters, to filter the results. You are expected to propose any kind of query params you may find useful to provide a rich user experience,
* the endpoint returns a JSON response with an array of matching results

## The database

The initial database is provided in a file [db.json](db.json).
You are expected to integrate the data in a mongodb local instance, and explain how to install, launch and populate the database. This database **must** be used by your code.
The schema of the provided data is :

```
{
  "type": "array",
  "items": {
    "type": "object"
    {
      "text": {
        "type": "string",
        "minLength": 1
      },
      "number": {
        "type": "number",
	     "minimum": 1,
	     "maximum": 1e+150
      },
      "found": {
        "type": "boolean"
      },
      "type": {
        "type": "string"
      }
    }
  }
}
```

Example:

```
[
  {
    "text": "1e+40 is the Eddington–Dirac number.",
    "number": 1e+40,
    "found": true,
    "type": "trivia"
  },
  {
    "text": "60 is the years of marriage until the diamond wedding anniversary.",
    "number": 60,
    "found": true,
    "type": "trivia"
  }
]
```

## How we review ?

Your application will be reviewed by at least two of our team member. We do take into consideration your experience level.

**We value quality over feature-completeness**. The goal of this code sample is to help us identify what you consider **production-ready** code. You should consider this code ready for final review with your colleague, i.e. this would be the last step before deploying to production.

The aspects of your code we will assess include:

* **Architecture**: how clean is the separation between the logical layers ? does the code make use of design patterns ? Are the RESTful principles followed ?
* **Clarity**: does the documentation (result.md) clearly and concisely explains the problem and solution? Are technical tradeoffs explained ?
* **Correctness**: does the code do what was asked ? If there is anything missing, does the result.md explain why it is missing?
* **Code quality**: is the code simple, easy to understand, and maintainable ? Are there any code smells or other red flags ? Does it follows coding principles such as the single responsibility principle ? Is the coding style consistent with the language's guidelines ? Is it consistent throughout the codebase ?
* **Usability**: is your API suitable to be used by many different clients, some with memory constraints, some with perfomance issues ? Can your API be used on different type of devices like web, mobile, or IoT ?
* **Security**: are there any obvious vulnerability ?
* **Testing**: how thorough are the automated tests ? Will they be difficult to change if the requirements of the application were to change ? Are there some unit or some integration tests ? We're not looking for full coverage (given time constraint) but just trying to get a feel for your testing skills.
* **Technical choices**: do choices of libraries, databases, architecture etc. seem appropriate for the chosen application ?
* **Documentation**: is your API documented ? With a decent handling of http status code ? With curl examples ? Does it expose a modern documentation (swagger-like) ?

Bonus point (those items are optional):

* **Scalability**: will technical choices scale well ? If not, is there a discussion of those choices in the result.md ?
* **Production-readiness**: does the code include monitoring ? logging ? proper error handling ?
* **Authentication**: does the API allows user authentication ?

## Notes

* We're evaluating you on completion of the scenario, but also on style. Do you commit often? Did you make code that you would enjoy reading ? Things like that are important to our team 👊
* We expect you to provide a code as you would do in any professional situation, with future maintenability and continuous deployment constraints in mind, with deployment or installation instructions, ...
* This should be fun! If you're stuck on an instruction or if something needs clarification, you can email me for help.


## Done ?

Great job! When you're all finished, open a PR and we'll review it together 🙌

## Thanks

* http://numbersapi.com

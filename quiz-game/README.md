# Quiz Game

A CLI program that will read in a quiz provided via a CSV file, and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question is asked immediately afterwards.

***
## Usage
### Flags
  1. `csv` a csv file in the format of question,answer. *Default: "problems.csv"*
  2. `shuffle` shuffle the problems if true, else ignore. *Default: true*
  3. `limit` time limit for the Quiz in seconds. *Default: 30*

### Features
  - Timed Quiz .i.e, the quiz is timed and will accept input only for that duration.
  - Shuffled Quiz .i.e, the problems will be shuffled randomly every time the program executes. 
  - Output the total number of questions correct and how many questions there were in total. *Example: 5/12*
  - Output the total duration of test taken by the User. *Example: 5.04 seconds*

***
## Next Steps
- Add choices/options for the Problem (MCQ).
- Allow Multi correct answers to be included.
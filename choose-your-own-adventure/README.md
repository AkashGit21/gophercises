# Choose Your Own Adventure

**Choose Your Own Adventure** is (was?) a series of books intended for children where as you read you would occasionally be given options about how you want to proceed. For instance, you might read about a boy walking in a cave when he stumbles across a dark passage or a ladder leading to an upper level and the reader will be presented with two options like:

  - *Option 1*: Turn to page 44 to go up the ladder.
  - *Option 2*: Turn to page 87 to venture down the dark passage.

***
## Objective
The goal of this exercise is to recreate this experience via a web application where each page will be a portion of the story, and at the end of every page the user will be given a series of options to choose from (or be told that they have reached the end of that particular story arc).

Stories will be provided via a JSON file.

### Note
- Stories could be cyclical if a user chooses options that keep leading to the same place. This isn’t likely to cause issues, but keep it in mind.
- For simplicity, all stories will have a story arc named “intro” that is where the story starts. That is, every JSON file will have a key with the value intro and this is where your story should start.

***
## Next Steps
- Create a command-line version of our Choose Your Own Adventure application where stories are printed out to the terminal and options are picked via typing in numbers (“Press 1 to venture …”).
- Consider how you would alter your program in order to support stories starting form a story-defined arc. That is, what if all stories didn’t start on an arc named intro? How would you redesign your program or restructure the JSON? This bonus exercises is meant to be as much of a thought exercise as an actual coding one.

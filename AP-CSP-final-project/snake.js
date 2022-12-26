/*   Functions   */

// Base function
function createSquare(x, y, id) {
	button(id, "");

	setProperty(id, "background-color", "blue");

	setPosition(id, x, y, 23, 23);
}

var game_over = false;

function endGame() {
  hideElement("game_title");
	setProperty("game_over_label", "text-color", "#d6d6d6");
	
  clearInterval(autoMove);
  game_over = true;
}

// Possible apple coordinates generator
var x_coors = [];
var y_coors = [];

for (var i = 0; i < 12; i++) {
	var x = 12;
	var y = 107;

	x = x + (25 * i);
	y = y + (25 * i);

	appendItem(x_coors, x);
	appendItem(y_coors, y);
}

// Moves apple to random location
function moveApple() {
	var rand1 = randomNumber(0, 11);
	var rand2 = randomNumber(0, 11);

	setPosition("apple", x_coors[ rand1 ], y_coors[ rand2 ], 25, 25);
}



/*   Initializations   */

// Creates apple
image("apple", "https://upload.wikimedia.org/wikipedia/commons/thumb/0/06/Red_apple.svg/" +
	"768px-Red_apple.svg.png");
setPosition("apple", x_coors[6], y_coors[6], 25, 25);

var score = 1; // Also length of snake
var snake = [ [12, 107] ];

// Creating the snake
createSquare(12, 107, "body0");



/*   Game mechanics   */

var direction;

// wasd sets direction
onEvent("game_screen", "keypress", function(event) {
  if (event.key == "w") {
    direction = 0;
  } else if (event.key == "a") {
    direction = 90;
  } else if (event.key == "s") {
    direction = 180;
  } else if (event.key == "d") {
    direction = 270;
  }
});

function moveSnake() {
  
  var x = snake[0][0];
  var y = snake[0][1];

  if (direction == 0) {
    y -= 25;
  } else if (direction == 90) {
    x -= 25;
  } else if (direction == 180) {
    y += 25;
  } else if (direction == 270) {
    x += 25;
  }
	
	
	var apple_x = getProperty("apple", "x");
	var apple_y = getProperty("apple", "y");

	for (var i = 0; i < score; i++ ) {    // score is # of items in snake array
	  
	  // Checks if any square is touching apple
		if ( (snake[i][0] == apple_x) && (snake[i][1] == apple_y) ) {
			moveApple();
			
			// Adds square to snake
			createSquare(snake[0][0], snake[0][1], "body" + score);
			
			// This so last square doesn't get deleted and snake[i] is defined
			appendItem(snake, [1, 1] );
			
			score += 1;
			setText("score_counter", "Score: " + score);
		}
	}
  
  // Adjusts snake array
  insertItem(snake, 0, [x, y]);
  removeItem(snake, score);
  
  // Runs for each coor in snake (score is amount of coors)
  for (var j = 0; j < score; j++) {
		// Checks if snake ran into itself
		if ( (j > 1) && (snake[0][0] == snake[j][0]) && (snake[0][1] == snake[j][1]) ) {
		  endGame();
		}
		
		// Checks if snake went out of bounds
		if ( (x < 10 || x > 310) || (y < 105 || y > 405) ) {
		  endGame();
		}
		
		// Updates snake's position
		if (!game_over) {
		  setPosition("body" + j, snake[j][0], snake[j][1]);
		}
	}
}


// Snake is moved every 150 ms
var autoMove = setInterval(moveSnake, 150);

autoMove;




/*
Day 1:  Created createSquare function to create snake
        Scapped idea to have board be made of squares
        Made basic board and started to make snake
Day 2:  Created basic movement, can move snake with wasd, doesn't continue moving
        Hid board, just white screen and blue square
        Coded function to wrap snake around board, untested
Day 3:  Completed wrap function
        Tried to use turtle to draw board, didn't work, scrapped it
        Started to work on true movement (learned setInterval() func)
        wasd sets direction
Day 4:  Created placeApple() function, randomly places apple on board
        Worked on collision detection (snake and apple) [WIP]
Day 5:  Took out redundant code and comments
(Sat)   Moved code from the wrapping function into the event listener
        -> Fixed bug where coordinates would be wrong after wrapping snake around screen
        Created checkApple() function, WORKED FIRST TIME!!!!
        Implemented a score counter with on-screen representation
Day 6:  Moved around some code for better readibility
(Sun)   Created addLength() function, works as intended
        Created updateBody(), isn't implemented correctly
Day 7:  Worked out a few annoying bugs and deleted unnecessary code
        Workerd on making updateBody word for multiple squares
        -> Maybe should make sure that one square can follow
        IDEA [Brayden L]: Last square is moved to where the snake is going (for movement)
        -> Need to combine snake_head and snake_body
Day 8:  Created new movement system, new coor inserted, last deleted
        Removed a bunch of code that was basically the same
        Timing STILL isn't working. Saying that the function that is called isn't a funciton
Day 9:  GOT SNAKE TO FOLLOW ITSELF!!!
        -> Changed up how snake is displayed & moved
        Deleted code that wasn't being used
        Started to work on a game over screen
Day 10: Got beginning and end screen working, basic UI
        -> Need to improve (better look, more buttons maybe)
        You can now die!!! ~ never thought I would be so happy to say that... ~
        Added high score functionality
        Removed beginning screen, useless
Day 11: Scrapped playing multiple times to work on true movement
[SB]    IT'S WORKING! IT'S WORKING! IT'S WORKING! IT'S WORKING!
        True movement works (and always was), just needed to move snake piece(s)
        Started to experiment with removing wrapping functionality
        Removed end screen as well, title changes to GAME OVER and play stops


Further Possible Improvements:
  Improve UI
  Be able to play multiple times
  Easy, Medium, Hard mode where the snake's movement speed is altered
*/

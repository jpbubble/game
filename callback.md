# CallBack


When in CallBack mode, the system will call for the next functions if it's been dubbed "needed".
Certain events can call this to happen.

Please note these functions are only called in CallBack mode (unless explicitly stated otherwise).



## bubble_init()

Well this one is called in all flowmodes, not only CallBack. This function is executed right after the script is loaded.



## bubble_quit()

Called when a the user tries to quit the application.
When not set this will just end the program immediately.
Please note, this function neither gets parameters no does it need to return anything.
This function should contain an EndCycle() or Crash() command to do the actual quitting.

Want to program to cause this effect manually?
Sure just call the function directly.


## bubble_keydown(keyname)

Called when a user pushes a key.
The keyname variable will contain the name of the key.
A full list of these keynames may be uploaded later.

## bubble_keyup(keyname)

Called when a user releases a key
The keyname variable will contain the name of the key.
A full list of these keynames may be uploaded later.

## bubble_mousemove(x,y)

Called when the mouse is moved. The arguments contain the coordinates.

## bubble_mousedown(x,y,button)

Called when the mouse button is hit

## bubble_mouseup(x,y,button)

Called when the mouse button is released

## bubble_draw()

Is called whenever the system expects input on what to draw.
After this function is called the results will be shown immediately on screen.

Please note! Although possible, it's in CallBack mode pretty pointless to put in any drawing commands outside of this function.
Either the results will never be seen or the outcome can be quite unpredictable.

## bubble_tick()

Executed when the set amount of nano seconds have past since the last check.
When your project handles this to quickly or to slowly set the amount of nanoseconds with the "SetTick" command.

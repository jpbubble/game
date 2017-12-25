# Command reference

Engines using this module have the next commands added to Lua.
(I will not teach you Lua, only give you an overview of the commands
lua uses).


## System

These are system based commands.

### Crash([exitcode=0])

- Exits bubble immediately causing an exit code for the underlying OS
- When you need an immediately shutdown you should always use Crash and never os.exit(). os.exit() has not regard for anything that needs to be unloaded before the system can safely be shutdown, and this can cause memory leaks. Crash will however do all needed unloading stuff.

### EndFlow()

- As soon as the end of the current cycle is reached bubble will exit.
- Contrary to Crash, EndFlow will not terminate the program immediately. The current flow will be completed and after that it ends and not sooner.
- Since "Static" does not work in flows this command is not available in Static mode.


### EngineVersion()

- Returns the current engine version number.

### Identify(key)

- Returns the value of an identify.gini variable.
- Yes! These are read-only!



## Events

Please note, how these are used and if you can even use these is really reliant on which flowmode you chose in your identify.gini
In fact this is where that setting can matter most.
Especially when aiming for CallBack these can work a little different than other modes.

### SetTick(nanoseconds)
- Static:   Sets waiting time for WaitTick()
- Cylic:    Sets waiting time to wait at the end of each cycle. WaitTick() can still be used, but best is to avoid that.
- CallBack: Sets the time between two "bubble_tick" calls.
Please note, the time is the minimal ammount of time that must have passed between two checks. If due to system lagging this took longer, there should be a shorter wait to compensate or no wait at all.
This routine is set up to make sure that even on faster computers projects would not go too fast.

### WaitTick()
Waits until enough time has been passed since the last check
- Static:   A regular use of this can be handy. Mostly recommended right after a Flip() command.
- Cyclic:   Best is to avoid this, although it may have some use somewhere, and that's why I left it in.
- CallBack: Not available.

### PollEvents()
Updates all input event data. 
- Static:   You will always need this to be up-to-date on data
- Cyclic:   It's automatically called at the start of each cycle, but you can still call it manually.
- CallBack: Not available as the system will call for event functions to help you handle this. The functions returning values depending on the outcome of this function do still work normally (as the callback mode calls this function automatically when it needs to).

Except for the Static mode this function will always be called automatically at the start of each cycle. This is handy to know for how the functions responding to this call have their values.

### KeyPressed(keyname)
Returns "true" if the key has been pressed since the last PollEvents() call or cycle start.

### KeyDown(keyname)
Returns "true" if the key was held down on the last PollEvents() call or cycle start.

### MouseHit(button)
Returns "true" if the mousebutton was hit since the last PollEvents() call or cycle start.

### MouseDown(button)
Returns "true" if the mousebutton was held down on the last PollEvents() call of cycle start.

### MouseCoords()
Returns the mouse coordinates as they were during the last PollEvents() call or cycle start.
~~~lua
x,y = MouseCoords()
~~~

### ClosureRequested()
Returns "true" if the user tried to shut down the program.



## Images/Graphics

### LoadImage(file,[target=""])

- Will load the image and return a key (as a string) which can be used to reference the file later
- When "target" is set this key will be of your choosing. Please note they are case sensitive. It will then return the same key by the way.
- When using "target" on a key that already exists it will automatically be freed, so you don't need to free it first.

### AssignImage(source,target,[autofreeoriginal=1])

- Will assign an image to a new target
- Unless you specifically set autofreeoriginal to 0 or any negative number the original will be freed
- Please note, this function only copies pointers, not actual data.

### FreeImage(imgkey)

- Will remove an image from memory
- If multiple pointers to the same image exist the image will not be delete from memory, just remove the pointer to it, if this is the last pointer then the image data will be removed from the memory.

### Cls()

- Clears the screen in the current cls color

### DrawImage(imgkey,x,y,[frame=0])

- Draws the image on the screen.
- Frame is only applicable when dealing with multi-picture resources. Bundles so to speak.
- First frame is always 0 and the highest always number of frames -1. So if there are 10 frames you have a margin of 0 till 9.

### ImageFrames(imgkey)

- Returns the number of frames an image has


### Flip()

- By default all drawing happens on a backbuffer. This command brings it up front
- When you use Bubble in "CallBack" mode, this function will not be availbe, as the action will automatically be performed after completing the callback function "bubble_draw()"


## Windows

I don't mean the OS Windows. I mean actions concerning the working windows. :P

## Win_GetSize()

- Returns the sizes of the Window

~~~Lua
width,height = Win_GetSize()
~~~

Like that.

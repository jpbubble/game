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

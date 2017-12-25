# Identify.gini

The file "Identify.gini" must always be present in a folder named "ID" inside the JCR6 resource a bubble program uses.
It contains several kinds of configuration data and maybe also some meta data if you like.

The definitions the bubble engine uses described below. All other definitons you put in there are ignored.


The file can look like this:
~~~
[vars]
TITLE=My Game
COPYRIGHT=(c) Me, myself and I, 2017
START.VM=DEFAULT
START.SCRIPT=Script/Start.lua
~~~


Yeah it's just a "variable name" plus its value.
Gini itself is case insensitve so "TITLE" and "TiTlE" will be seen as the same variable.
The values may be case sensitive though.
The order in which you put them in doesn't matter. Gini will parse everything anyway before Bubble can handle it, and Bubble can easily read all parsed data from Gini's memory.



Here are the variables you can define and what they mean.
Variables marked with a "*" are required. Others are optional


### TITLE
Contains the title of your project. It will be used to create the title of the window.


### ENGINE *
Should contain the engine name used. If it's not correct Bubble will refuse it.
When you just use the basic Bubble engine it should contain the word "BUBBLE".
But as Bubble is a base that can be combined with other engines the name can then be different.
Bubble checks this value to prevent resources being mixed up with engines. Can give rather messy effects, you know.

### WIDTH
Width of the window


### HEIGHT
Height of the window

### CONSOLE.BCK.R / CONSOLE.BCK.G / CONSOLE.BCK.B
These can be used for the background of your debug console. 
When not set the background will be black.
The values may be numbers from 0 till 255

### CONSOLE.CMD.R / CONSOLE.CMD.G / CONSOLE.CMD.B
The command font color of your debug console.
When not set the color will be white.
The values may be numbers from 0 till 255

### FLOW.MODE
When not set it will be "Static".
There are three modes in which Bubble can operate and it's important you use the correct one.
1. Static
   - Static is very basic. It will just load one script and execute its bubble_init() function if it exists and then its bubble_main() function. Once all that is done it will just exit Bubble.
   - Static can be a good way for beginners to get on the road or for simple experimental programs. It may not be the best approach for full projects, but maybe there is.
   - Static only considers "MAIN" as its main vm in multiscripting. It can use alternate vms for "help", but not for "flow" changes.
   - You will need to perform a "PollEvents()" call every now and then to get ahold of mouseclicks and key hits etc.
2. Cyclic
   - Upon starting the starting vm will execute bubble_init() if it exists and after that it will keep repeating bubble_cycle() eternally unless the script tells bubble to stop.
   - Prior to the start of a cycle all events triggered will be read. "PollEvents()" can still be used, but only do so if you know what you are doing.
   - By changing flows you can switch between vms and bubble will then repeat their bubble_cycle() in stead. This can be handy for software with several work screens. My RPGs "Star Story" and "The Fairy Tale REVAMPED" heavily rely on that kind of working (although not written in a Bubble related engine, but the mechanism is the same. The LAURA II engine works also a bit in this Cyclic manner).
3. CallBack
   - In CallBack mode, Bubble will as in the other two modes first read bubble_init(), but after that it will just respond on events and call functions based on these. "PollEvents()" will not be available in this mode as it will very likely spook things up.
   - I will write another .md file on the functions the callback mode calls on later.
   - The change flow mechanism can also be used in the same way as in the Cyclic mode.
The names of these mode are CASE SENSITIVE. Look well in how I used lower and upper case!


### START.VM
Set the basic flow VM to start with. By default this is "MAIN", but you can desire any name you want.
When you use the "Static" mode, this value should not be touched as multi-flowing is not supported in Static (and changing this value in Static will cause bubble to throw an error).

### START.SCRIPT
This is the script with which bubble should start your project. When not set it will look for "Script/Main.lua".
Oh yeah, this value can safely be changed in all modes, including Static ;)



## Notes
- Like I said everything that Bubble doesn't know is ignored. It can be read by the Lua scripts though if you desire.
- Always take in mind though that more variables can get a function as Bubble gets expanded or reaches higher levels of sophistication. I do recommend therefore to name your own vars in a way you can be sure it will never turn into a reserved variable for Bubble.



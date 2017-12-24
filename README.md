Broks Ultimately Brings you a Basic Lua Engine => BUBBLE

# Bubble Game Engine

This the base for the Game Engine. These are just imports. The actual files to build an executable binary are in another repo.
In other words these files do nothing by themselves, but the files in the actual executables repos do.

Why this setup?
- Most of all for extendability. The bubble engine has been setup to be easily attachable other setups creating sub-engines that are more able to deal with certain tasks. Lua is as an interpreted lanugae not too fast (already faster than you'd expect from an interpreted language), plus the Go version of Lua is a bit slower, so allowing the engines to do more can be desirable, but at the same time why give users of the software made with Bubble more than they need? Only takes up RAM. Wll, I guess I motivated my approach enough now :P
- Please note that my Blitzmax engines are already a bit set up in this perspective. LAURA II was a combination of JCR6, GALE and small library with RPG specific routines that was actually meant for the DuCraL engine that never got finished, and of course Kthura :P
- With Kthura I already named a heavy system. It's not yet decided if Kthura makes it to Go, but if it does, I only want it in engines for projects that NEED Kthura.



# Requirements to build

- Tricky units for Go
- veandco's SDL2 libraries for Go -- https://github.com/veandco/go-sdl2
- The libs deemed optional in the documentation of SDL are NOT optional as far as Bubble is concerned!!!
- Shopify's Lua engine -- https://github.com/Shopify/go-lua
- The jpbubble/Base library

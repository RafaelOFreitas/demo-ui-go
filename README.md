# Build cross platform desktop applications with HTML, CSS and script!

![Sciter](img/sciter.png)

### Examples: 

+ [Hello World](ui-hello-world/).
+ [Notes](ui-notes/).
+ [Login Discord-clone](ui-login/).
+ [To-Do List](ui-to-do-list/).
+ [Splash Screen](ui-splash-screen/).
+ [Project Tools](ui-example-project/).

### Guide:

+ Create Cross-Platform desktop applications with HTML, CSS and script!

+ Creating desktop screens is easy using the [sciter](https://sciter.com/).

+ See more on [GitHub](https://github.com/sciter-sdk/go-sciter).

+ Perhaps these [videos](https://www.youtube.com/playlist?list=PLub5C2vM5SjKvkbFfposhyg1V2gpXnviM) help you understand a little bit about the library.

+ Some [Tutorials](https://sciter.com/tutorials/).

+ Understand about [Sciter architecture](https://sciter.com/developers/engine-architecture/).

+ [Documentation](https://sciter.com/developers/sciter-docs/).

+ About [Licensing and pricing](https://sciter.com/prices/).

+ Sciter Wizard for [Visual Studio Code](https://sciter.com/sciter-assistant-for-visual-studio-code/).

+ [Development Tools](https://sciter.com/developers/development-tools/).

+ My impressions:
  
  - Ease of use.
  - With a few kb, we do a lot, unlike the Electron. Which also produces cross-platform desktop applications.
  - I didn't find the documentation and items of this in an easy way, however on the [website](https://sciter.com/) you can search for problems efficiently.
  - Embeddable, can be used with any programming language.
  - Does not support all CSS 3 features 😥.
  - We do not use JavaScript, but TIScript is an extended version of ECMAScript. Which is not necessarily a bad thing, but at the beginning it produces a certain discomfort, see more [here](https://sciter.com/developers/for-web-programmers/tiscript-vs-javascript/)

## How to test location:

+ If the OS is Linux or MacOS just run `config-sciter.sh`
  
+ Configure Sciter according to the OS
  + At the moment, only Go 1.10 or higher is supported.
  + Download [sciter-sdk](https://sciter.com/download/).
  + Extract the sciter library `sciter-sdk`.
  + The runtime libraries are in /bin bin.lnx | bin.osx suffixed like dll sooudylib.
  + Windows: just copy `bin\64\sciter.dll` for `c:\windows\system32`
  + Linux:
  
    ```shell
        user@user:~$ cd sciter-sdk/bin.lnx/x64
        user@user:~$ export LIBRARY_PATH=$PWD
        user@user:~$ echo $PWD >> libsciter.conf
        user@user:~$ sudo cp libsciter.conf /etc/ld.so.conf.d/
        user@user:~$ sudo ldconfig
        user@user:~$ ldconfig -p | grep sciter
    ```

  + OSX:
  
    ```shell
        cd sciter-sdk/bin.osx/
        export DYLD_LIBRARY_PATH=$PWD
    ```

  + Set up the GCC environment for CGO.
  + [mingw64-gcc](https://sourceforge.net/projects/mingw-w64/) (5.2.0 e 7.2.0 are tested) is recommended for users of Windows.
  + At the Linux, gcc (4.8 or superior) e gtk + -3.0 are necessary.
    + Install the build-essential package (The command installs several new packages including gcc, g++ and make.) typing:

    ```shell
        user@user:~$ sudo apt install build-essential
    ```

    + Instale o gtk + -3.0
  
    ```shell
        user@user:~$ sudo apt-get install build-essential libgtk-3-dev
        user@user:~$ dpkg -L libgtk-3-dev | grep '\.pc'
        user@user:~$ pkg-config --modversion gtk+-3.0
    ```
  
  + Import the Sciter library with:

    ```shell
        go get -x github.com/sciter-sdk/go-sciter
    ```
  

+ My impressions:
  
  - Ease of use.
  - With a few kb, we do a lot, unlike the Electron. Which also produces cross-platform desktop applications.
  - I didn't find the documentation and items of this in an easy way, however on the [website](https://sciter.com/) you can search for problems efficiently.
  - Embeddable, can be used with any programming language.
  - Does not support all CSS 3 features 😥.
  - We do not use JavaScript, but TIScript is an extended version of ECMAScript. Which is not necessarily a bad thing, but at the beginning it produces a certain discomfort, see more [here](https://sciter.com/developers/for-web-programmers/tiscript-vs-javascript)
# voop

Voop is meant to be simple and lightweight videoart tool.

It can sync [relatively short] video clips with running music software via Ableton Link.

If you want to check what Voop can do you should run it with a path to your video

```
go run . --folder="folder/with/your/videos"
```

it requires running [Carabiner](https://github.com/Deep-Symmetry/carabiner) on 127.0.0.1:17000

and opencv installed locally

on mac it usually can be done with brew
```
brew install opencv
```

Voop has some shortcuts:
'>' - next video<br/>
'<' - previous video<br/>
'/' - random video<br/>
']' - next folder<br/>
'[' - previous folder<br/>
'§' - random folder<br/>
'=' - faster<br/>
'-' - slower<br/>
'0' - default speed<br/>
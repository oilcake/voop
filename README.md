# voop
Voop is meant to be simple and lightweight videoart tool.<br/>
It can sync [relatively short] video clips with running music software via Ableton Link.
## TL;DR
Place [this](https://ln5.sync.com/dl/efbca6d10/4z9tgsq7-ifxw5mm2-fsqm7hek-683t2gaq) in the "./samples" in the root of this repo.
open [Carabiner](https://github.com/Deep-Symmetry/carabiner), and run Voop:
```
go run .
```


You can run it with a path to your videos

```
go run . --folder="folder/with/your/videos"
```

Running [Carabiner](https://github.com/Deep-Symmetry/carabiner) on 127.0.0.1:17000 is a must - because currently Link is the only supported transport.
The simplest way is to download a pre built [binary](https://github.com/Deep-Symmetry/carabiner/releases)

You also need OpenCV.<br/>
On mac it can be installed via brew
```
brew install opencv
```
Note that video codec is critical - no application in the world will be backward-happy with temporal encoding like mp4. Best option so far is Apple ProRes.
Here's some [ProRes encoded samples](https://ln5.sync.com/dl/efbca6d10/4z9tgsq7-ifxw5mm2-fsqm7hek-683t2gaq) 
if you place it in ./samples in the root of this repo you can run voop with no arguments:
```
go run .
```
you can also run Voop with customized config:
```
go run . --config="path/to/your/config"
```
if no config is specified Voop will be started with default ./config.yml, which can be customized as well. It is also a good overview of actions that could be done with Voop.

Some of Voop's default shortcuts: <br/>
'>' - next video<br/>
'<' - previous video<br/>
'/' - random video<br/>
']' - next folder<br/>
'[' - previous folder<br/>
'§' - random folder<br/>
'=' - faster<br/>
'-' - slower<br/>
'0' - default speed<br/>
